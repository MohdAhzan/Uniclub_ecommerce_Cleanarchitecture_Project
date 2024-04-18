package interfaces

import (
	"mime/multipart"
	"project/pkg/utils/models"
)

type Helper interface {
	PasswordHashing(string) (string, error)
	GenerateTokenClients(user models.UserDetailsResponse) (string, error)
	GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, error)
	CompareHashAndPassword(hashedPassword string, password string) error

	TwilioSetup(accountSID string, authToken string)
	TwilioSendOTP(phoneNo string, serviceSID string) (string, error)
	TwilioVerifyOTP(serviceSID string, code string, phoneNo string) error

	AddImageToAwsS3(file *multipart.FileHeader) (string, error)

	SendMailToPhone(To, Subject, Msg string) error

	// CacheManage(key string,)
}
