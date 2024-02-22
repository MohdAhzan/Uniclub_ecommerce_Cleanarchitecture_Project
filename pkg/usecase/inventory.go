package usecase

import (
	"errors"
	interfaces "project/pkg/repository/interface"
	"project/pkg/utils/models"
)

type InventoryUseCase struct {
	repository interfaces.InventoryRepository
}

func NewInventoryUseCase(repo interfaces.InventoryRepository) *InventoryUseCase {
	return &InventoryUseCase{
		repository: repo,
	}

}

func (Inv *InventoryUseCase) AddInventory(inventory models.AddInventory) error {

	exists, err := Inv.repository.CheckCategoryID(inventory.CategoryID)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("category of this ID doesn't exist ")
	}

	Err := Inv.repository.AddInventory(inventory)

	if Err != nil {
		return Err
	}
	return nil
}

func (Inv *InventoryUseCase) GetProductsForAdmin() ([]models.Inventories, error) {
	productDetails, err := Inv.repository.ListProducts()
	if err != nil {
		return []models.Inventories{}, err
	}
	return productDetails, nil
}

func (Inv *InventoryUseCase) GetProductsForUsers() ([]models.Inventories, error) {

	productDetails, err := Inv.repository.ListProducts()
	if err != nil {
		return []models.Inventories{}, err
	}

	return productDetails, nil
}

func (Inv *InventoryUseCase) DeleteInventory(pid int) error {

	err := Inv.repository.DeleteInventory(pid)
	if err != nil {
		return err
	}

	return nil
}

func (Inv *InventoryUseCase) EditInventory(pid int, model models.EditInventory) error {

	err := Inv.repository.EditInventory(pid, model)

	if err != nil {
		return err
	}

	return nil

}
