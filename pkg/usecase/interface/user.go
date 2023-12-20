package interfaces

import (
	"project/pkg/utils/models"
)

type UserUseCase interface {
	UserSignup(user models.UserDetails) (models.TokenUsers, error)
	UserLoginHandler(user models.UserLogin) (models.TokenUsers, error)
}
