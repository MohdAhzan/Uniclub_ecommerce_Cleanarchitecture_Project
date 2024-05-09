package repository

import (
	"errors"
	"fmt"
	"project/pkg/utils/domain"
	"project/pkg/utils/models"
	"time"

	"gorm.io/gorm"
)

type orderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{
		DB: db,
	}
}

func (o *orderRepository) OrderItems(userID, addressID, payment_id, couponID int, TotalCartPrice float64) (int, error) {

	var orderID int

	if couponID == 0 {

		err := o.DB.Raw(`INSERT INTO orders (user_id,address_id,payment_method_id,final_price,coupon_used_by_id)
		VALUES (?, ?,?, ?,NULL)
		RETURNING id`, userID, addressID, payment_id, TotalCartPrice).Scan(&orderID).Error
		if err != nil {
			return 0, err
		}

	} else {

		err := o.DB.Raw(`INSERT INTO orders (user_id,address_id,payment_method_id,final_price,coupon_used_by_id)
		VALUES (?, ?,?, ?,?)
		RETURNING id`, userID, addressID, payment_id, TotalCartPrice, couponID).Scan(&orderID).Error
		if err != nil {
			return 0, err
		}

	}

	return orderID, nil

}

func (o *orderRepository) AddOrderProducts(orderID int, cart []models.GetCart) error {

	for _, data := range cart {

		if err := o.DB.Exec(`INSERT INTO order_items (order_id,inventory_id,quantity,total_price) VALUES(?,?,?,?)`, orderID, data.ProductID, data.Quantity, data.TotalPrice).Error; err != nil {
			return err
		}
	}

	return nil
}

func (o *orderRepository) GetOrders(userID int) ([]domain.Order, error) {

	var orders []domain.Order

	if err := o.DB.Raw("select * from orders where user_id = ?", userID).Scan(&orders).Error; err != nil {
		return []domain.Order{}, err
	}

	return orders, nil
}

func (o *orderRepository) GetOrderImages(orderID int) ([]string, error) {

	var images []string

	err := o.DB.Raw("SELECT inventories.image FROM order_items JOIN inventories ON inventories.product_id = order_items.inventory_id JOIN orders ON orders.id = order_items.order_id WHERE orders.id = ? ", orderID).Scan(&images).Error
	if err != nil {
		return nil, err
	}
	return images, nil
}
func (o *orderRepository) GetOrderProductIDs(orderID int) ([]int, error) {

	var inventoryIDs []int

	err := o.DB.Raw("select inventory_id from order_items where order_id = ? ", orderID).Scan(&inventoryIDs).Error
	if err != nil {
		return nil, err
	}
	return inventoryIDs, nil
}

func (o *orderRepository) GetOrderAddress(orderID int) (domain.Address, models.OrderData, error) {

	var address domain.Address

	err := o.DB.Raw("SELECT a.* FROM addresses a JOIN orders o ON a.id = o.address_id where o.id = ? ", orderID).Scan(&address).Error
	if err != nil {
		return domain.Address{}, models.OrderData{}, err
	}

	var orderData models.OrderData

	err = o.DB.Raw("select payment_method_id,order_status, final_price,payment_status from orders where id = ?", orderID).Scan(&orderData).Error
	if err != nil {
		return domain.Address{}, models.OrderData{}, err
	}
	return address, orderData, nil

}

func (o *orderRepository) CheckOrderStatusByID(orderID int) (string, error) {

	var status string

	err := o.DB.Raw("select order_status from orders where id = ?", orderID).Scan(&status).Error
	if err != nil {
		return "", err
	}

	return status, nil
}

func (o *orderRepository) CancelOrder(orderID int) error {

	result := o.DB.Exec(`UPDATE orders SET order_status = 'CANCELED' ,updated_at = ? WHERE id = ? `, time.Now(), orderID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return errors.New("no rows updated")
	}
	return nil
}

func (o *orderRepository) ReturnOrder(orderID int) error {

	result := o.DB.Exec(`UPDATE orders SET order_status='RETURN_REQUESTED',updated_at = ? WHERE id = ?`, time.Now(), orderID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return errors.New("no rows updated")
	}
	return nil
}

func (o *orderRepository) CheckOrderByID(orderID int) error {

	var count int
	err := o.DB.Raw("SELECT count(*) from orders where id = ?", orderID).Scan(&count).Error
	if err != nil {
		return err
	}

	if count < 1 {
		return errors.New("no orders not availble in this id")
	}

	return nil

}

func (o *orderRepository) EditOrderStatus(orderID int, status string) error {

	err := o.DB.Exec(`UPDATE orders SET order_status = $1 ,updated_at =$2  WHERE id = $3 `, status, time.Now(), orderID).Error
	if err != nil {
		return err
	}
	return nil
}

func (o *orderRepository) GetPaymentMethodsByID(PaymentMethodID int) (string, error) {
	var paymentMethod string
	err := o.DB.Raw("select payment_name from payment_methods where id = ?", PaymentMethodID).Scan(&paymentMethod).Error
	if err != nil {
		return "", err
	}
	return paymentMethod, nil
}

func (o *orderRepository) FindOrderAmount(orderID int) (float64, error) {

	var orderAmount float64

	err := o.DB.Raw("select final_price from orders where id = ?", orderID).Scan(&orderAmount).Error
	if err != nil {
		return 0, err
	}

	return orderAmount, nil
}

func (o *orderRepository) FindOrderedUserID(orderID int) (int, error) {

	var userid int

	err := o.DB.Raw("select user_id from orders where id = ?", orderID).Scan(&userid).Error
	if err != nil {
		return 0, err
	}

	return userid, nil
}

func (o *orderRepository) GetPaymentStatusByID(orderID int) (string, error) {

	var paymentStatus string

	err := o.DB.Raw("select payment_status  from orders where id = ?", orderID).Scan(&paymentStatus).Error
	if err != nil {
		return "", err
	}

	return paymentStatus, nil
}

func (o *orderRepository) GetAllOrderItemsByOrderID(orderID int) ([]domain.EachProductOrderDetails, error) {

	var Allmodel []domain.EachProductOrderDetails

	var data []models.EachOrderData

	err := o.DB.Raw("SELECT inventory_id as product_id , SUM(quantity) AS total_quantity,CAST(SUM(total_price) AS DECIMAL(10, 2)) AS total_price FROM order_items WHERE order_id = ? GROUP BY inventory_id", orderID).Scan(&data).Error
	if err != nil {
		return []domain.EachProductOrderDetails{}, err

	}

	for _, value := range data {
		var model domain.EachProductOrderDetails
		model.ProductID = uint(value.ProductID)
		model.Quantity = uint(value.TotalQuantity)
		model.ProductPrice = value.TotalPrice
		Allmodel = append(Allmodel, model)
	}
	fmt.Println(data)
	return Allmodel, nil

}

func (o *orderRepository) CheckIndividualOrders(orderID, pID int) (int, error) {

	var count int

	err := o.DB.Raw("select count(*) from order_items where order_id = ? and inventory_id = ?", orderID, pID).Scan(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (o *orderRepository) DeleteProductInOrder(orderID, pID int) (float64, error) {

	var productPrice float64

	err := o.DB.Raw("DELETE FROM order_items WHERE order_id = $1 AND inventory_id = $2 RETURNING total_price", orderID, pID).Scan(&productPrice).Error
	if err != nil {
		return 0, err
	}

	return productPrice, nil
}

func (o *orderRepository) UpdateFinalOrderPrice(orderID int, NewPrice float64) error {

	err := o.DB.Exec(`UPDATE orders SET final_price = ? where id = ?`, NewPrice, orderID).Error
	if err != nil {
		return err
	}
	return nil
}
