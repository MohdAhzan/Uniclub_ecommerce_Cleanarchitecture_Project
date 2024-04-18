package interfaces

import (
	"project/pkg/utils/domain"
	"project/pkg/utils/models"
)

type AdminRepository interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error)
	GetUsers() ([]models.UserDetailsAtAdmin, error)
	GetUserByID(id int) (domain.Users, error)
	UpdateBlockUserByID(user domain.Users) error
	OrderReturnApprove(orderID int) error
	GetUserIDbyorderID(orderID int) (int, error)
}
