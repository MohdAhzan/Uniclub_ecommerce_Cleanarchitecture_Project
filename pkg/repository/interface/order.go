package interfaces

import "project/pkg/utils/models"

type OrderRepository interface {
	OrderItems(userID, address int, TotalCartPrice float64) (int, error)
	AddOrderProducts(orderID int, cart []models.GetCart) error
}
