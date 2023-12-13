package routes

import (
	"project/pkg/api/handler"

	"github.com/gin-gonic/gin"
)

func UserRoutes(engine *gin.RouterGroup,
	userHandler *handler.UserHandler) {

	engine.POST("/signup", userHandler.UserSignUp)
}
