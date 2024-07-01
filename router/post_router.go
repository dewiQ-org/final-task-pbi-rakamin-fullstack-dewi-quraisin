package router

import (
	"project-api-golang/config"
	"project-api-golang/handler"
	"project-api-golang/middleware"
	"project-api-golang/repository"
	"project-api-golang/service"

	"github.com/gin-gonic/gin"
)

func PostRouter(api *gin.RouterGroup) {
	postRepository := repository.NewPostRepository(config.DB)
	postService := service.NewPostService(postRepository)
	postHandler := handler.NewPostHandler(postService)

	r := api.Group("/upload")
	r.Use(middleware.JWTMiddleware())
	r.POST("/", postHandler.Create)
}
