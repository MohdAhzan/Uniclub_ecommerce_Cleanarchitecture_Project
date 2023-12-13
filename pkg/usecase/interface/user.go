package interfaces

import (
	"project/pkg/utils/models"
)

type UserUseCase interface {
	UserSignup(user models.UserDetails) (models.TokenUsers, error)
}
