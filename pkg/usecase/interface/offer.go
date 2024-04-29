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
}
