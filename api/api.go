package api

import (
	v1 "ginApi/api/v1"
	"ginApi/storage"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RouterOptions struct {
	Storage storage.StorageI
}

func New(opt *RouterOptions) *gin.Engine {
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "*")
	router.Use(cors.New(corsConfig))
	handlerV1 := v1.New(&v1.HandlerV1Options{
		Storage: opt.Storage,
	})
	v1Router := router.Group("/v1")
	v1Router.POST("/create", handlerV1.CreateUSer)
	v1Router.GET("/users/:id", handlerV1.GetUser)
	v1Router.PUT("/update", handlerV1.Update)
	v1Router.DELETE("/delete/:id", handlerV1.Delete)
	return router
}
