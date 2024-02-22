package interfaces

import "project/pkg/utils/models"

type InventoryUseCase interface {
	AddInventory(inventory models.AddInventory) error
	GetProductsForAdmin() ([]models.Inventories, error)
	GetProductsForUsers() ([]models.Inventories, error)
	DeleteInventory(pid int) error
	EditInventory(pid int, model models.EditInventory) error
}
