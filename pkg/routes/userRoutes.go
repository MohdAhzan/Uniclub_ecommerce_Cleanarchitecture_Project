package routes

import (
	"project/pkg/api/handler"
	"project/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(engine *gin.RouterGroup,
	userHandler *handler.UserHandler,
	otpHandler *handler.OtpHandler, inventoryHandler *handler.InventaryHandler,
	cartHandler *handler.CartHandler, orderHandler *handler.OrderHandler) {

	engine.POST("/signup", userHandler.UserSignUp)
	engine.POST("/login", userHandler.UserLoginHandler)
	engine.POST("/otplogin", otpHandler.SendOTPHandler)
	engine.POST("/verifyotp", otpHandler.VerifyOTPHandler)

	engine.Use(middleware.UserAuthMiddleware)
	{
		home := engine.Group("/home")
		{
			home.GET("", inventoryHandler.GetProductsForUsers)
			home.POST("/add_to_cart", cartHandler.AddtoCart)
		}

		cart := engine.Group("/cart")
		{
			cart.GET("", cartHandler.GetCart)
			cart.DELETE("/remove", cartHandler.RemoveCart)
			// cart.PUT("/cartQuantity/plus", cartHandler.PlusCartQuantity)
			// cart.PUT("/cartQuantity/minus", cartHandler.MinusCartQuantity)
		}

		profile := engine.Group("/profile")

		{
			profile.GET("/details", userHandler.GetUserDetails)
			profile.GET("/address", userHandler.GetAddressess)
			profile.POST("/address", userHandler.AddAddressess)
			profile.DELETE("/address", userHandler.DeleteAddress)

			edit := engine.Group("/edit")
			{
				edit.PUT("/account", userHandler.EditUserDetails)
				edit.PUT("/address", userHandler.EditAddress)

				edit.PUT("/password", userHandler.ChangePassword)
			}

			orders := profile.Group("/orders")
			{
				// 	orders.GET("", orderHandler.GetOrders)
				// 	orders.GET("/:id", orderHandler.GetIndividualOrderDetails)
				// 	orders.DELETE("", orderHandler.CancelOrder)
				// 	orders.PUT("/return", orderHandler.ReturnOrder)
				// orders.GET("/checkout", orderHandler.Checkout)
				orders.POST("", orderHandler.OrderFromCart)
			}

		}
	}
}
