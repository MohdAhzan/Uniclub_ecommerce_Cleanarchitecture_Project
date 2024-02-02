package repository

import (
	"errors"
	"fmt"
	interfaces "project/pkg/repository/interface"
	"project/pkg/utils/models"

	"gorm.io/gorm"
)

type InventoryRepository struct {
	DB *gorm.DB
}

func NewInventoryRepository(DB *gorm.DB) interfaces.InventoryRepository {
	return &InventoryRepository{
		DB: DB,
	}
}

func (inv *InventoryRepository) AddInventory(Inventory models.AddInventory) error {

	query := `INSERT INTO Inventories (category_id,product_name,size,stock,price) VALUES (?,?,?,?,?);`

	err := inv.DB.Exec(query, Inventory.CategoryID, Inventory.ProductName, Inventory.Size, Inventory.Stock, Inventory.Price).Error

	if err != nil {
		return err
	}

	return nil
}

func (inv *InventoryRepository) CheckCategoryID(CategoryID int) (bool, error) {
	var i int

	err := inv.DB.Raw("select count(*) from categories where id= ?", CategoryID).Scan(&i).Error
	if err != nil {
		return false, err
	}
	if i == 0 {
		return false, err
	} else {
		return true, err
	}

}

func (inv *InventoryRepository) ListProducts() ([]models.Inventories, error) {

	var productDetails []models.Inventories

	err := inv.DB.Raw("select * from inventories").Scan(&productDetails).Error
	if err != nil {
		return []models.Inventories{}, err
	}

	return productDetails, nil
}

func (Inv *InventoryRepository)DeleteInventory(pid int) error {

	result := Inv.DB.Exec("DELETE from inventories WHERE product_id=?", pid)
	errDelete := fmt.Sprintf("No product is in inventory of id %d ", pid)
	if result.RowsAffected < 1 {
		return errors.New(errDelete)
	}
	return nil
}
