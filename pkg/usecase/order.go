package usecase

import (
	"fmt"
	services "project/pkg/repository/interface"
	interfaces "project/pkg/usecase/interface"
	"project/pkg/utils/models"
)

type OrderUseCase struct {
	orderRepo   services.OrderRepository
	cartRepo    services.CartRepository
	cartUseCase interfaces.CartUseCase
}

func NewOrderUseCase(orderRepo services.OrderRepository, cartRepo services.CartRepository, cartUseCase interfaces.CartUseCase) *OrderUseCase {
	return &OrderUseCase{
		orderRepo:   orderRepo,
		cartRepo:    cartRepo,
		cartUseCase: cartUseCase,
	}
}
func (o OrderUseCase) OrderFromCart(order models.Order) error {

	cart, err := o.cartUseCase.GetCart(order.UserID)
	if err != nil {
		return err
	}
	fmt.Println("model GET CART", cart)
	var TotalCartPrice float64
	for _, data := range cart.CartData {
		TotalCartPrice += data.TotalPrice
	}

	orderID, err := o.orderRepo.OrderItems(order.UserID, order.AddressID, TotalCartPrice)
	if err != nil {
		return err
	}

	err = o.orderRepo.AddOrderProducts(orderID, cart.CartData)
	if err != nil {
		return err
	}
	for _, data := range cart.CartData {
		if err := o.cartUseCase.RemoveCart(order.UserID, data.ProductID); err != nil {
			return err
		}

	}
	return nil

}
