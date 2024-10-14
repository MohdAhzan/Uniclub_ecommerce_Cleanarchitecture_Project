package usecase

import (
	"errors"
	"fmt"
	interfaces "github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/repository/interface"
	"github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/models"
)

type wishlistUsecase struct {
	wishRepo interfaces.WishlistRepository
	invRepo  interfaces.InventoryRepository
}

func NewWishlistUsecase(wRepo interfaces.WishlistRepository, inv interfaces.InventoryRepository) *wishlistUsecase {
	return &wishlistUsecase{
		wishRepo: wRepo,
		invRepo:  inv,
	}
}

func (w wishlistUsecase) AddToWishlist(userID, InventoryID int) error {

	stock, err := w.invRepo.CheckStock(InventoryID)
	if err != nil {

		return err
	}

	if stock < 1 {
		return errors.New("no products available in this id")
	}

	count, is_deleted, err := w.wishRepo.CheckInWishlist(userID, InventoryID)
	if err != nil {
		return err
	}

	fmt.Println(is_deleted, "checking wishlist")

	if is_deleted {
		fmt.Println("present but not deleted", is_deleted)
		err := w.wishRepo.WishlistProductActivate(userID, InventoryID)
		if err != nil {
			return err
		}
		return nil

	}

	if count > 0 {
		return errors.New("your product is already added to wishlist")
	}
	fmt.Println("not present in wishlist so adding new product")
	err = w.wishRepo.AddToWishlist(userID, InventoryID)
	if err != nil {
		return err
	}

	return nil
}

func (w wishlistUsecase) GetWishlist(userID int) ([]models.Inventories, error) {

	//check if user has any wishlist

	count, err := w.wishRepo.CheckWishlist(userID)
	if err != nil {
		return []models.Inventories{}, err
	}
	if count <= 0 {
		return []models.Inventories{}, errors.New("your wishlist is empty please add some products first")
	}

	wishlistData, err := w.wishRepo.GetWishlist(userID)
	if err != nil {
		return []models.Inventories{}, err
	}

	return wishlistData, nil
}

func (w wishlistUsecase) RemoveFromWishlist(userID, pid int) error {

	//check if user has any wishlist

	count, err := w.wishRepo.CheckWishlist(userID)
	if err != nil {
		return err
	}
	if count <= 0 {
		return errors.New("your wishlist is empty")
	}

	err = w.wishRepo.RemoveFromWishlist(pid)
	if err != nil {
		return err
	}
	return nil
}
