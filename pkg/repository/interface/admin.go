package interfaces

import (
	"project/pkg/domain"
	"project/pkg/utils/models"
)

type AdminRepository interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error)
}
