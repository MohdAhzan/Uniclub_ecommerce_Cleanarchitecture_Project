package repository

import (
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

func (o *orderRepository) OrderItems(userID, addressID int, price float64) (int, error) {

	var orderID int
	err := o.DB.Raw(` INSERT INTO orders (user_id,address_id,price)
    VALUES (?, ?, ?)
    RETURNING id
    `, userID, addressID, price).Scan(&orderID).Error

	if err != nil {
		return 0, err
	}

	return orderID, nil

}
