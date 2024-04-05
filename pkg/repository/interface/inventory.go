package interfaces

import "project/pkg/utils/models"

type InventoryRepository interface {
	AddInventory(Inventory models.AddInventory, URL string) (models.InventoryResponse, error)
	CheckCategoryID(CategoryID int) (bool, error)
	ListProducts() ([]models.Inventories, error)
	DeleteInventory(pid int) error
	EditInventory(pid int, model models.EditInventory) error
	CheckProduct(productName string, size string) (bool, error)
	CheckStock(pid int) (int, error)
	GetCategoryID(pid int) (int, error)
	GetProductImages(pid int) (string, error)
	FindStock(pid int) (int, error)
	FindPrice(pid int) (float64, error)
	SearchProducts(pdtName string) ([]models.Inventories, error)
}
