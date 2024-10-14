package usecase

import (
	"errors"
	"fmt"
	helper "github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/helper/interface"
	interfaces "github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/repository/interface"
	services "github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/usecase/interface"
	domain "github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/domain"
	"github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/models"
	"strconv"
	"time"

	"github.com/jinzhu/copier"
	"github.com/jung-kurt/gofpdf"
	"golang.org/x/crypto/bcrypt"
)

type adminUseCase struct {
	adminRepository interfaces.AdminRepository
	orderRepository interfaces.OrderRepository
	userRepository  interfaces.UserRepository
	helper          helper.Helper
}

func NewAdminUsecase(repo interfaces.AdminRepository, h helper.Helper, o interfaces.OrderRepository, u interfaces.UserRepository) services.AdminUseCase {
	return &adminUseCase{
		adminRepository: repo,
		helper:          h,
		orderRepository: o,
		userRepository:  u,
	}
}

func (ad *adminUseCase) LoginHandler(adminDetails models.AdminLogin) (domain.TokenAdmin, error) {

	adminCompareDetails, err := ad.adminRepository.LoginHandler(adminDetails)
	if err != nil {
		return domain.TokenAdmin{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(adminCompareDetails.Password), []byte(adminDetails.Password))
	if err != nil {
		return domain.TokenAdmin{}, err
	}
	var adminDetailsResponse models.AdminDetailsResponse

	err = copier.Copy(&adminDetailsResponse, &adminCompareDetails)
	if err != nil {
		return domain.TokenAdmin{}, err
	}

	access, err := ad.helper.GenerateTokenAdmin(adminDetailsResponse)

	if err != nil {
		return domain.TokenAdmin{}, err
	}

	return domain.TokenAdmin{
		Admin:       adminDetailsResponse,
		AccessToken: access,
	}, nil
}

func (ad *adminUseCase) GetUsers() ([]models.UserDetailsAtAdmin, error) {
	users, err := ad.adminRepository.GetUsers()
	if err != nil {
		return []models.UserDetailsAtAdmin{}, errors.New("Error fetching UserDetails")
	}
	return users, nil
}

func (ad *adminUseCase) BlockUser(id int) error {
	user, err := ad.adminRepository.GetUserByID(id)
	if err != nil {
		return err
	}

	if user.Blocked {
		return errors.New("user already blocked")
	} else {
		user.Blocked = true
	}

	err = ad.adminRepository.UpdateBlockUserByID(user)
	if err != nil {
		return err
	}

	return nil
}

func (ad *adminUseCase) UnBlockUser(id int) error {
	user, err := ad.adminRepository.GetUserByID(id)
	if err != nil {
		return err
	}
	if user.Blocked {
		user.Blocked = false
	} else {
		return errors.New("user already unblocked")
	}

	err = ad.adminRepository.UpdateBlockUserByID(user)
	if err != nil {
		return err
	}
	return nil
}

func (ad *adminUseCase) EditOrderStatus(orderID int, status string) error {

	err := ad.orderRepository.CheckOrderByID(orderID)
	if err != nil {
		return err
	}

	err = ad.orderRepository.EditOrderStatus(orderID, status)
	if err != nil {
		return err
	}

	return nil

}

func (ad *adminUseCase) MakePaymentStatusAsPaid(orderID int) error {

	err := ad.orderRepository.CheckOrderByID(orderID)
	if err != nil {
		return err
	}

	err = ad.adminRepository.MakePaymentStatusAsPaid(orderID)
	if err != nil {

		return err
	}

	return nil
}

func (ad *adminUseCase) OrderReturnApprove(orderID int) error {

	err := ad.orderRepository.CheckOrderByID(orderID)
	if err != nil {
		return err
	}

	status, err := ad.orderRepository.CheckOrderStatusByID(orderID)
	if err != nil {
		return err
	}

	userID, err := ad.adminRepository.GetUserIDbyorderID(orderID)
	if err != nil {
		return err
	}
	user, err := ad.adminRepository.GetUserByID(userID)
	if err != nil {
		return err
	}

	msg := fmt.Sprintf("User : %s has not requested to return the product", user.Name)
	if status == "RETURNED" {
		return fmt.Errorf("user :%s has already returned the product", user.Name)

	} else if status != "RETURN_REQUESTED" {
		return fmt.Errorf(msg)
	}

	err = ad.adminRepository.OrderReturnApprove(orderID)
	if err != nil {
		return err
	}

	orderAmount, err := ad.orderRepository.FindOrderAmount(orderID)
	if err != nil {
		return err
	}
	var model models.AddMoneytoWallet

	model.UserID = userID
	model.Amount = orderAmount
	model.TranscationType = "PDT_RETURNED"

	err = ad.userRepository.AddMoneytoWallet(model)
	if err != nil {
		return err
	}

	mailSub := "Order Return Approval "
	mailMsg := "Hey " + user.Name + "...your request for Order Return has been approved and Order Amount is transfered to your Wallet!!! "
	err = ad.helper.SendMailToPhone(user.Email, mailSub, mailMsg)
	if err != nil {
		return err
	}
	return nil

}

func (ad *adminUseCase) GetAllOrderDetails() (domain.AdminOrdersResponse, error) {

	orders, err := ad.adminRepository.GetAllOrderDetailsByStatus()
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	return orders, nil
}

func (ad *adminUseCase) NewPaymentMethod(pMethod string) error {

	err := ad.adminRepository.AddNewPaymentMethod(pMethod)
	if err != nil {
		return err
	}

	return nil
}

func (ad *adminUseCase) GetAllPaymentMethods() ([]models.GetPaymentMethods, error) {

	data, err := ad.adminRepository.GetAllPaymentMethods()
	if err != nil {
		return []models.GetPaymentMethods{}, err
	}

	return data, nil
}

func (ad *adminUseCase) DeletePaymentMethod(paymentID int) error {

	err := ad.adminRepository.DeletePaymentMethod(paymentID)
	if err != nil {
		return err
	}
	return nil

}

func (ah *adminUseCase) FilteredSalesReport(timePeriod string) (models.SalesReport, error) {
	if timePeriod == "" {
		err := errors.New("please provide a timePeriod")
		return models.SalesReport{}, err
	}

	if timePeriod != "week" && timePeriod != "month" && timePeriod != "year" {
		err := errors.New("invalid timePeriod ")
		return models.SalesReport{}, err
	}

	startTime, endTime := ah.helper.GetTimeFromPeriod(timePeriod)
	fmt.Println("starttime", startTime)
	fmt.Println("ENDDTIMEEE", endTime)
	saleReport, err := ah.adminRepository.FilteredSalesReport(startTime, endTime)

	if err != nil {
		return models.SalesReport{}, err
	}
	return saleReport, nil
}

func (ad *adminUseCase) SalesByDate(dayInt int, monthInt int, yearInt int) ([]models.OrderDetailsAdmin, error) {

	if dayInt == 0 && monthInt == 0 && yearInt == 0 {
		return []models.OrderDetailsAdmin{}, errors.New("day, month, and year field are empty ")
	}

	if dayInt < 0 || monthInt < 0 || yearInt < 0 {
		return []models.OrderDetailsAdmin{}, errors.New("enter valid day, month, and year")
	}

	if yearInt >= 2020 {
		if monthInt == 0 && dayInt == 0 {

			body, err := ad.adminRepository.SalesByYear(yearInt)
			if err != nil {
				return []models.OrderDetailsAdmin{}, err
			}

			return body, nil
		} else if monthInt > 0 && monthInt <= 12 && dayInt == 0 {

			body, err := ad.adminRepository.SalesByMonth(yearInt, monthInt)
			if err != nil {
				return []models.OrderDetailsAdmin{}, err
			}

			return body, nil
		} else if monthInt > 0 && monthInt <= 12 && dayInt > 0 && dayInt <= 31 {

			body, err := ad.adminRepository.SalesByDay(yearInt, monthInt, dayInt)
			if err != nil {
				return []models.OrderDetailsAdmin{}, err
			}

			return body, nil
		}
	}

	return []models.OrderDetailsAdmin{}, errors.New("invalid format")
}

func (ad *adminUseCase) PrintSalesReport(sales []models.OrderDetailsAdmin) (*gofpdf.Fpdf, error) {

	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.AddPage()
	pdf.SetFont("Arial", "B", 22)
	pdf.SetTextColor(31, 73, 125)
	pdf.CellFormat(0, 20, "Total Sales Report", "0", 1, "C", false, 0, "")

	pdf.SetFont("Arial", "", 16)
	pdf.SetTextColor(0, 0, 0)
	var FinalAmount float64
	for _, item := range sales {
		pdf.CellFormat(0, 10, "Product: "+item.ProductName, "0", 1, "L", false, 0, "")
		amount := strconv.FormatFloat(item.TotalAmount, 'f', 2, 64)
		pdf.CellFormat(0, 10, "Amount Sold: $"+amount, "0", 1, "L", false, 0, "")
		pdf.Ln(5)
		FinalAmount += item.TotalAmount
	}
	pdf.SetFont("Arial", "", 18)
	pdf.SetTextColor(0, 0, 0)
	FinalTotal := strconv.FormatFloat(FinalAmount, 'f', 2, 64)
	pdf.CellFormat(0, 10, " Total Amount Sold: "+FinalTotal, "0", 1, "L", false, 0, "")

	pdf.SetFont("Arial", "I", 12)
	pdf.SetTextColor(150, 150, 150)

	pdf.Cell(0, 10, "Generated by Uniclub India Pvt Ltd. - "+time.Now().Format("2006-01-02 15:04:05"))

	return pdf, nil
}

func (ad *adminUseCase) ChangePassword(changePasswordDetails models.AdminPasswordChange, id int) error {

	if changePasswordDetails.NewPassword == changePasswordDetails.CurrentPassword {

		return fmt.Errorf("New Password is same as old one.Try again!!!")
	}

	hashedPass, err := ad.adminRepository.GetAdminHashPassword(id)
	if err != nil {
		return err
	}

	fmt.Println("hashed PAss", hashedPass)

	err = ad.helper.CompareHashAndPassword(hashedPass, changePasswordDetails.CurrentPassword)
	fmt.Println(changePasswordDetails.CurrentPassword, "currentPass")

	if err != nil {
		return fmt.Errorf("Incorrect Password Try again!!!")

	}

	if len(changePasswordDetails.NewPassword) < 3 {

		return fmt.Errorf("Password is too short")
	}

	if changePasswordDetails.NewPassword != changePasswordDetails.ConfirmPassword {

		return fmt.Errorf("Password mismatch Try again!!!")

	}

	newHashedPass, err := ad.helper.PasswordHashing(changePasswordDetails.NewPassword)
	if err != nil {
		return err
	}

	err = ad.adminRepository.UpdateAdminPass(id, newHashedPass)
	if err != nil {
		return err
	}

	return nil
}
