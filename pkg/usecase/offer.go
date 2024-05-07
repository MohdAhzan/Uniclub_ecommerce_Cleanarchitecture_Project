package usecase

import (
	"errors"
	"fmt"
	interfaces "project/pkg/repository/interface"
	domain "project/pkg/utils/domain"
	"project/pkg/utils/models"
)

type offerUsecase struct {
	offRepo interfaces.OfferRepository
	catRepo interfaces.CategoryRepository
	invRepo interfaces.InventoryRepository
}

func NewOfferUseCase(ofRep interfaces.OfferRepository, cat interfaces.CategoryRepository, inv interfaces.InventoryRepository) *offerUsecase {
	return &offerUsecase{
		offRepo: ofRep,
		catRepo: cat,
		invRepo: inv,
	}
}

func (of offerUsecase) AddCategoryOffer(model models.AddCategoryOffer) error {

	exists, err := of.catRepo.CheckCategoryByID(model.CategoryID)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("no category in this ID")
	}

	count, err := of.offRepo.CheckCategoryOfferExist(model.CategoryID)
	if err != nil {
		return err
	}
	if count >= 1 {
		return errors.New("offer for this category already exists")
	}
	err = of.offRepo.AddNewCategoryOffer(model)
	if err != nil {
		return err
	}

	return nil
}

func (of offerUsecase) GetAllCategoryOffers() ([]domain.CategoryOffers, error) {

	catOffers, err := of.offRepo.GetAllCategoryOffers()
	if err != nil {
		return []domain.CategoryOffers{}, err
	}

	return catOffers, nil
}

func (of offerUsecase) EditCategoryOffer(newDiscount float64, cID int) error {

	count, err := of.offRepo.CheckCategoryOfferExist(cID)
	if err != nil {
		return err
	}
	if count < 1 {
		return errors.New("no offers to edit for this category")
	}
	err = of.offRepo.EditCategoryOffer(newDiscount, cID)
	if err != nil {
		return err
	}

	return nil
}

func (of offerUsecase) ValidorInvalidCategoryOffers(status bool, cID int) error {

	count, err := of.offRepo.CheckCategoryOfferExist(cID)
	if err != nil {
		return err
	}
	if count < 1 {
		return errors.New("no offers for this category please add one first")
	}

	oldStatus, err := of.offRepo.CheckCategoryOfferStatus(cID)
	if err != nil {
		return err
	}
	if oldStatus == status {
		errMsg := fmt.Sprintf("offer status is already %t", status)
		return errors.New(errMsg)

	}

	err = of.offRepo.ValidorInvalidCategoryOffers(status, cID)
	if err != nil {
		return err
	}

	return nil
}

func (of offerUsecase) AddInventoryOffer(model models.AddInventoryOffer) error {

	stockCount, err := of.invRepo.FindStock(model.InventoryID)
	if err != nil {
		return err
	}

	if stockCount < 1 {
		return errors.New("no product exist in this product ID")
	}

	count, err := of.offRepo.CheckInventoryOfferExist(model.InventoryID)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("offer for this product already exists")
	}

	err = of.offRepo.AddNewInventoryOffer(model)
	if err != nil {
		return err
	}

	return nil

}

func (of offerUsecase) GetInventoryOffers() ([]models.GetInventoryOffers, error) {

	offerData, err := of.offRepo.GetInventoryOffers()
	if err != nil {
		return []models.GetInventoryOffers{}, err

	}

	return offerData, nil

}

func (of offerUsecase) EditInventoryOffer(newDiscount float64, InventoryID int) error {

	count, err := of.offRepo.CheckInventoryOfferExist(InventoryID)
	if err != nil {
		return err
	}
	if count < 1 {
		return errors.New("no offers to edit for this product")
	}
	err = of.offRepo.EditInventoryOffer(newDiscount, InventoryID)
	if err != nil {
		return err
	}

	return nil

}

func (of offerUsecase) ValidorInvalidInventoryOffers(status bool, inventoryID int) error {

	count, err := of.offRepo.CheckInventoryOfferExist(inventoryID)
	if err != nil {
		return err
	}
	if count < 1 {
		return errors.New("no offer for this product please add one first")
	}

	oldStatus, err := of.offRepo.CheckInventoryOfferStatus(inventoryID)
	if err != nil {
		return err

	}
	fmt.Println("old status", oldStatus)
	if oldStatus == status {
		errMsg := fmt.Sprintf("offer status is already %t", status)
		return errors.New(errMsg)

	}
	fmt.Println("new status", status)
	err = of.offRepo.ValidorInvalidInventoryOffers(status, inventoryID)
	if err != nil {
		return err
	}

	return nil
}
