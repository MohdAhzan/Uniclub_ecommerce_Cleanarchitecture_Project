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
	offerRepo   services.OfferRepository
	invRepo     services.InventoryRepository
}

func NewOrderUseCase(orderRepo services.OrderRepository, cartRepo services.CartRepository, cartUseCase interfaces.CartUseCase, userRepo services.UserRepository, h helper.Helper, c services.CouponRepository, o services.OfferRepository, inv services.InventoryRepository) *OrderUseCase {
	return &OrderUseCase{
		orderRepo:   orderRepo,
		cartRepo:    cartRepo,
		cartUseCase: cartUseCase,
		userRepo:    userRepo,
		helper:      h,
		couponRepo:  c,
		offerRepo:   o,
		invRepo:     inv,
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

		Total = data.DiscountedPrice + Total
		fmt.Println("total according to cart", Total)

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

func (o OrderUseCase) Checkout(userID, couponID int) (models.CheckOut, error) {

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

	//final price after applying coupon if any
	var Final_Price float64
	Final_Price = TotalCartPrice
	used, err := o.couponRepo.CheckIfUserUsedCoupon(userID, couponID)
	if err != nil {
		return models.CheckOut{}, err
	}
	if used {
		return models.CheckOut{}, errors.New("this coupon is alreay used")
	}

	if !used {

		coupon, err := o.couponRepo.FindCouponDetails(couponID)
		if err != nil {
			return models.CheckOut{}, err
		}

		totalDiscount := (Final_Price * float64(coupon.DiscountRate)) / 100
		Final_Price = Final_Price - totalDiscount
	}

	// fmt.Println(cart_id, "\n", address, "\n", cart.CartData, "\n", TotalCartPrice)
	orderDetails.CartID = cart_id
	orderDetails.Addresses = address
	orderDetails.Products = cart.CartData
	orderDetails.TotalPrice = TotalCartPrice
	orderDetails.Final_Price = Final_Price

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

func (o OrderUseCase) GetEachProductOrderDetails(orderID, userID int) (domain.OrderDetailsSeparate, error) {

	var AllData domain.OrderDetailsSeparate

	AllData.ID = userID

	user, err := o.userRepo.GetUserDetails(userID)
	if err != nil {
		return domain.OrderDetailsSeparate{}, err
	}

	AllData.Username = user.Name

	//get address used by user for order
	orderAddress, orderData, err := o.orderRepo.GetOrderAddress(orderID)
	if err != nil {
		return domain.OrderDetailsSeparate{}, err
	}

	AllData.Address = orderAddress
	//get order Status
	AllData.OrderStatus = orderData.Order_status
	//get payment method
	paymentMethod, err := o.orderRepo.GetPaymentMethodsByID(orderData.Payment_method_id)
	if err != nil {
		return domain.OrderDetailsSeparate{}, err
	}
	AllData.PaymentMethod = paymentMethod

	individualOrders, err := o.orderRepo.GetAllOrderItemsByOrderID(orderID)
	if err != nil {
		return domain.OrderDetailsSeparate{}, err
	}

	fmt.Println(individualOrders)

	var AllIndividualOrders []domain.EachProductOrderDetails

	for _, value := range individualOrders {
		var model domain.EachProductOrderDetails
		model.ProductID = uint(value.ProductID)
		model.Quantity = uint(value.Quantity)
		model.ProductPrice = value.ProductPrice

		CategoryID, err := o.invRepo.GetCategoryID(int(value.ProductID))
		if err != nil {
			return domain.OrderDetailsSeparate{}, err
		}
		catOFF, catOfferName, err := o.offerRepo.GetCategoryOfferDiscountPercentage(CategoryID)
		if err != nil {
			return domain.OrderDetailsSeparate{}, err
		}

		pdtOFF, pdtOfferName, err := o.offerRepo.GetInventoryOfferDiscountPercentage(int(value.ProductID))
		if err != nil {
			return domain.OrderDetailsSeparate{}, err
		}

		totalOFF := catOFF + pdtOFF

		totalDiscount := (value.ProductPrice * totalOFF) / 100
		value.DiscountPrice = value.ProductPrice - totalDiscount

		model.DiscountPrice = value.DiscountPrice
		OfferName := catOfferName + "," + pdtOfferName
		model.ProductOffer = OfferName
		AllIndividualOrders = append(AllIndividualOrders, model)
	}

	orderPaymentStatus, err := o.orderRepo.GetPaymentStatusByID(orderID)
	if err != nil {
		return domain.OrderDetailsSeparate{}, err
	}
	AllData.PaymentStatus = orderPaymentStatus

	AllData.EachProductOrderDetails = AllIndividualOrders

	AllData.Total = orderData.Final_Price

	return AllData, nil
}

func (o OrderUseCase) CancelProductInOrder(orderID, pID, user_id int) (domain.OrderDetailsSeparate, error) {

	err := o.orderRepo.CheckOrderByID(orderID)
	if err != nil {
		return domain.OrderDetailsSeparate{}, err
	}

	status, err := o.orderRepo.CheckOrderStatusByID(orderID)
	if err != nil {
		return domain.OrderDetailsSeparate{}, err
	}
	if status == "CANCELED" {
		return domain.OrderDetailsSeparate{}, errors.New("this order is already cancelled")
	}
	if status == "RETURNED" {
		return domain.OrderDetailsSeparate{}, errors.New("this order is already Returned")

	}
	if status == "DELIVERED" {
		return domain.OrderDetailsSeparate{}, errors.New("this order is already DELIVERED ")

	}

	count, err := o.orderRepo.CheckIndividualOrders(orderID, pID)
	if err != nil {
		return domain.OrderDetailsSeparate{}, err
	}

	if count < 1 {
		return domain.OrderDetailsSeparate{}, errors.New("no product is in this order ID ")
	}

	//delete from order_items
	//take that product order amount from it

	productPrice, err := o.orderRepo.DeleteProductInOrder(orderID, pID)
	if err != nil {
		return domain.OrderDetailsSeparate{}, err
	}

	//apply the discount to make the product price same as when user ordered ...

	categoryID, err := o.invRepo.GetCategoryID(pID)
	if err != nil {
		return domain.OrderDetailsSeparate{}, err
	}
	catOFF, _, err := o.offerRepo.GetCategoryOfferDiscountPercentage(categoryID)
	if err != nil {
		return domain.OrderDetailsSeparate{}, err
	}

	pdtOFF, _, err := o.offerRepo.GetInventoryOfferDiscountPercentage(pID)
	if err != nil {
		return domain.OrderDetailsSeparate{}, err
	}
	totalOFF := catOFF + pdtOFF
	pdtDiscountPrice := (productPrice * totalOFF) / 100
	newProductPrice := productPrice - pdtDiscountPrice

	OrderPrice, err := o.orderRepo.FindOrderAmount(orderID)
	if err != nil {
		return domain.OrderDetailsSeparate{}, err
	}

	//reduce the product order price from total order price
	NewPrice := OrderPrice - newProductPrice
	err = o.orderRepo.UpdateFinalOrderPrice(orderID, NewPrice)
	if err != nil {
		return domain.OrderDetailsSeparate{}, err
	}

	// check if user paid the order if paid transfer the amount to his wallet
	paymentStatus, err := o.orderRepo.GetPaymentStatusByID(orderID)
	if err != nil {
		return domain.OrderDetailsSeparate{}, err
	}
	if paymentStatus == "PAID" {
		var wallet models.AddMoneytoWallet

		wallet.Amount = newProductPrice
		wallet.UserID = user_id
		wallet.TranscationType = "PDT_CANCELLED"
		err := o.userRepo.AddMoneytoWallet(wallet)
		if err != nil {
			return domain.OrderDetailsSeparate{}, err
		}

	}

	orderData, err := o.GetEachProductOrderDetails(orderID, user_id)
	if err != nil {
		return domain.OrderDetailsSeparate{}, err
	}

	return orderData, nil
}
