package interfaces

import "project/pkg/utils/models"

type WishlistUsecase interface {
	AddToWishlist(userID, inventoryID int) error
	GetWishlist(userID int) ([]models.Inventories, error)
	RemoveFromWishlist(userID, pID int) error
}
