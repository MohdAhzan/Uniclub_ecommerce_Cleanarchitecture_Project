package repository

import (
	"errors"
	"fmt"
	"project/pkg/utils/models"

	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(DB *gorm.DB) *CartRepository {
	return &CartRepository{
		db: DB,
	}
}

func (c *CartRepository) GetCartID(userID int) (int, error) {

	var cartID int

	if err := c.db.Raw("select id from carts where user_id = ?", userID).Scan(&cartID).Error; err != nil {
		return 0, err
	}
	return cartID, nil
}

func (c *CartRepository) CreateNewCart(UserID int) (int, error) {

	var cartID int

	// fmt.Println("userID IN REPO CREATE NEW CART", UserID)
	err := c.db.Exec("INSERT INTO carts (user_id) VALUES ($1)", UserID).Error
	if err != nil {
		return 0, err
	}

	err = c.db.Raw("select id from carts where user_id = ? ", UserID).Scan(&cartID).Error
	if err != nil {
		return 0, err
	}
	fmt.Println("ADDED ")
	return cartID, nil
}

func (c *CartRepository) CheckIfItemIsAlreadyAdded(cart_id, pid int) (bool, error) {
	//check items if already added to cart if its added update the quantity
	var count int

	err := c.db.Raw("select count(*) from cart_items where cart_id = $1 and product_id = $2 ", cart_id, pid).Scan(&count).Error
	if err != nil {
		return false, err
	}
	if count <= 0 {
		return false, nil
	}

	return count > 0, nil

}

func (c *CartRepository) UpdateCartQuantity(cartID, pid, quantity int) error {
	err := c.db.Exec("UPDATE cart_items SET quantity = $1 where cart_id =$2 and product_id =$3", quantity, cartID, pid).Error
	if err != nil {
		return err
	}
	return nil

}

func (c *CartRepository) AddtoCartItems(cartID, pid int) error {

	err := c.db.Exec("insert into cart_items (cart_id,product_id) values ($1,$2)", cartID, pid).Error
	if err != nil {
		return errors.New("error adding product to cartITems")
	}

	return nil

}

func (c *CartRepository) GetProductIDs(cartID int) ([]int, error) {
	var pid []int

	if err := c.db.Raw("select product_id from cart_items where cart_id = ?", cartID).Scan(&pid).Error; err != nil {
		return nil, err
	}

	return pid, nil
}

func (c *CartRepository) GetProductNames(pID int) (string, error) {
	var pName string

	err := c.db.Raw("select product_name from inventories where product_id = ?", pID).Scan(&pName).Error
	if err != nil {
		return "", err
	}
	return pName, nil
}

func (c *CartRepository) FindCartQuantity(pid, cartID int) (int, error) {

	var quantity int

	err := c.db.Raw("select quantity from cart_items where product_id = $1 and cart_id = $2", pid, cartID).Scan(&quantity).Error
	if err != nil {
		return 0, err
	}
	return quantity, nil
}

func (c *CartRepository) RemoveCartItems(pid, cartID int) error {

	result := c.db.Exec("delete from cart_items where product_id = $1 and cart_id = $2", pid, cartID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return errors.New("nothing deleted")
	}

	return nil
}

func (c *CartRepository) GetCartAddress(userID int) (models.Address, error) {
	var addressID int
	if err := c.db.Raw("select address_id from orders where user_id = ?", userID).Scan(&addressID).Error; err != nil {
		return models.Address{}, err
	}

	var address models.Address

	err := c.db.Raw("select user_id,name,address,land_mark,city,pincode,state,phone from addresses where user_id = ?", userID).Scan(&address).Error
	if err != nil {
		return models.Address{}, err
	}
	return address, nil
}
