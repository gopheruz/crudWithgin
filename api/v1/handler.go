package v1

import (
	"ginApi/models"
	"ginApi/storage"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handlerV1 struct {
	Storage storage.StorageI
}

type HandlerV1Options struct {
	Storage storage.StorageI
}

func New(options *HandlerV1Options) *handlerV1 {

	return &handlerV1{
		Storage: options.Storage,
	}
}

func ValidateGetAllParams(ctx *gin.Context) (*models.GetAllUsersParams, error) {
	var (
		limit int = 10
		page  int = 1
		err   error
	)
	if ctx.Query("limit") != "" {
		limit, err = strconv.Atoi(ctx.Query("limit"))
		if err != nil {
			return nil, err
		}
	}
	if ctx.Query("page") != "" {
		page, err = strconv.Atoi(ctx.Query("page"))
		if err != nil {
			return nil, err
		}
	}
	return &models.GetAllUsersParams{
		Limit: int32(limit),
		Page: int32(page),
		Search: ctx.Query("search"),
	}, nil
}
