package api

import (
	_ "ginApi/api/docs"
	v1 "ginApi/api/v1"
	"ginApi/storage"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RouterOptions struct {
	Storage storage.StorageI
}

// @title           Swagger for user api
// @version         1.0
// @description     This is a blog service api.
// @BasePath  /v1
func New(opt *RouterOptions) *gin.Engine {
	router := gin.Default()

	// CORS configuration
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "*")
	router.Use(cors.New(corsConfig))

	// API v1 handlers
	handlerV1 := v1.New(&v1.HandlerV1Options{
		Storage: opt.Storage,
	})

	// Grouping v1 endpoints
	v1Router := router.Group("/v1")
	{
		v1Router.POST("/create", handlerV1.CreateUser)
		v1Router.GET("/users/:id", handlerV1.GetUser)
		v1Router.PUT("/update/:id", handlerV1.Update) // assuming update with an ID
		v1Router.DELETE("/delete/:id", handlerV1.Delete)
		v1Router.GET("/getbyemail/:email", handlerV1.GetByEmailHandler)
		v1Router.GET("/getall", handlerV1.GetAll)
	}

	// Swagger endpoint for documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
