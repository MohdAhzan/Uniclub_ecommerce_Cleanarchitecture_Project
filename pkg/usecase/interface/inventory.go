package interfaces

import (
	"mime/multipart"
	"project/pkg/utils/models"
)

type InventoryUseCase interface {
	AddInventory(inventory models.AddInventory, file *multipart.FileHeader) (models.InventoryResponse, error)
	GetProductsForAdmin() ([]models.Inventories, error)
	GetProductsForUsers() ([]models.Inventories, error)
	DeleteInventory(pid int) error
	EditInventory(pid int, model models.EditInventory) error
	SearchProducts(pdtName string) ([]models.Inventories, error)
}
