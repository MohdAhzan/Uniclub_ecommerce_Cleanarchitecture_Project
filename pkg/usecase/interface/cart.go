package interfaces

import (
	"project/pkg/utils/models"
)

type CartUseCase interface {
	AddtoCart(pid, UserID int) (models.CartResponse, error)
	GetCart(userID int)(models.CartResponse,error)
}
