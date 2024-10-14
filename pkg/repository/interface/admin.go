package interfaces

import (
	"github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/domain"
	"github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/models"
	"time"
)

type AdminRepository interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error)
	GetUsers() ([]models.UserDetailsAtAdmin, error)
	GetUserByID(id int) (domain.Users, error)
	UpdateBlockUserByID(user domain.Users) error
	OrderReturnApprove(orderID int) error
	GetUserIDbyorderID(orderID int) (int, error)
	MakePaymentStatusAsPaid(orderID int) error
	GetAllOrderDetailsByStatus() (domain.AdminOrdersResponse, error)
	AddNewPaymentMethod(pMethod string) error
	GetAllPaymentMethods() ([]models.GetPaymentMethods, error)
	DeletePaymentMethod(paymentID int) error
	FilteredSalesReport(startTime time.Time, endTime time.Time) (models.SalesReport, error)
	SalesByDay(yearInt int, monthInt int, dayInt int) ([]models.OrderDetailsAdmin, error)
	SalesByMonth(yearInt int, monthInt int) ([]models.OrderDetailsAdmin, error)
	SalesByYear(yearInt int) ([]models.OrderDetailsAdmin, error)
	GetAdminHashPassword(id int) (string, error)
	UpdateAdminPass(id int, NewPass string) error
}
