package v1

import (
	"database/sql"
	"ginApi/models"
	"ginApi/storage/repo"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func (h *handlerV1) CreateUser(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"error":   err.Error(),
		})
		return
	}
	if user.Email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Email must not be empty",
			"error":   "Email not found",
		})
		return
	}

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, err := regexp.MatchString(emailRegex, user.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	}
	if !matched {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "User not created",
			"error":   "Invalid email address",
		})
		return
	}
	createdUser, err := h.Storage.User().Create(&repo.User{
		ID:        uuid.New(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: time.Now(),
		DeletedAt: sql.NullTime{},
	})
	if err != nil {
		pqErr, ok := err.(*pq.Error)
		if ok && pqErr.Code.Name() == "unique_violation" && pqErr.Constraint == "email_unique" {
			ctx.JSON(http.StatusConflict, gin.H{
				"message": "Email already exists",
				"error":   pqErr.Error(),
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, createdUser)
}
func (h *handlerV1) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "UserId not fount in url",
			"error":   "id=null_string",
		})
		return
	}
	result, err := h.Storage.User().Get(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"mesage": "your id not found from database",
			"error":  "invalid ID",
		})
		return
	}
	ctx.JSON(200, result)

}

func (h *handlerV1) Update(ctx *gin.Context) {
	var user models.UpdateUser
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
			"error":   err,
		})
		return
	}
	result, err := h.Storage.User().Update(&repo.UpdateUser{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
			"error":   err,
		})
		return
	}
	ctx.JSON(200, result)
}

func (h *handlerV1) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "UserId not fount in url",
			"error":   "id=null_string",
		})
		return
	}
	err := h.Storage.User().Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Filed to delete from postgres",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "User deleted",
	})
}

func (h *handlerV1) GetByEmailHandler(ctx *gin.Context) {
	email := ctx.Param("email")
	if email == "" {
		ctx.JSON(500, gin.H{
			"message": "email not found from database",
			"error":   "email not given by user or invalid request",
		})
		return
	}
	datbaseRsult, err := h.Storage.User().GetByEmail(email)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": " error from database",
			"err":     err.Error(),
		})
		return
	}
	ctx.JSON(200, datbaseRsult)
}

func (h handlerV1) GetAll(ctx *gin.Context) {
	var params models.GetAllUsersParams
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(500, gin.H{
			"message": "filed to unmarshall JSON",
			"error":   err.Error(),
		})
		return
	}

	users, err := h.Storage.User().GetAll(&repo.GetAllUsersParams{
		Search: params.Search,
		Limit:  params.Limit,
		Page:   params.Page,
	})
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Filed to get from database",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(200, users)

}
