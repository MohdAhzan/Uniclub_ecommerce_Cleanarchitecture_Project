package repository

import (
	"errors"
	"fmt"
	interfaces "project/pkg/repository/interface"
	"project/pkg/utils/models"
	"strings"

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

func (inv *InventoryRepository) AddInventory(Inventory models.AddInventory, URL string) (models.InventoryResponse, error) {

	query := `INSERT INTO Inventories (category_id,product_name,size,stock,price,image) VALUES (?,?,?,?,?,?);`

	err := inv.DB.Exec(query, Inventory.CategoryID, Inventory.ProductName, Inventory.Size, Inventory.Stock, Inventory.Price, URL).Error

	if err != nil {
		return models.InventoryResponse{}, err
	}

	var InventoryResponse models.InventoryResponse

	return InventoryResponse, nil
}

func (inv *InventoryRepository) ListProducts() ([]models.Inventories, error) {

	var productDetails []models.Inventories

	err := inv.DB.Raw("select * from inventories").Scan(&productDetails).Error
	if err != nil {
		return []models.Inventories{}, err
	}

	return productDetails, nil
}

func (Inv *InventoryRepository) DeleteInventory(pid int) error {

	result := Inv.DB.Exec("DELETE from inventories WHERE product_id=?", pid)
	errDelete := fmt.Sprintf("No product is in inventory of id %d ", pid)
	if result.RowsAffected < 1 {
		return errors.New(errDelete)
	}
	return nil
}

func (inv *InventoryRepository) EditInventory(pid int, model models.EditInventory) error {
	fmt.Println("{{{{PRODUCT ID}}}}", pid)
	result := inv.DB.Exec("UPDATE inventories SET category_id = $1, product_name = $2, size = $3, stock = $4 ,price = $5 WHERE product_id = $6", model.CategoryID, model.ProductName, model.Size, model.Stock, model.Price, pid)

	if result.RowsAffected < 1 {
		return errors.New("error")
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
		return false, nil
	} else {
		return true, nil
	}

}

func (inv *InventoryRepository) CheckProduct(ProductName string, size string) (bool, error) {

	var count int

	err := inv.DB.Raw("select count(*) from inventories where product_name = ? and size = ?", ProductName, size).Scan(&count).Error

	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func (inv *InventoryRepository) CheckStock(pid int) (int, error) {

	var stock int

	err := inv.DB.Raw("select stock from inventories where product_id = ?", pid).Scan(&stock).Error
	if err != nil {
		return 0, err
	}

	return stock, nil
}

func (inv *InventoryRepository) GetProductImages(pid int) (string, error) {

	var image string
	err := inv.DB.Raw("select image from inventories where product_id =? ", pid).Scan(&image).Error
	if err != nil {
		return "", err
	}
	return image, nil
}

func (inv *InventoryRepository) GetCategoryID(pid int) (int, error) {
	var categoryID int
	err := inv.DB.Raw("select category_id from inventories where product_id =? ", pid).Scan(&categoryID).Error
	if err != nil {
		return 0, err
	}
	return categoryID, nil

}

func (c *InventoryRepository) FindStock(pid int) (int, error) {

	var stock int

	err := c.DB.Raw("select stock from inventories where product_id = ?", pid).Scan(&stock).Error
	if err != nil {
		return 0, err
	}
	return stock, nil
}

func (c *InventoryRepository) FindPrice(pid int) (float64, error) {
	var price float64
	err := c.DB.Raw("select price from inventories where product_id = ?", pid).Scan(&price).Error
	if err != nil {
		return 0, err
	}
	return price, nil
}

func (i *InventoryRepository) SearchProducts(pdtName string) ([]models.Inventories, error) {

	var products []models.Inventories

	pdtName = strings.TrimSpace(pdtName)
	
	err := i.DB.Raw("select product_id,category_id,product_name,image,stock,price from inventories where product_name ilike ?", "%"+pdtName+"%").Scan(&products).Error
	if err != nil {
		return []models.Inventories{}, err
	}

	return products, nil
}
