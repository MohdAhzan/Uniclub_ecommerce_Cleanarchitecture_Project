package interfaces

import "project/pkg/utils/models"

type OtpUseCase interface {
	SendOTP(phoneNo string) error
	VerifyOTP(code models.VerifyData) (models.TokenUsers, error)
}
