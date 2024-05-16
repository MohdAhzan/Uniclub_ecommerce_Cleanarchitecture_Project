package interfaces

import (
	"project/pkg/utils/domain"
	"project/pkg/utils/models"

	"github.com/jung-kurt/gofpdf"
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
	PrintInvoice(orderID, userID int) (*gofpdf.Fpdf, error)
}
