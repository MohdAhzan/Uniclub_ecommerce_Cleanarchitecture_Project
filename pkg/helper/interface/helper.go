package interfaces

import (
	"project/pkg/utils/models"
)

type Helper interface {
	PasswordHashing(string) (string, error)
	GenerateTokenClients(user models.UserDetailsResponse) (string, string, error)
	GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, string, error)
}
