package interfaces

import (
	"project/pkg/utils/models"
)

type OrderUseCase interface {
	OrderFromCart(order models.Order) error
}
