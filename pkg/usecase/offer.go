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
}

func NewOfferUseCase(ofRep interfaces.OfferRepository, cat interfaces.CategoryRepository) *offerUsecase {
	return &offerUsecase{
		offRepo: ofRep,
		catRepo: cat,
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

	count, err := of.offRepo.CheckOfferExist(model.CategoryID)
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

	count, err := of.offRepo.CheckOfferExist(cID)
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

	count, err := of.offRepo.CheckOfferExist(cID)
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
