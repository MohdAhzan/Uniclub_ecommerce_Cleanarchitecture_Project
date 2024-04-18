package repository

import (
	"errors"
	"fmt"
	interfaces "project/pkg/repository/interface"
	"project/pkg/utils/domain"
	"project/pkg/utils/models"
	"time"

	"gorm.io/gorm"
)

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {

	return &adminRepository{
		db: DB,
	}

}

func (ad *adminRepository) LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error) {

	var adminCompareDetails domain.Admin

	if err := ad.db.Raw("select * from admins where email = ? ", adminDetails.Email).Scan(&adminCompareDetails).Error; err != nil {
		return domain.Admin{}, err
	}

	return adminCompareDetails, nil
}

func (ad *adminRepository) GetUsers() ([]models.UserDetailsAtAdmin, error) {

	var count int
	if err := ad.db.Raw("select count(*) from users").Scan(&count).Error; err != nil {
		return []models.UserDetailsAtAdmin{}, err
	}
	if count < 1 {
		return []models.UserDetailsAtAdmin{}, errors.New("empty users in database")
	}

	var userDetails []models.UserDetailsAtAdmin

	if err := ad.db.Raw("select id,name,email,phone,blocked from users").Scan(&userDetails).Error; err != nil {
		return []models.UserDetailsAtAdmin{}, err
	}

	return userDetails, nil

}

func (ad *adminRepository) GetUserByID(userID int) (domain.Users, error) {

	var count int

	if err := ad.db.Raw("select count(*) from users where id = ?", userID).Scan(&count).Error; err != nil {
		return domain.Users{}, err
	}

	if count < 1 {
		return domain.Users{}, errors.New("user for the given id doesn't exists")
	}

	query := fmt.Sprintf("select * from users where id = '%d'", userID)

	var userDetails domain.Users

	if err := ad.db.Raw(query).Scan(&userDetails).Error; err != nil {
		return domain.Users{}, err
	}
	return userDetails, nil
}

//blockes and unblockes users

func (ad *adminRepository) UpdateBlockUserByID(user domain.Users) error {

	fmt.Println("now id =", user.ID)
	if err := ad.db.Exec("update users set blocked = ? where id = ?", user.Blocked, user.ID).Error; err != nil {
		return err
	}
	return nil
}

func (ad *adminRepository) OrderReturnApprove(orderID int) error {

	if err := ad.db.Exec(`UPDATE orders SET order_status = 'RETURNED',updated_at = ? WHERE id = ?`, time.Now(), orderID).Error; err != nil {
		return err
	}
	return nil
}

func (ad *adminRepository) GetUserIDbyorderID(orderID int) (int, error) {

	var userID int

	err := ad.db.Raw("select user_id from orders where id=?", orderID).Scan(&userID).Error
	if err != nil {
		return 0, err
	}
	return userID, nil
}
