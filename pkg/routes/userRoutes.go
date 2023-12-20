package routes

import (
	"project/pkg/api/handler"

	"github.com/gin-gonic/gin"
)

func UserRoutes(engine *gin.RouterGroup,
	userHandler *handler.UserHandler,
	otpHandler *handler.OtpHandler) {

	engine.POST("/signup", userHandler.UserSignUp)
	engine.POST("/login", userHandler.UserLoginHandler)
	engine.POST("/otplogin", otpHandler.SendOTPHandler)
	engine.POST("/verifyotp", otpHandler.VerifyOTPHandler)
}
