package helper

import (
	"errors"
	"fmt"
	cfg "project/pkg/config"
	"project/pkg/utils/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/twilio/twilio-go"
	openApi "github.com/twilio/twilio-go/rest/verify/v2"
	"golang.org/x/crypto/bcrypt"
)

type helper struct {
	cfg cfg.Config
}

type AuthCustomClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func NewHelper(config cfg.Config) *helper {
	return &helper{
		cfg: config,
	}
}

func (h *helper) PasswordHashing(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", errors.New("internal server error")
	}

	hash := string(hashedPassword)
	return hash, nil
}

func (h *helper) GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, error) {
	accessTokenClaims := &AuthCustomClaims{
		Id:    admin.ID,
		Email: admin.Email,
		Role:  "admin",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 20)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte("adminaccesstokena983274uhweirbt"))
	if err != nil {
		return "", err
	}

	// refreshTokenClaims := &AuthCustomClaims{
	// 	Id:    admin.ID,
	// 	Email: admin.Email,
	// 	Role:  "admin",
	// 	RegisteredClaims: jwt.RegisteredClaims{
	// 		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 20)),
	// 		IssuedAt:  jwt.NewNumericDate(time.Now()),
	// 	},
	// }

	// refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	// refreshTokenString, err := refreshToken.SignedString([]byte("adminrefreshToken988243rwcfsdsjfyf74cysf38"))
	// if err != nil {
	// 	return "", "", nil
	// }

	return accessTokenString, nil

}

func (h *helper) GenerateTokenClients(user models.UserDetailsResponse) (string, error) {
	accessTokenClaims := &AuthCustomClaims{
		Id:    user.Id,
		Email: user.Email,
		Role:  "client",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 20)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// refreshTokenClaims := &AuthCustomClaims{
	// 	Id:    user.Id,
	// 	Email: user.Email,
	// 	Role:  "client",
	// 	RegisteredClaims: jwt.RegisteredClaims{
	// 		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
	// 		IssuedAt:  jwt.NewNumericDate(time.Now()),
	// 	},
	// }

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte("useraccesstokenasdioufou23854284jsdf9823jsdfh"))
	if err != nil {
		return "", err
	}

	// refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	// refreshTokenString, err := refreshToken.SignedString([]byte("userrefreshtokenasdgfr23788h23cy86qnw3dr367d4ye2"))
	// if err != nil {
	// 	return "", "", err
	// }

	return accessTokenString, nil

}

func (h *helper) CompareHashAndPassword(hashPass string, pass string) error {

	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(pass))

	if err != nil {
		return err
	}

	return nil

}

var client *twilio.RestClient

func (h *helper) TwilioSetup(accountSID string, authToken string) {

	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSID,
		Password: authToken,
	})
}

func (h *helper) TwilioSendOTP(phoneNo string, serviceSID string) (string, error) {
	// fmt.Println("phone no is =", phoneNo, "     and servicesid is =", serviceSID)

	to := "+91" + phoneNo
	params := &openApi.CreateVerificationParams{}
	params.SetTo(to)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(serviceSID, params)
	// fmt.Println("VErificatoino Params", params)
	if err != nil {

		return " ", err
	}
	fmt.Println("verificatoin SID", *resp.Sid)
	return *resp.Sid, nil

}

func (h *helper) TwilioVerifyOTP(serviceSID string, code string, phoneNo string) error {

	params := &openApi.CreateVerificationCheckParams{}
	params.SetTo("+91" + phoneNo)
	params.SetCode(code)
	resp, err := client.VerifyV2.CreateVerificationCheck(serviceSID, params)

	if err != nil {
		fmt.Println("ERRORR is", err)
		return err
	}

	if *resp.Status == "approved" {
		return nil
	}

	return errors.New("failed to validate otp")

}
