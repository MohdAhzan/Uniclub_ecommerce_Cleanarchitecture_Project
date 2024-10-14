package interfaces

import "github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/models"

type WishlistRepository interface {
	AddToWishlist(userID, InventoryID int) error
	CheckInWishlist(userID, InventoryID int) (int, bool, error)
	WishlistProductActivate(userID, InventoryID int) error
	CheckWishlist(userID int) (int, error)
	GetWishlist(userID int) ([]models.Inventories, error)
	RemoveFromWishlist(pid int) error
}
