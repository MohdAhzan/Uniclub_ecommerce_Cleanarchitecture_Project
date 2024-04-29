package interfaces

import (
	"project/pkg/utils/domain"
	"project/pkg/utils/models"
)

type OfferRepository interface {
	AddNewCategoryOffer(model models.AddCategoryOffer) error
	GetOfferDiscountPercentage(CategoryID int) (float64, error)
	CheckOfferExist(categoryID int) (int, error)
	GetAllCategoryOffers() ([]domain.CategoryOffers, error)
	EditCategoryOffer(newDiscount float64, cID int) error
	CheckCategoryOfferStatus(cID int) (bool, error)
	ValidorInvalidCategoryOffers(status bool, CID int) error
}
