package repository

import (
	"project/pkg/utils/models"

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
	err := o.DB.Raw(` INSERT INTO orders (user_id,address_id,price)
    VALUES (?, ?, ?)
    RETURNING id
    `, userID, addressID, TotalCartPrice).Scan(&orderID).Error

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
