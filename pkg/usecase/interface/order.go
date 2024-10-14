package interfaces

import (
	"github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/domain"
	"github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/models"
)

type OrderUseCase interface {
	OrderFromCart(order models.Order, couponID int) error
	Checkout(userID, couponID int) (models.CheckOut, error)
	GetOrders(id int) ([]domain.OrderDetailsWithImages, error)
	GetOrderDetailsByOrderID(orderID, userID int) (domain.OrderDetails, error)
	CancelOrder(orderID, userID int) error
	ReturnOrder(orderID, userID int) error
	GetEachProductOrderDetails(orderID, userID int) (domain.OrderDetailsSeparate, error)
	CancelProductInOrder(orderID, pID, user_id int) (domain.OrderDetailsSeparate, error)
}
