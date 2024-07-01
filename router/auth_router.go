package router

import (
	"project-api-golang/config"
	"project-api-golang/handler"
	"project-api-golang/repository"
	"project-api-golang/service"

	"github.com/gin-gonic/gin"
)

func AuthRouter(api *gin.RouterGroup) {
	authRepository := repository.NewAuthRepository(config.DB)
	authService := service.NewAuthService(authRepository)
	authHandler := handler.NewAuthHandler(authService)

	api.POST("/register", authHandler.Register)
	api.POST("/login", authHandler.Login)
}
