package interfaces

import "github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/models"

type OtpUseCase interface {
	SendOTP(phoneNo string) error
	VerifyOTP(code models.VerifyData) (models.TokenUsers, error)
}
