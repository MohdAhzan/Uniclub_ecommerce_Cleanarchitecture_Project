package repository

import (
	"errors"
	"project/pkg/utils/models"

	"gorm.io/gorm"
)

type wishlistRepository struct {
	DB *gorm.DB
}

func NewWishlistRepository(db *gorm.DB) *wishlistRepository {

	return &wishlistRepository{
		DB: db,
	}
}

func (w *wishlistRepository) AddToWishlist(userID, inventoryID int) error {

	err := w.DB.Exec(`INSERT INTO wishlists (user_id,inventory_id) VALUES($1,$2)`, userID, inventoryID).Error
	if err != nil {
		return err
	}
	return nil

}

func (w *wishlistRepository) CheckInWishlist(userID, InventoryID int) (int, bool, error) {

	var is_deleted bool

	result := w.DB.Raw("select is_deleted from wishlists where user_id = ? and inventory_id = ? ", userID, InventoryID).Scan(&is_deleted)
	if result.Error != nil {
		return 0, false, result.Error
	}
	var count int
	err := w.DB.Raw("select count(*) from wishlists where user_id = ? and inventory_id = ?", userID, InventoryID).Scan(&count).Error
	if err != nil {
		return 0, false, err
	}
	return count, is_deleted, nil
}

func (w *wishlistRepository) WishlistProductActivate(userID, InventoryID int) error {

	result := w.DB.Exec(`UPDATE wishlists SET is_deleted = false where user_id = ? and inventory_id = ?`, userID, InventoryID)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return errors.New("no wishlist is activated")
	}

	return nil
}

func (w *wishlistRepository) CheckWishlist(userID int) (int, error) {

	var count int

	err := w.DB.Raw("select count(*) from wishlists where user_id = ?", userID).Scan(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil

}

func (w *wishlistRepository) GetWishlist(userID int) ([]models.Inventories, error) {

	var wishlistData []models.Inventories

	err := w.DB.Raw("select i.* from wishlists w join inventories i  on i.product_id = w.inventory_id where user_id = ? and is_deleted = false", userID).Scan(&wishlistData).Error
	if err != nil {
		return []models.Inventories{}, err
	}

	return wishlistData, nil

}

func (w *wishlistRepository) RemoveFromWishlist(pid int) error {

	err := w.DB.Exec(`UPDATE wishlists SET is_deleted = TRUE WHERE inventory_id = ?`, pid).Error
	if err != nil {
		return err
	}

	return nil
}
