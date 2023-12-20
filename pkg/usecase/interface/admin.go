package interfaces

import (
	"project/pkg/domain"
	"project/pkg/utils/models"
)

type AdminUseCase interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.TokenAdmin, error)
	GetUsers() ([]models.UserDetailsAtAdmin, error)
	BlockUser(id string) error
	UnBlockUser(id string) error
}
