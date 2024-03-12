package interfaces

import "project/pkg/utils/models"

type InventoryRepository interface {
	AddInventory(Inventory models.AddInventory, URL string) (models.InventoryResponse, error)
	CheckCategoryID(CategoryID int) (bool, error)
	ListProducts() ([]models.Inventories, error)
	DeleteInventory(pid int) error
	EditInventory(pid int, model models.EditInventory) error
	CheckProduct(productName string, size string) (bool, error)
}
