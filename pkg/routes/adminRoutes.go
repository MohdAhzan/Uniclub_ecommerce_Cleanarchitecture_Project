package routes

import (
	"project/pkg/api/handler"
	"project/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(engine *gin.RouterGroup,
	adminHandler *handler.AdminHandler) {

	engine.POST("/adminlogin", adminHandler.LoginHandler)

	engine.Use(middleware.AdminAuthMiddleware)
	{
		userManagement := engine.Group("/users")
		{
			userManagement.GET("", adminHandler.GetUsers)
			userManagement.PUT("/block", adminHandler.BlockUser)
			userManagement.PUT("/unblock", adminHandler.UnBlockUser)
		}

	}
}
