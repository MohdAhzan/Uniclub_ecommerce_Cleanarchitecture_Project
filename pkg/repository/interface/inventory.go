package interfaces

import "project/pkg/utils/models"

type InventoryRepository interface {
	AddInventory(Inventory models.AddInventory) error
	CheckCategoryID(CategoryID int) (bool, error)
	ListProducts() ([]models.Inventories, error)
	DeleteInventory(pid int)error
	EditInventory(pid int ,model models.EditInventory)error
}
