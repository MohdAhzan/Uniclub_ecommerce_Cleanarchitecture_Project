package interfaces

import (
	"project/pkg/utils/domain"
	"project/pkg/utils/models"
)

type OrderRepository interface {
	OrderItems(userID, address int, TotalCartPrice float64) (int, error)
	AddOrderProducts(orderID int, cart []models.GetCart) error
	GetOrders(userID int) ([]domain.Order, error)
	GetOrderImages(orderID int) ([]string, error)
	GetOrderAddress(orderID int) (domain.Address, models.OrderData, error)
	CheckOrderStatusByID(orderID int) (string, error)
	CancelOrder(orderID int) error
	ReturnOrder(orderID int)error
}
