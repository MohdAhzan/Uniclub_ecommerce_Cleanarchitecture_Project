package interfaces

import (
	"project/pkg/utils/domain"
	"project/pkg/utils/models"
)

type OrderUseCase interface {
	OrderFromCart(order models.Order) error
	Checkout(userID int) (models.CheckOut, error)
	GetOrders(id int) ([]domain.OrderDetailsWithImages, error)
	GetOrderDetailsByOrderID(orderID, userID int) (domain.OrderDetails, error)
	CancelOrder(orderID int) error
}
