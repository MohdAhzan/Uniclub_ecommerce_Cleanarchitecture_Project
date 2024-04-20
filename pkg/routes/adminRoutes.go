package routes

import (
	"project/pkg/api/handler"
	"project/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(engine *gin.RouterGroup,
	adminHandler *handler.AdminHandler,
	categoryHandler *handler.CategoryHandler,
	inventoryHandler *handler.InventaryHandler,
) {

	engine.POST("/adminlogin", adminHandler.LoginHandler)

	engine.Use(middleware.AdminAuthMiddleware)
	// Create an HTTP client
	// client := &http.Client

	{
		userManagement := engine.Group("/users")
		{
			userManagement.GET("", adminHandler.GetUsers)
			userManagement.PUT("/block", adminHandler.BlockUser)
			userManagement.PUT("/unblock", adminHandler.UnBlockUser)
		}

		categorymanagement := engine.Group("/category")
		{
			categorymanagement.GET("", categoryHandler.GetCategory)
			categorymanagement.POST("", categoryHandler.AddCategory)
			categorymanagement.PUT("", categoryHandler.UpdateCategory)
			categorymanagement.DELETE("", categoryHandler.DeleteCategory)

		}

		productmanagement := engine.Group("/products")
		{
			productmanagement.POST("", inventoryHandler.AddInventory)
			productmanagement.GET("", inventoryHandler.GetProductsForAdmin)
			productmanagement.DELETE("", inventoryHandler.DeleteInventory)
			productmanagement.PUT("/:id/edit_details", inventoryHandler.EditInventoryDetails)
		}
		ordermanagement := engine.Group("/orders")
		{
			ordermanagement.GET("",adminHandler.GetAllOrderDetails)				
			// ordermanagement.GET("/:id",GetOrderDetailsByID)
			ordermanagement.PUT("/payment-status", adminHandler.MakePaymentStatusAsPaid)
			ordermanagement.PUT("/status", adminHandler.EditOrderStatus)
			ordermanagement.PUT("/return", adminHandler.OrderReturnApprove)
		}

	}
}
