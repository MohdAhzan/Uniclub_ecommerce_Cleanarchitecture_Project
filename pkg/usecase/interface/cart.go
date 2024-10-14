package interfaces

import (
	"github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/models"
)

type CartUseCase interface {
	AddtoCart(pid, UserID, quantity int) (models.CartResponse, error)
	GetCart(userID int) (models.CartResponse, error)
	RemoveCart(userID, pid int) error
	DecreaseCartQuantity(userID, quantity, pID int) error
}
