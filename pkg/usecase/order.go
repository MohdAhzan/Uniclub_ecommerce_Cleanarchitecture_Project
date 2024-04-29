package usecase

import (
	"errors"
	"fmt"
	helper "project/pkg/helper/interface"
	services "project/pkg/repository/interface"
	interfaces "project/pkg/usecase/interface"
	"project/pkg/utils/domain"
	"project/pkg/utils/models"
)

type OrderUseCase struct {
	orderRepo   services.OrderRepository
	cartRepo    services.CartRepository
	cartUseCase interfaces.CartUseCase
	userRepo    services.UserRepository
	helper      helper.Helper
	couponRepo  services.CouponRepository
}

func NewOrderUseCase(orderRepo services.OrderRepository, cartRepo services.CartRepository, cartUseCase interfaces.CartUseCase, userRepo services.UserRepository, h helper.Helper, c services.CouponRepository) *OrderUseCase {
	return &OrderUseCase{
		orderRepo:   orderRepo,
		cartRepo:    cartRepo,
		cartUseCase: cartUseCase,
		userRepo:    userRepo,
		helper:      h,
		couponRepo:  c,
	}
}
func (o OrderUseCase) OrderFromCart(order models.Order, couponID int) error {

	cart, err := o.cartUseCase.GetCart(order.UserID)
	if err != nil {
		return err
	}

	fmt.Println("model GET CART", cart)
	var Total float64
	for _, data := range cart.CartData {

		Total += data.DiscountedPrice

	}

	used, err := o.couponRepo.CheckIfUserUsedCoupon(order.UserID, couponID)
	if err != nil {
		return err
	}
	if used {
		return errors.New("this coupon is alreay used")
	}

	if !used {

		coupon, err := o.couponRepo.FindCouponDetails(couponID)
		if err != nil {
			return err
		}

		totalDiscount := (Total * float64(coupon.DiscountRate)) / 100
		Total = Total - totalDiscount
	}
	if Total == 0 {
		return errors.New("no items in Cart")
	}

	orderID, err := o.orderRepo.OrderItems(order.UserID, order.AddressID, order.PaymentID, couponID, Total)
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

func (o OrderUseCase) Checkout(userID int) (models.CheckOut, error) {

	var orderDetails models.CheckOut

	cart_id, err := o.cartRepo.GetCartID(userID)
	if err != nil {
		return models.CheckOut{}, err
	}

	// o.orderRepo.GetOrderedProducts()

	cart, err := o.cartUseCase.GetCart(userID)
	if err != nil {
		return models.CheckOut{}, err
	}
	// fmt.Println("cart", cart)

	address, err := o.cartRepo.GetCartAddress(userID)
	if err != nil {
		return models.CheckOut{}, err
	}
	fmt.Println(address)
	var TotalCartPrice float64
	for _, data := range cart.CartData {
		TotalCartPrice += data.TotalPrice
	}

	// fmt.Println(cart_id, "\n", address, "\n", cart.CartData, "\n", TotalCartPrice)
	orderDetails.CartID = cart_id
	orderDetails.Addresses = address
	orderDetails.Products = cart.CartData
	orderDetails.TotalPrice = TotalCartPrice

	// fmt.Println(orderDetails)
	return orderDetails, nil
}

func (o OrderUseCase) GetOrders(userID int) ([]domain.OrderDetailsWithImages, error) {

	orders, err := o.orderRepo.GetOrders(userID)
	if err != nil {
		return []domain.OrderDetailsWithImages{}, err
	}

	var result []domain.OrderDetailsWithImages
	for _, data := range orders {
		var or domain.OrderDetailsWithImages
		images, err := o.orderRepo.GetOrderImages(int(data.ID))
		if err != nil {
			return []domain.OrderDetailsWithImages{}, err
		}
		paymentMethod, err := o.orderRepo.GetPaymentMethodsByID(int(data.PaymentMethodID))
		if err != nil {
			return []domain.OrderDetailsWithImages{}, err
		}
		inventoryIDs, err := o.orderRepo.GetOrderProductIDs(int(data.ID))
		if err != nil {
			return []domain.OrderDetailsWithImages{}, err
		}
		or.OrderDetails = data
		or.ProductID = inventoryIDs
		or.Images = images
		or.PaymentMethod = paymentMethod
		result = append(result, or)
	}

	return result, nil
}

func (o OrderUseCase) GetOrderDetailsByOrderID(orderID, userID int) (domain.OrderDetails, error) {

	user, err := o.userRepo.GetUserDetails(userID)
	if err != nil {
		return domain.OrderDetails{}, err
	}

	address, orderData, err := o.orderRepo.GetOrderAddress(orderID)
	if err != nil {
		return domain.OrderDetails{}, err
	}

	errMSg := fmt.Sprintf("No OrderID of %d", orderID)

	if address.Address == "" {
		return domain.OrderDetails{}, errors.New(errMSg)
	}

	var orderDetails domain.OrderDetails

	orderDetails.ID = orderID
	orderDetails.Username = user.Name
	orderDetails.Address = address
	orderDetails.OrderStatus = orderData.Order_status

	paymentMethod, err := o.orderRepo.GetPaymentMethodsByID(orderData.Payment_method_id)
	if err != nil {
		return domain.OrderDetails{}, err
	}

	orderDetails.PaymentMethod = paymentMethod
	orderDetails.Total = orderData.Final_Price
	orderDetails.PaymentStatus = orderData.PaymentStatus

	fmt.Println(orderDetails)
	return orderDetails, nil
}

func (o OrderUseCase) CancelOrder(orderID, userID int) error {

	status, err := o.orderRepo.CheckOrderStatusByID(orderID)
	if err != nil {
		return err
	}
	msg := fmt.Sprintf("Order can't be cancelled as it is %s Kindly return the product", status)

	if status == "CANCELED" {
		return errors.New("the order is already cancelled")
	} else if status != "PENDING" {
		return errors.New(msg)
	}

	err = o.orderRepo.CancelOrder(orderID)
	if err != nil {
		return err
	}

	// only refund if the user has already prepaid orderAmount

	paymentStatus, err := o.orderRepo.GetPaymentStatusByID(orderID)
	if err != nil {
		return err
	}

	if paymentStatus == "PAID" {

		orderAmount, err := o.orderRepo.FindOrderAmount(orderID)
		if err != nil {
			return err
		}
		var model models.AddMoneytoWallet

		model.UserID = userID
		model.Amount = orderAmount
		model.TranscationType = "PDT_CANCELLED"

		err = o.userRepo.AddMoneytoWallet(model)
		if err != nil {
			return err
		}
	}

	return nil
}

func (o OrderUseCase) ReturnOrder(orderID, userID int) error {

	status, err := o.orderRepo.CheckOrderStatusByID(orderID)

	if err != nil {
		return err
	}
	msg := fmt.Sprintf("cannot return the order ,already %s", status)

	if status == "RETURNED" {
		return errors.New("cannot return the order, already returned ")
	} else if status != "DELIVERED" {
		return errors.New(msg)
	}

	err = o.orderRepo.ReturnOrder(orderID)
	if err != nil {
		return err

	}

	user, err := o.userRepo.GetUserDetails(userID)
	if err != nil {
		return err
	}
	Sub := "ORDER RETURN REQUEST"
	mailMsg := fmt.Sprintf("Dear %s your request for returning the product of ID: %d is on process", user.Name, orderID)
	err = o.helper.SendMailToPhone(user.Email, Sub, mailMsg)
	if err != nil {
		return err
	}

	return nil

}
