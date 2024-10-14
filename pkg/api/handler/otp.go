package handler

import (
	"net/http"
	services "github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/usecase/interface"
	response "github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/Response"
	models "github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/models"

	"github.com/gin-gonic/gin"
)

type OtpHandler struct {
	otpUseCase services.OtpUseCase
}

func NewOtpHandler(useCase services.OtpUseCase) *OtpHandler {
	return &OtpHandler{
		otpUseCase: useCase,
	}
}

func (otp *OtpHandler) SendOTPHandler(c *gin.Context) {
	var phone models.OTPData
	if err := c.BindJSON(&phone); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
	}

	err := otp.otpUseCase.SendOTP(phone.PhoneNumber)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not send OTP", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "OTP sent successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

func (otp *OtpHandler) VerifyOTPHandler(c *gin.Context) {
	var code models.VerifyData
	if err := c.BindJSON(&code); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
	}
	users, err := otp.otpUseCase.VerifyOTP(code)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not verify OTP", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Succesfully verified OTP", users, nil)
	c.JSON(http.StatusOK, successRes)
}
