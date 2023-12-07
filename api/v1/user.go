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

// @Router /create [post]
// @Summary Create a user
// @Description Create user
// @Tags user
// @Accept json
// @Produce json
// @Param user body models.CreateUser true "CreateUser"
// @Success 200 {object} models.User
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

// @Router /users/{id} [get]
// @Summary Get user by id
// @Description Get user by id
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.User
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
	ctx.JSON(200, models.User{
		ID:        result.ID,
		FirstName: result.FirstName,
		LastName:  result.LastName,
		Email:     result.Email,
		CreatedAt: result.CreatedAt,
	})

}

// @Router /update/{id} [put]
// @Summary Update user by id
// @Description Update user by id
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Param user body models.UpdateUser true "UpdateUser"
// @Success 200 {object} models.User
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
		ID:        ctx.Param("id"),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(200, models.User{
		ID:        result.ID,
		FirstName: result.FirstName,
		LastName:  result.LastName,
		Email:     result.Email,
		CreatedAt: result.CreatedAt,
	})
}

// @Router /delete/{id} [delete]
// @Summary Delete user by ID
// @Description Delete user by ID
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "id"
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

// @Router /getbyemail/{email} [get]
// @Summary Get users by email
// @Description Get user by email
// @Tags user
// @Accept json
// @Produce json
// @Param email path string true "email"
// @Success 200 {object} models.User
func (h *handlerV1) GetByEmailHandler(ctx *gin.Context) {
	email := ctx.Param("email")
	if email == "" {
		ctx.JSON(500, gin.H{
			"message": "email not found from database",
			"error":   "email not given by user or invalid request",
		})
		return
	}
	result, err := h.Storage.User().GetByEmail(email)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": " error from database",
			"err":     err.Error(),
		})
		return
	}
	ctx.JSON(200, models.User{
		ID:        result.ID,
		FirstName: result.FirstName,
		LastName:  result.LastName,
		Email:     result.Email,
		CreatedAt: result.CreatedAt,
	})
}

// @Router /getall [get]
// @Summary Get all users
// @Description Get all users
// @Tags user
// @Accept json
// @Produce json
// @Param filter query models.GetAllUsersParams  ture "GetAllUsersParams"
// @Success 200 {object} models.GetAllUsersResult
func (h handlerV1) GetAll(ctx *gin.Context) {

	params, err := ValidateGetAllParams(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
			"error":   err.Error(),
		})
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
