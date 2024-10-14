package interfaces

import "github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/models"

type WishlistUsecase interface {
	AddToWishlist(userID, inventoryID int) error
	GetWishlist(userID int) ([]models.Inventories, error)
	RemoveFromWishlist(userID, pID int) error
}
