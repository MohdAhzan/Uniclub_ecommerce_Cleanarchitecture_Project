package routes

import (
	"project/pkg/api/handler"
	"project/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(engine *gin.RouterGroup,
	userHandler *handler.UserHandler,
	otpHandler *handler.OtpHandler, inventoryHandler *handler.InventaryHandler) {

	engine.POST("/signup", userHandler.UserSignUp)
	engine.POST("/login", userHandler.UserLoginHandler)
	engine.POST("/otplogin", otpHandler.SendOTPHandler)
	engine.POST("/verifyotp", otpHandler.VerifyOTPHandler)

	engine.Use(middleware.UserAuthMiddleware)
	{
		home := engine.Group("/home")
		{
			home.GET("", inventoryHandler.GetProductsForUsers)
		}

		profile := engine.Group("/profile")

		{
			profile.GET("/details", userHandler.GetUserDetails)
			profile.GET("/address", userHandler.GetAddressess)
			profile.POST("/address", userHandler.AddAddressess)

			edit := engine.Group("/edit")
			{
				edit.PUT("/account", userHandler.EditUserDetails)
				edit.PUT("/address", userHandler.EditAddress)


				edit.PUT("/password", userHandler.ChangePassword)
			}

		}
	}
}
