package interfaces

import (
	"project/pkg/utils/domain"
	"project/pkg/utils/models"
)

type OrderRepository interface {
	OrderItems(userID, address, paymentID, couponID int, TotalCartPrice float64) (int, error)
	AddOrderProducts(orderID int, cart []models.GetCart) error

	GetOrders(userID int) ([]domain.Order, error)
	GetOrderImages(orderID int) ([]string, error)
	GetOrderProductIDs(orderID int) ([]int, error)
	GetOrderAddress(orderID int) (domain.Address, models.OrderData, error)
	CheckOrderStatusByID(orderID int) (string, error)
	CancelOrder(orderID int) error
	ReturnOrder(orderID int) error
	CheckOrderByID(orderID int) error
	EditOrderStatus(order int, status string) error
	FindOrderAmount(orderID int) (float64, error)
	FindOrderedUserID(orderID int) (int, error)
	GetPaymentMethodsByID(PaymentMethodID int) (string, error)
	GetPaymentStatusByID(orderID int) (string, error)
	GetAllOrderItemsByOrderID(orderID int) ([]domain.EachProductOrderDetails, error)
	CheckIndividualOrders(orderID, pID int) (int, error)
	CheckOrderCount(orderID int) (int, error)
	DeleteProductInOrder(orderID, pID int) (float64, error)
	UpdateFinalOrderPrice(orderID int, NewPrice float64) error
	GetOrderProductNames(pID int) (string, error)
}
