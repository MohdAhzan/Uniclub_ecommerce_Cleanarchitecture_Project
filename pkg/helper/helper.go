package helper

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"mime/multipart"
	cfg "project/pkg/config"
	"project/pkg/utils/models"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"

	"github.com/golang-jwt/jwt/v4"
	"github.com/twilio/twilio-go"
	openApi "github.com/twilio/twilio-go/rest/verify/v2"
	"golang.org/x/crypto/bcrypt"

	"net/smtp"
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
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 12)),
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
	//
	// refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	// refreshTokenString, err := refreshToken.SignedString([]byte("adminrefreshToken988243rwcfsdsjfyf74cysf38"))
	// if err != nil {
	// 	return "", "", nil
	// }

	return accessTokenString, nil

}

func (h *helper) GenerateTokenClients(user models.UserDetailsResponse) (string,string, error) {
	accessTokenClaims := &AuthCustomClaims{
		Id:    user.Id,
		Email: user.Email,
		Role:  "client",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 1)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	refreshTokenClaims := &AuthCustomClaims{
		Id:    user.Id,
		Email: user.Email,
		Role:  "client",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte("useraccesstokenasdioufou23854284jsdf9823jsdfh"))
	if err != nil {
		return "","", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte("userrefreshtokenasdgfr23788h23cy86qnw3dr367d4ye2"))
	if err != nil {
		return "", "", err
	}

	return accessTokenString,refreshTokenString ,nil

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

func (h *helper) AddImageToAwsS3(file *multipart.FileHeader) (string, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-southeast-2"))
	if err != nil {
		return "", err
	}

	client := s3.NewFromConfig(cfg)

	uploader := manager.NewUploader(client)
	f, openErr := file.Open()
	if openErr != nil {
		return "", openErr
	}
	defer f.Close()

	result, uploadErr := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String("jpeg123"),
		Key:         aws.String(file.Filename),
		Body:        f,
		ACL:         types.ObjectCannedACLPublicRead,
		ContentType: aws.String("image/png"),
	})

	if uploadErr != nil {
		fmt.Println("uploadERR", uploadErr)
		return "", uploadErr
	}

	return result.Location, nil
}

func (h *helper) SendMailToPhone(To, Subject, Msg string) error {

	TO := []string{To}

	//setup authentication
	auth := smtp.PlainAuth("", h.cfg.SMTP_USERNAME, h.cfg.SMTP_PASSWORD, h.cfg.SMTP_HOST)

	//message body
	msg := []byte("To: " + TO[0] + "\r\n" +
		"Subject: " + Subject + "\r\n" +
		"\r\n" +
		Msg + "\r\n")
	//send mail to recipient
	err := smtp.SendMail(h.cfg.SMTP_HOST+":"+h.cfg.SMTP_PORT, auth, h.cfg.SMTP_USERNAME, TO, msg)
	if err != nil {
		return err
	}
	return nil

}

func (h *helper) GenerateReferralCode() (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 12

	// Initialize the result string.
	result := make([]byte, length)

	// Generate a random index for each character in the result string.
	for i := range result {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[idx.Int64()]
	}

	return string(result), nil
}

func (h *helper) GetTimeFromPeriod(timePeriod string) (time.Time, time.Time) {

	endDate := time.Now()

	if timePeriod == "week" {
		startDate := endDate.AddDate(0, 0, -6)
		return startDate, endDate
	}

	if timePeriod == "month" {
		startDate := endDate.AddDate(0, -1, 0)
		return startDate, endDate
	}

	if timePeriod == "year" {
		startDate := endDate.AddDate(-1, 0, 0)
		return startDate, endDate
	}

	return endDate.AddDate(0, 0, -6), endDate

}

func (h *helper) ConvertToExel(sales []models.OrderDetailsAdmin) (*excelize.File, error) {

	filename := "../salesReport/sales_report.xlsx"
	file := excelize.NewFile()

	file.SetCellValue("Sheet1", "A1", "Product")
	file.SetCellValue("Sheet1", "B1", "Amount Sold")

	// Bold style for headings
	boldStyle, err := file.NewStyle(`{"font":{"bold":true}}`)
	if err != nil {
		return nil, err
	}

	file.SetCellStyle("Sheet1", "A1", "B1", boldStyle)

	var Total float64
	var Limit int
	for i, sale := range sales {
		col1 := fmt.Sprintf("A%d", i+2)
		col2 := fmt.Sprintf("B%d", i+2)

		file.SetCellValue("Sheet1", col1, sale.ProductName)
		file.SetCellValue("Sheet1", col2, sale.TotalAmount)
		Limit = i + 3
		Total += sale.TotalAmount

	}
	col1 := fmt.Sprintf("A%d", Limit)
	file.SetCellValue("Sheet1", col1, "Final Total")
	col2 := fmt.Sprintf("B%d", Limit)
	file.SetCellValue("Sheet1", col2, Total)

	// Larger font size for 'Final Total'
	largerFontStyle, err := file.NewStyle(`{"font":{"size":10}}`)
	if err != nil {
		return nil, err
	}
	file.SetCellStyle("Sheet1", col1, col2, largerFontStyle)

	if err := file.SaveAs(filename); err != nil {
		return nil, err
	}

	return file, nil

	// var Total float64
	// for i, sale := range sales {
	// 	col1 := fmt.Sprintf("A%d", i+2)
	// 	col2 := fmt.Sprintf("B%d", i+2)

	// 	file.SetCellValue("Sheet1", col1, sale.ProductName)
	// 	file.SetCellValue("Sheet1", col2, sale.TotalAmount)
	// 	Total += sale.TotalAmount
	// }

	// Limit := len(sales) + 2
	// col1 := fmt.Sprintf("A%d", Limit)
	// file.SetCellValue("Sheet1", col1, "Final Total")
	// col2 := fmt.Sprintf("B%d", Limit)
	// file.SetCellValue("Sheet1", col2, Total)

	// // Larger font size for 'Final Total'
	// largerFontStyle, err := file.NewStyle(`{"font":{"size":10}}`)
	// if err != nil {
	// 	return nil, err
	// }
	// file.SetCellStyle("Sheet1", col1, col2, largerFontStyle)

	// if err := file.SaveAs(filename); err != nil {
	// 	return nil, err
	// }

	// return file, nil

}
