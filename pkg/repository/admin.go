package repository

import (
	"errors"
	"fmt"
	interfaces "github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/repository/interface"
	"github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/domain"
	"github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/models"
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

func (ad *adminRepository) AddNewPaymentMethod(pMethod string) error {

	err := ad.db.Exec("INSERT INTO payment_methods (payment_name) VALUES(?)", pMethod).Error
	if err != nil {
		return err
	}
	return nil
}

func (ad *adminRepository) GetAllPaymentMethods() ([]models.GetPaymentMethods, error) {

	var data []models.GetPaymentMethods

	err := ad.db.Raw("select id,payment_name from payment_methods").Scan(&data).Error
	if err != nil {
		return []models.GetPaymentMethods{}, err
	}

	return data, nil
}

func (ad *adminRepository) DeletePaymentMethod(paymentID int) error {

	result := ad.db.Exec(`DELETE FROM payment_methods where id = ?`, paymentID)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return errors.New("no payment Method in this id")
	}
	return nil
}

func (ad *adminRepository) FilteredSalesReport(startTime time.Time, endTime time.Time) (models.SalesReport, error) {
	var salesReport models.SalesReport
	querry := `
		SELECT COALESCE(SUM(final_price),0) 
		FROM orders WHERE payment_status='PAID'
		AND created_at >= ? AND created_at <= ?
		`
	result := ad.db.Raw(querry, startTime, endTime).Scan(&salesReport.TotalSalesAmount)
	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}

	result = ad.db.Raw("SELECT COUNT(*) FROM orders where created_at >= ? AND created_at <= ?", startTime, endTime).Scan(&salesReport.TotalOrders)
	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}

	querry = `
		SELECT COUNT(*) FROM orders 
		WHERE payment_status = 'PAID' and 
		created_at >= ? AND created_at <= ?
		`

	result = ad.db.Raw(querry, startTime, endTime).Scan(&salesReport.CompletedOrders)
	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}

	querry = `
		SELECT COUNT(*) FROM orders WHERE 
		order_status = 'PENDING' AND created_at >= ? AND created_at<=?
		`
	result = ad.db.Raw(querry, startTime, endTime).Scan(&salesReport.PendingOrders)
	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}
	querry = `
		SELECT COUNT(*) FROM orders WHERE 
		order_status = 'CANCELED' AND created_at >= ? AND created_at<=?
		`
	result = ad.db.Raw(querry, startTime, endTime).Scan(&salesReport.CancelledOrders)
	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}
	querry = `
		SELECT COUNT(*) FROM orders WHERE 
		order_status = 'RETURNED' AND created_at >= ? AND created_at<=?
		`
	result = ad.db.Raw(querry, startTime, endTime).Scan(&salesReport.ReturnedOrders)
	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}

	var inventoryID int
	querry = `
		SELECT inventory_id FROM order_items 
		GROUP BY inventory_id order by SUM(quantity) DESC LIMIT 1
		`
	result = ad.db.Raw(querry).Scan(&inventoryID)
	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}

	result = ad.db.Raw("SELECT product_name FROM inventories WHERE product_id = ?", inventoryID).Scan(&salesReport.MostSelledProduct)
	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}
	return salesReport, nil
}

func (ad *adminRepository) SalesByYear(yearInt int) ([]models.OrderDetailsAdmin, error) {
	var orderDetails []models.OrderDetailsAdmin

	query := `SELECT i.product_name, SUM(oi.total_price) AS total_amount
	FROM orders o
	JOIN order_items oi ON o.id = oi.order_id
	JOIN inventories i ON oi.inventory_id = i.product_id
	WHERE o.payment_status = 'PAID'
	AND EXTRACT(YEAR FROM o.created_at) = ?
	GROUP BY i.product_name;`

	if err := ad.db.Raw(query, yearInt).Scan(&orderDetails).Error; err != nil {
		return []models.OrderDetailsAdmin{}, err
	}

	return orderDetails, nil
}

func (ad *adminRepository) SalesByMonth(yearInt int, monthInt int) ([]models.OrderDetailsAdmin, error) {
	var orderDetails []models.OrderDetailsAdmin

	query := `SELECT i.product_name, SUM(oi.total_price) AS total_amount
	FROM orders o
	JOIN order_items oi ON o.id = oi.order_id
	JOIN inventories i ON oi.inventory_id = i.product_id
	WHERE o.payment_status = 'PAID'
	AND EXTRACT(YEAR FROM o.created_at) = ?
	AND EXTRACT(MONTH FROM o.created_at) = ?
	GROUP BY i.product_name;`

	if err := ad.db.Raw(query, yearInt, monthInt).Scan(&orderDetails).Error; err != nil {
		return []models.OrderDetailsAdmin{}, err
	}

	return orderDetails, nil
}

func (ad *adminRepository) SalesByDay(yearInt int, monthInt int, dayInt int) ([]models.OrderDetailsAdmin, error) {
	var orderDetails []models.OrderDetailsAdmin

	query := `SELECT i.product_name, SUM(oi.total_price) AS total_amount
	FROM orders o
	JOIN order_items oi ON o.id = oi.order_id
	JOIN inventories i ON oi.inventory_id = i.product_id
	WHERE o.payment_status = 'PAID'
	AND EXTRACT(YEAR FROM o.created_at) = ?
	AND EXTRACT(MONTH FROM o.created_at) = ?
	AND EXTRACT(DAY FROM o.created_at) = ?
	GROUP BY i.product_name;`

	if err := ad.db.Raw(query, yearInt, monthInt, dayInt).Scan(&orderDetails).Error; err != nil {
		return []models.OrderDetailsAdmin{}, err
	}

	return orderDetails, nil
}

func (ad *adminRepository) GetAdminHashPassword(id int) (string, error) {

	var hashPass string

	if err := ad.db.Raw("SELECT  password from admins where id = ?", id).Scan(&hashPass).Error; err != nil {
		return "", err
	}

	return hashPass, nil
}

func (ad *adminRepository) UpdateAdminPass(id int, NewPass string) error {

	result := ad.db.Exec(`UPDATE admins SET password = ? where id = ?`, NewPass, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return errors.New("nothing updated")
	}

	return nil

}
