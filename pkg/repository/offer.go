package repository

import (
	"fmt"
	"project/pkg/utils/domain"
	"project/pkg/utils/models"
	"time"

	"gorm.io/gorm"
)

type OfferRepository struct {
	DB *gorm.DB
}

func NewOfferRepository(db *gorm.DB) *OfferRepository {
	return &OfferRepository{
		DB: db,
	}
}

func (of *OfferRepository) AddNewCategoryOffer(model models.AddCategoryOffer) error {

	fmt.Println("models passed to repo offer", model)

	err := of.DB.Exec(`INSERT INTO category_offers (category_id,discount_rate,created_at,valid_till,is_active) Values(?,?,?,?,?)`, model.CategoryID, model.DiscountRate, time.Now(), model.ValidTill, true).Error
	if err != nil {
		return err
	}

	return nil
}

func (of *OfferRepository) GetOfferDiscountPercentage(CategoryID int) (float64, error) {

	var discountRate float64

	err := of.DB.Raw("select discount_rate from category_offers where category_id = ?", CategoryID).Scan(&discountRate).Error
	if err != nil {
		return 0, err
	}
	return discountRate, nil

}

func (of *OfferRepository) CheckOfferExist(categoryID int) (int, error) {

	var count int
	err := of.DB.Raw("select count(*) from category_offers where category_id = ?", categoryID).Scan(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (of *OfferRepository) GetAllCategoryOffers() ([]domain.CategoryOffers, error) {

	var catOffers []domain.CategoryOffers

	err := of.DB.Raw("select * from category_offers").Scan(&catOffers).Error
	if err != nil {
		return []domain.CategoryOffers{}, err
	}

	return catOffers, nil

}

func (of *OfferRepository) EditCategoryOffer(newDiscount float64, cID int) error {

	err := of.DB.Exec(`update category_offers  set discount_rate = ? where category_id = ?`, newDiscount, cID).Error
	if err != nil {
		return err
	}
	return nil
}

func (of *OfferRepository) CheckCategoryOfferStatus(cID int) (bool, error) {

	var status bool

	err := of.DB.Raw("select is_active from category_offers where category_id = ?", cID).Scan(&status).Error
	if err != nil {
		return false, err
	}
	return status, nil
}

func (of *OfferRepository) ValidorInvalidCategoryOffers(status bool, CID int) error {

	err := of.DB.Exec(`update category_offers  set is_active = ? where category_id = ?`, status, CID).Error
	if err != nil {
		return err
	}
	return nil
}
