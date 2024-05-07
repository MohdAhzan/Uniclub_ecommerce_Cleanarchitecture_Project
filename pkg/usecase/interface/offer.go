package interfaces

import (
	"project/pkg/utils/domain"
	"project/pkg/utils/models"
)

type OfferUsecase interface {
	AddCategoryOffer(model models.AddCategoryOffer) error
	GetAllCategoryOffers() ([]domain.CategoryOffers, error)
	EditCategoryOffer(newDiscount float64, cID int) error
	ValidorInvalidCategoryOffers(status bool, cID int) error

	AddInventoryOffer(model models.AddInventoryOffer) error
	GetInventoryOffers() ([]models.GetInventoryOffers, error)
	EditInventoryOffer(newDiscount float64, InventoryID int) error
	ValidorInvalidInventoryOffers(status bool, inventoryID int) error
}
