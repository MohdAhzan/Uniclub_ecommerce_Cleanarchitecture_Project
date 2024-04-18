package interfaces

import (
	"project/pkg/utils/domain"
	"project/pkg/utils/models"
)

type AdminUseCase interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.TokenAdmin, error)
	GetUsers() ([]models.UserDetailsAtAdmin, error)
	BlockUser(id int) error
	UnBlockUser(id int) error
	OrderReturnApprove(orderID int) error
}
