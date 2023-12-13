package routes

import (
	"project/pkg/api/handler"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(engine *gin.RouterGroup,
	adminHandler *handler.AdminHandler) {

	engine.POST("/adminlogin", adminHandler.LoginHandler)

}
