package usecase

import (
	"errors"
	"fmt"
	helper "project/pkg/helper/interface"
	interfaces "project/pkg/repository/interface"
	services "project/pkg/usecase/interface"
	domain "project/pkg/utils/domain"
	"project/pkg/utils/models"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type adminUseCase struct {
	adminRepository interfaces.AdminRepository
	orderRepository interfaces.OrderRepository
	helper          helper.Helper
}

func NewAdminUsecase(repo interfaces.AdminRepository, h helper.Helper, o interfaces.OrderRepository) services.AdminUseCase {
	return &adminUseCase{
		adminRepository: repo,
		helper:          h,
		orderRepository: o,
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
		return []models.UserDetailsAtAdmin{}, err
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

func (ad *adminUseCase) OrderReturnApprove(orderID int) error {

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

	mailSub := "Order Return Approval"
	mailMsg := "Hey " + user.Name + "...your request for Order Return has been approved!!!"
	err = ad.helper.SendMailToPhone(user.Email, mailSub, mailMsg)
	if err != nil {
		return err
	}
	return nil

}
