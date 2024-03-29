package usecase

import (
	"errors"
	"fmt"
	"mime/multipart"
	helper_interfaces "project/pkg/helper/interface"
	interfaces "project/pkg/repository/interface"

	"project/pkg/utils/models"
)

type InventoryUseCase struct {
	repository interfaces.InventoryRepository
	helper     helper_interfaces.Helper
}

func NewInventoryUseCase(repo interfaces.InventoryRepository, h helper_interfaces.Helper) *InventoryUseCase {
	return &InventoryUseCase{
		repository: repo,
		helper:     h,
	}

}

func (Inv *InventoryUseCase) AddInventory(inventory models.AddInventory, file *multipart.FileHeader) (models.InventoryResponse, error) {

	exists, err := Inv.repository.CheckCategoryID(inventory.CategoryID)
	if err != nil {
		return models.InventoryResponse{}, err
	}
	if !exists {
		return models.InventoryResponse{}, errors.New("category of this ID doesn't exist ")
	}

	exists, err = Inv.repository.CheckProduct(inventory.ProductName, inventory.Size)

	if err != nil {
		return models.InventoryResponse{}, err
	}
	if exists {
		errMsg := fmt.Sprintf("Product %s of Size: %s already exists", inventory.ProductName, inventory.Size)
		return models.InventoryResponse{}, errors.New(errMsg)
	}

	var URL string

	URL, err = Inv.helper.AddImageToAwsS3(file)
	if err != nil {
		return models.InventoryResponse{}, err
	}

	inventoryResponse, Err := Inv.repository.AddInventory(inventory, URL)

	if Err != nil {
		return models.InventoryResponse{}, Err
	}
	return inventoryResponse, nil
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
