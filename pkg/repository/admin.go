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

func (adm *adminRepository) MakePaymentStatusAsPaid(orderID int) error {

	if err := adm.db.Exec(`UPDATE orders SET payment_status = 'PAID',updated_at = ? WHERE id = ?`, time.Now(), orderID).Error; err != nil {
		return err
	}
	return nil
}

func (ad *adminRepository) GetAllOrderDetailsByStatus() (domain.AdminOrdersResponse, error) {

	var orderData domain.AdminOrdersResponse

	var pending []domain.AdminOrderDetails

	err := ad.db.Raw(`SELECT o.id, o.created_at, o.updated_at,u.name, a.*,o.payment_method, o.price as total , o.order_status, o.payment_status 
FROM 
    orders o 
JOIN 
    users u ON o.user_id = u.id 
JOIN 
    addresses a ON o.address_id = a.id 
WHERE 
    o.order_status = 'PENDING'`).Scan(&pending).Error
	fmt.Println(pending, "pending orders.............................................")
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	var shipped []domain.AdminOrderDetails

	err = ad.db.Raw(`SELECT o.id, o.created_at, o.updated_at,u.name, a.*,o.payment_method, o.price as total, o.order_status, o.payment_status 
	FROM 
		orders o 
	JOIN 
		users u ON o.user_id = u.id 
	JOIN 
		addresses a ON o.address_id = a.id 
	WHERE 
		o.order_status = 'SHIPPED' `).Scan(&shipped).Error
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	var delivered []domain.AdminOrderDetails

	err = ad.db.Raw(`SELECT o.id, o.created_at, o.updated_at,u.name, a.*,o.payment_method, o.price as total, o.order_status, o.payment_status 
	FROM 
		orders o 
	JOIN 
		users u ON o.user_id = u.id 
	JOIN 
		addresses a ON o.address_id = a.id 
	WHERE 
		o.order_status = 'DELIVERED'`).Scan(&delivered).Error
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	var cancelled []domain.AdminOrderDetails

	err = ad.db.Raw(`SELECT o.id, o.created_at, o.updated_at,u.name, a.*,o.payment_method, o.price, o.order_status, o.payment_status 
	FROM 
		orders o 
	JOIN 
		users u ON o.user_id = u.id 
	JOIN 
		addresses a ON o.address_id = a.id 
	WHERE 
		o.order_status = 'CANCELED'`).Scan(&cancelled).Error
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	var return_requested []domain.AdminOrderDetails

	err = ad.db.Raw(`SELECT o.id, o.created_at, o.updated_at,u.name , a.*,o.payment_method, o.price, o.order_status, o.payment_status 
	FROM 
		orders o 
	JOIN 
		users u ON o.user_id = u.id 
	JOIN 
		addresses a ON o.address_id = a.id 
	WHERE 
		o.order_status ='RETURN_REQUESTED'`).Scan(&return_requested).Error
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	var returned []domain.AdminOrderDetails

	err = ad.db.Raw(`SELECT o.id, o.created_at, o.updated_at,u.name , a.*,o.payment_method, o.price, o.order_status, o.payment_status 
	FROM 
		orders o 
	JOIN 
		users u ON o.user_id = u.id 
	JOIN 
		addresses a ON o.address_id = a.id 
	WHERE 
		o.order_status = 'RETURNED'`).Scan(&returned).Error
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	orderData.PENDING = pending
	orderData.SHIPPED = shipped
	orderData.DELIVERED = delivered
	orderData.CANCELED = cancelled
	orderData.RETURN_REQUESTED = return_requested
	orderData.RETURNED = returned

	return orderData, nil

}
