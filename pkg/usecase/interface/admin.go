package interfaces

import (
	"project/pkg/utils/domain"
	"project/pkg/utils/models"

	"github.com/jung-kurt/gofpdf"
)

type AdminUseCase interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.TokenAdmin, error)
	GetUsers() ([]models.UserDetailsAtAdmin, error)
	BlockUser(id int) error
	UnBlockUser(id int) error
	OrderReturnApprove(orderID int) error
	EditOrderStatus(orderID int, status string) error
	MakePaymentStatusAsPaid(orderID int) error
	GetAllOrderDetails() (domain.AdminOrdersResponse, error)
	NewPaymentMethod(pMethod string) error
	GetAllPaymentMethods() ([]models.GetPaymentMethods, error)
	DeletePaymentMethod(paymentID int) error
	FilteredSalesReport(timePeriod string) (models.SalesReport, error)
	PrintSalesReport(sales []models.OrderDetailsAdmin) (*gofpdf.Fpdf, error)
	SalesByDate(dayInt int, monthInt int, yearInt int) ([]models.OrderDetailsAdmin, error)
}
