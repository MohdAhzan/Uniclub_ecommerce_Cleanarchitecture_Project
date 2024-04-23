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

		search := engine.Group("/search")
		{
			search.GET("", inventoryHandler.SearchProducts)
		}
		cart := engine.Group("/cart")
		{
			cart.GET("", cartHandler.GetCart)
			cart.PUT("", cartHandler.DecreaseCartQuantity)
			cart.DELETE("/remove", cartHandler.RemoveCart)
		}

		profile := engine.Group("/profile")

		{
			profile.GET("/details", userHandler.GetUserDetails)
			profile.GET("/address", userHandler.GetAddressess)
			profile.POST("/address", userHandler.AddAddressess)
			profile.DELETE("/address", userHandler.DeleteAddress)

			wallet := engine.Group("/wallet")
			{
				wallet.GET("", userHandler.GetWallet)
			}

			edit := engine.Group("/edit")
			{
				edit.PUT("/account", userHandler.EditUserDetails)
				edit.PUT("/address", userHandler.EditAddress)

				edit.PUT("/password", userHandler.ChangePassword)
			}
		}
		orders := engine.Group("/orders")
		{
			orders.GET("", orderHandler.GetOrders)
			orders.GET("/:id", orderHandler.GetOrderDetailsByOrderID)
			orders.DELETE("", orderHandler.CancelOrder)
			orders.PUT("/return", orderHandler.ReturnOrder)
			orders.GET("/checkout", orderHandler.Checkout)
			orders.POST("", orderHandler.OrderFromCart)
		}

	}
}
