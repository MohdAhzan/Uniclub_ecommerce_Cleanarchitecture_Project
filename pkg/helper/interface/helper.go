package interfaces

import (
	"mime/multipart"
	"github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/models"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type Helper interface {
	PasswordHashing(string) (string, error)
	GenerateTokenClients(user models.UserDetailsResponse) (string, string, error)
	GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, error)
	CompareHashAndPassword(hashedPassword string, password string) error

	TwilioSetup(accountSID string, authToken string)
	TwilioSendOTP(phoneNo string, serviceSID string) (string, error)
	TwilioVerifyOTP(serviceSID string, code string, phoneNo string) error

	AddImageToAwsS3(file *multipart.FileHeader) (string, error)

	SendMailToPhone(To, Subject, Msg string) error

	GenerateReferralCode() (string, error)
	ConvertToExel(sales []models.OrderDetailsAdmin) (*excelize.File, error)
	GetTimeFromPeriod(timePeriod string) (time.Time, time.Time)
}
