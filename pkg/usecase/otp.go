package usecase

import (
	"project/pkg/config"
	helper_interface "project/pkg/helper/interface"
	interfaces "project/pkg/repository/interface"
	services "project/pkg/usecase/interface"
)

type otpUseCase struct {
	cfg           config.Config
	otpRepository interfaces.OtpRepository
	helper        helper_interface.Helper
}

func NewOtpUseCase(cfg config.Config, otpRepo interfaces.OtpRepository, helper helper_interface.Helper) services.OtpUseCase {

	return &otpUseCase{
		cfg:           cfg,
		otpRepository: otpRepo,
		helper:        helper,
	}
}
