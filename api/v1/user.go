package v1

import (
	"database/sql"
	"ginApi/models"
	"ginApi/storage/repo"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *handlerV1) CreateUSer(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": "internal server error",
			"error":   err.Error(),
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
		ctx.JSON(200, gin.H{
			"message": "internal server error",
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
			"error":  "invalij ID",
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

func (h handlerV1) Delete(ctx *gin.Context) {
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
