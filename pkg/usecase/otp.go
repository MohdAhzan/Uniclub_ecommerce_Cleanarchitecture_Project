package usecase

import (
	"errors"
	"fmt"
	"project/pkg/config"
	helper_interface "project/pkg/helper/interface"
	interfaces "project/pkg/repository/interface"
	services "project/pkg/usecase/interface"
	"project/pkg/utils/models"

	"github.com/jinzhu/copier"
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

func (otp *otpUseCase) SendOTP(phoneNo string) error {

	ok := otp.otpRepository.FindUserByMobileNumber(phoneNo)
	if !ok {
		return errors.New("the user does not exist")
	}
	otp.helper.TwilioSetup(otp.cfg.DBACCOUNTSID, otp.cfg.DBAUTHTOKEN)
	_, err := otp.helper.TwilioSendOTP(phoneNo, otp.cfg.DBSERVICESID)

	if err != nil {

		return err

	}

	return nil
}

func (otp *otpUseCase) VerifyOTP(code models.VerifyData) (models.TokenUsers, error) {
	otp.helper.TwilioSetup(otp.cfg.DBACCOUNTSID, otp.cfg.DBAUTHTOKEN)
	fmt.Printf("the otp recieved is %v \n and phone number is %v", code.Code, code.PhoneNumber)

	err := otp.helper.TwilioVerifyOTP(otp.cfg.DBSERVICESID, code.Code, code.PhoneNumber)
	if err != nil {
		return models.TokenUsers{}, errors.New("error while verifying")
	}
	// if user is authenticated using OTP send back user details
	userDetails, err := otp.otpRepository.UserDetailsUsingPhone(code.PhoneNumber)
	if err != nil {
		return models.TokenUsers{}, err
	}

	tokenString,refreshtokenString, err := otp.helper.GenerateTokenClients(userDetails)

	if err != nil {
		return models.TokenUsers{}, err
	}

	var user models.UserDetailsResponse

	err = copier.Copy(&user, &userDetails)
	if err != nil {
		return models.TokenUsers{}, err
	}

	return models.TokenUsers{
		Users: user,
		AccessToken: tokenString,
		RefreshToken: refreshtokenString,
	}, nil

}
