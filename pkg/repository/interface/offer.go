package interfaces

import (
	"github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/domain"
	"github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/models"
)

type OfferRepository interface {
	AddNewCategoryOffer(model models.AddCategoryOffer) error
	GetCategoryOfferDiscountPercentage(CategoryID int) (float64, string, error)
	CheckCategoryOfferExist(categoryID int) (int, error)
	GetAllCategoryOffers() ([]domain.CategoryOffers, error)
	EditCategoryOffer(newDiscount float64, cID int) error
	CheckCategoryOfferStatus(cID int) (bool, error)
	ValidorInvalidCategoryOffers(status bool, CID int) error

	CheckInventoryOfferExist(InventoryID int) (int, error)
	AddNewInventoryOffer(model models.AddInventoryOffer) error
	GetInventoryOffers() ([]models.GetInventoryOffers, error)
	EditInventoryOffer(newDiscount float64, InventoryID int) error
	CheckInventoryOfferStatus(inventoryID int) (bool, error)
	ValidorInvalidInventoryOffers(status bool, InventoryID int) error
	GetInventoryOfferDiscountPercentage(InventoryId int) (float64, string, error)
}
