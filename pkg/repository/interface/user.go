package interfaces

import "project/pkg/utils/models"

type UserRepository interface {
	UserSignup(user models.UserDetails) (models.UserDetailsResponse, error)
	CheckUserAvailability(email string) bool
	FindUserByEmail(user models.UserLogin) (models.UserSignInResponse, error)
	UserBlockStatus(email string) (bool, error)
}
