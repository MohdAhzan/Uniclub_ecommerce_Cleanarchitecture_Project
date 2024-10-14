package repository

import (
	"fmt"
	"github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/domain"
	"github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/models"
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

	err := of.DB.Exec(`INSERT INTO category_offers (category_id,offer_name,discount_rate,created_at,valid_till,is_active) Values(?,?,?,?,?,?)`, model.CategoryID, model.OfferName, model.DiscountRate, time.Now(), model.ValidTill, true).Error
	if err != nil {
		return err
	}

	return nil
}

func (of *OfferRepository) GetCategoryOfferDiscountPercentage(CategoryID int) (float64, string, error) {

	var discountRate float64

	err := of.DB.Raw("select discount_rate from category_offers where category_id = ?", CategoryID).Scan(&discountRate).Error
	if err != nil {
		return 0, "null", err
	}

	var offerName string
	err = of.DB.Raw("select offer_name from category_offers where category_id = ?", CategoryID).Scan(&offerName).Error
	if err != nil {
		return 0, "null", err
	}

	return discountRate, offerName, nil

}

func (of *OfferRepository) CheckCategoryOfferExist(categoryID int) (int, error) {

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

func (of *OfferRepository) CheckInventoryOfferExist(InventoryID int) (int, error) {

	var count int

	err := of.DB.Raw("select count(*) from inventory_offers where inventory_id = ?", InventoryID).Scan(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil

}

func (of *OfferRepository) AddNewInventoryOffer(model models.AddInventoryOffer) error {

	err := of.DB.Exec(`INSERT INTO inventory_offers (inventory_id,offer_name,discount_rate,created_at,valid_till,is_active) Values(?,?,?,?,?,?)`, model.InventoryID, model.OfferName, model.DiscountRate, time.Now(), model.ValidTill, true).Error
	if err != nil {
		return err
	}

	return nil
}
func (of *OfferRepository) GetInventoryOffers() ([]models.GetInventoryOffers, error) {

	var model []models.GetInventoryOffers
	err := of.DB.Raw("select o.id,o.inventory_id,i.product_name,o.offer_name,o.discount_rate,o.created_at,o.valid_till,o.is_active from inventory_offers o join inventories i on i.product_id = o.inventory_id").Scan(&model).Error
	if err != nil {
		return []models.GetInventoryOffers{}, err
	}
	fmt.Println(model)

	return model, nil
}

func (of *OfferRepository) CheckInventoryOfferStatus(inventoryID int) (bool, error) {

	var status bool

	err := of.DB.Raw("select is_active from inventory_offers where inventory_id = ?", inventoryID).Scan(&status).Error
	if err != nil {
		return false, err
	}
	return status, nil
}

func (of *OfferRepository) ValidorInvalidInventoryOffers(status bool, InventoryID int) error {

	err := of.DB.Exec(`update inventory_offers  set is_active = ? where inventory_id = ?`, status, InventoryID).Error
	if err != nil {
		return err
	}
	return nil
}

func (of *OfferRepository) EditInventoryOffer(newDiscount float64, InventoryID int) error {

	err := of.DB.Exec(`update inventory_offers  set discount_rate = ? where inventory_id = ?`, newDiscount, InventoryID).Error
	if err != nil {
		return err
	}
	return nil

}

func (of *OfferRepository) GetInventoryOfferDiscountPercentage(InventoryId int) (float64, string, error) {
	var discountRate float64

	err := of.DB.Raw("select discount_rate from inventory_offers where inventory_id = ?", InventoryId).Scan(&discountRate).Error
	if err != nil {
		return 0, "null", err
	}

	var offerName string
	err = of.DB.Raw("select offer_name from inventory_offers where inventory_id = ?", InventoryId).Scan(&offerName).Error
	if err != nil {
		return 0, "null", err
	}

	return discountRate, offerName, nil

}
