package interfaces

import (
	"project/pkg/utils/domain"
	"project/pkg/utils/models"
)

type OrderUseCase interface {
	OrderFromCart(order models.Order, couponID int) error
	Checkout(userID, couponID int) (models.CheckOut, error)
	GetOrders(id int) ([]domain.OrderDetailsWithImages, error)
	GetOrderDetailsByOrderID(orderID, userID int) (domain.OrderDetails, error)
	CancelOrder(orderID, userID int) error
	ReturnOrder(orderID, userID int) error
}
