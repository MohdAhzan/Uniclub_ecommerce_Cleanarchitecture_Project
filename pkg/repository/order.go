package repository

import (
	"errors"
	"project/pkg/utils/domain"
	"project/pkg/utils/models"
	"time"

	"gorm.io/gorm"
)

type orderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{
		DB: db,
	}
}

func (o *orderRepository) OrderItems(userID, addressID int, TotalCartPrice float64) (int, error) {

	var orderID int
	err := o.DB.Raw(`INSERT INTO orders (user_id,address_id,price)
    VALUES (?, ?, ?)
    RETURNING id`, userID, addressID, TotalCartPrice).Scan(&orderID).Error

	if err != nil {
		return 0, err
	}

	return orderID, nil

}

func (o *orderRepository) AddOrderProducts(orderID int, cart []models.GetCart) error {

	for _, data := range cart {
		var pID int
		if err := o.DB.Raw("select product_id from inventories where product_name = ?", data.ProductName).Scan(&pID).Error; err != nil {
			return err
		}
		if err := o.DB.Exec(`INSERT INTO order_items (order_id,inventory_id,quantity,total_price) VALUES(?,?,?,?)`, orderID, pID, data.Quantity, data.TotalPrice).Error; err != nil {
			return err
		}
	}

	return nil
}

func (o *orderRepository) GetOrders(userID int) ([]domain.Order, error) {

	var orders []domain.Order

	if err := o.DB.Raw("select * from orders where user_id = ?", userID).Scan(&orders).Error; err != nil {
		return []domain.Order{}, err
	}

	return orders, nil
}

func (o *orderRepository) GetOrderImages(orderID int) ([]string, error) {

	var images []string

	err := o.DB.Raw("SELECT inventories.image FROM order_items JOIN inventories ON inventories.product_id = order_items.inventory_id JOIN orders ON orders.id = order_items.order_id WHERE orders.id = ? ", orderID).Scan(&images).Error
	if err != nil {
		return nil, err
	}
	return images, nil
}

func (o *orderRepository) GetOrderAddress(orderID int) (domain.Address, models.OrderData, error) {

	var address domain.Address

	err := o.DB.Raw("SELECT a.* FROM addresses a JOIN orders o ON a.id = o.address_id where o.id = ? ", orderID).Scan(&address).Error
	if err != nil {
		return domain.Address{}, models.OrderData{}, err
	}

	var orderData models.OrderData

	err = o.DB.Raw("select payment_method,order_status, price,payment_status from orders where id = ?", orderID).Scan(&orderData).Error
	if err != nil {
		return domain.Address{}, models.OrderData{}, err
	}
	return address, orderData, nil

}

func (o *orderRepository) CheckOrderStatusByID(orderID int) (string, error) {

	var status string

	err := o.DB.Raw("select order_status from orders where id = ?", orderID).Scan(&status).Error
	if err != nil {
		return "", err
	}

	return status, nil
}

func (o *orderRepository) CancelOrder(orderID int) error {

	result := o.DB.Exec(`UPDATE orders SET order_status = 'CANCELED' ,updated_at = ? WHERE id = ? `, time.Now(), orderID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return errors.New("no rows updated")
	}
	return nil
}

func (o *orderRepository) ReturnOrder(orderID int) error {

	result := o.DB.Exec(`UPDATE orders SET order_status='RETURN_REQUESTED',updated_at = ? WHERE id = ?`, time.Now(), orderID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return errors.New("no rows updated")
	}
	return nil
}
