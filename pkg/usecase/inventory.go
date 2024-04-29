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
	offerRepo  interfaces.OfferRepository
}

func NewInventoryUseCase(repo interfaces.InventoryRepository, h helper_interfaces.Helper, off interfaces.OfferRepository) *InventoryUseCase {
	return &InventoryUseCase{
		repository: repo,
		helper:     h,
		offerRepo:  off,
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

	//check if any offers are there

	for i, Product := range productDetails {

		// if the category id of these products are in offer table discount the price to new one

		DiscountRate, err := Inv.offerRepo.GetOfferDiscountPercentage(Product.CategoryID)
		if err != nil {
			return []models.Inventories{}, err
		}

		var discount float64

		if DiscountRate > 0 {
			discount = (Product.Price * float64(DiscountRate)) / 100
		}

		//Discounted Price = Original Price - (Original Price * (Discount Percentage / 100))

		Product.DiscountedPrice = Product.Price - discount

		fmt.Println("discounted Price", Product.DiscountedPrice)
		fmt.Println("ORginal Price", Product.Price)

		productDetails[i].DiscountedPrice = Product.DiscountedPrice
	}
	return productDetails, nil
}

func (Inv *InventoryUseCase) GetProductsForUsers() ([]models.Inventories, error) {

	productDetails, err := Inv.repository.ListProducts()
	if err != nil {
		return []models.Inventories{}, err
	}

	//check if any offers are there

	for i, Product := range productDetails {

		// if the category id of these products are in offer table discount the price to new one

		DiscountRate, err := Inv.offerRepo.GetOfferDiscountPercentage(Product.CategoryID)
		if err != nil {
			return []models.Inventories{}, err
		}

		var discount float64

		if DiscountRate > 0 {
			discount = (Product.Price * float64(DiscountRate)) / 100
		}

		//Discounted Price = Original Price - (Original Price * (Discount Percentage / 100))

		Product.DiscountedPrice = Product.Price - discount

		fmt.Println("discounted Price", Product.DiscountedPrice)
		fmt.Println("ORginal Price", Product.Price)

		productDetails[i].DiscountedPrice = Product.DiscountedPrice
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

func (Inv *InventoryUseCase) SearchProducts(pdtName string) ([]models.Inventories, error) {

	productDetails, err := Inv.repository.SearchProducts(pdtName)
	if err != nil {
		return []models.Inventories{}, err
	}

	//check if any offers are there

	for i, Product := range productDetails {

		// if the category id of these products are in offer table discount the price to new one

		DiscountRate, err := Inv.offerRepo.GetOfferDiscountPercentage(Product.CategoryID)
		if err != nil {
			return []models.Inventories{}, err
		}

		var discount float64

		if DiscountRate > 0 {
			discount = (Product.Price * float64(DiscountRate)) / 100
		}

		//Discounted Price = Original Price - (Original Price * (Discount Percentage / 100))

		Product.DiscountedPrice = Product.Price - discount

		fmt.Println("discounted Price", Product.DiscountedPrice)
		fmt.Println("ORginal Price", Product.Price)

		productDetails[i].DiscountedPrice = Product.DiscountedPrice
	}

	return productDetails, nil

}
