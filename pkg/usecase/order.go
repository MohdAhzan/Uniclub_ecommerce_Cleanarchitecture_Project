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

	// cartID, err := o.cartRepo.GetCartID(order.UserID)
	// if err != nil {
	// 	return err
	// }

	cart, err := o.cartUseCase.GetCart(order.UserID)
	if err != nil {
		return err
	}
	var TotalCartPrice float64
	for _, data := range cart.CartData {
		TotalCartPrice += data.TotalPrice
	}

	orderID, err := o.orderRepo.OrderItems(order.UserID, order.AddressID, TotalCartPrice)
	if err != nil {
		return err
	}
	fmt.Println(orderID)
	return nil

}
