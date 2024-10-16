package usecase

import (
	"errors"
	"fmt"
	interfaces "github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/repository/interface"
	"github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/models"
)

type CartUseCase struct {
	CartRepo interfaces.CartRepository
	invRepo  interfaces.InventoryRepository
	offRepo  interfaces.OfferRepository
}

func NewCartUseCase(CartRepo interfaces.CartRepository, InvRepo interfaces.InventoryRepository, off interfaces.OfferRepository) *CartUseCase {
	return &CartUseCase{
		CartRepo: CartRepo,
		invRepo:  InvRepo,
		offRepo:  off,
	}
}

func (u CartUseCase) AddtoCart(pid, userID, quantity int) (models.CartResponse, error) {

	stock, err := u.invRepo.CheckStock(pid)
	if err != nil {
		return models.CartResponse{}, errors.New("error fetching stock")
	}
	if stock < 1 {
		return models.CartResponse{}, errors.New("out of stock")
	}
	quantityError := fmt.Sprintf("Can't add %d products as stock is %d", quantity, stock)
	if quantity > stock {
		return models.CartResponse{}, errors.New(quantityError)
	}

	cartID, err := u.CartRepo.GetCartID(userID)
	if err != nil {
		return models.CartResponse{}, err
	}

	if cartID == 0 {
		// if no cart exists create Newcart
		cartID, err = u.CartRepo.CreateNewCart(userID)
		if err != nil {
			return models.CartResponse{}, errors.New("can't create cart for user")
		}

	}

	var cartQuantity int
	exist, Err := u.CartRepo.CheckIfItemIsAlreadyAdded(cartID, pid)
	if Err != nil {
		return models.CartResponse{}, Err
	}
	if exist {
		cartQuantity, err = u.CartRepo.FindCartQuantity(pid, cartID)
		if err != nil {
			return models.CartResponse{}, errors.New("error adding to cart")
		}

	}

	// creating cartItems for user
	newQuantity := cartQuantity + quantity

	errMsg := fmt.Sprintf("Can't add %d products as stock is %d", newQuantity, stock)
	if cartQuantity+quantity > stock {
		return models.CartResponse{}, errors.New(errMsg)
	}

	if cartQuantity == 0 {
		err = u.CartRepo.AddtoCartItems(cartID, pid, quantity)
		if err != nil {
			return models.CartResponse{}, errors.New("failed to AddCartItems")
		}
	} else {
		err := u.CartRepo.UpdateCartQuantity(cartID, pid, newQuantity)
		if err != nil {
			return models.CartResponse{}, err
		}
	}
	var cartProduct models.GetCart
	cartProduct.ProductID = pid
	cartProduct.ProductName, err = u.CartRepo.GetProductNames(pid)
	if err != nil {
		return models.CartResponse{}, err
	}

	cartProduct.Image, err = u.invRepo.GetProductImages(pid)
	if err != nil {
		return models.CartResponse{}, err
	}
	cartProduct.Category_id, err = u.invRepo.GetCategoryID(pid)
	if err != nil {
		return models.CartResponse{}, err
	}
	cartProduct.Quantity, err = u.CartRepo.FindCartQuantity(pid, cartID)
	if err != nil {
		return models.CartResponse{}, err
	}

	cartProduct.StockAvailable = stock - cartProduct.Quantity

	price, err := u.invRepo.FindPrice(pid)
	if err != nil {
		return models.CartResponse{}, err
	}

	cartProduct.TotalPrice = price * float64(cartProduct.Quantity)
	//check if any offers are there

	// if the category id of these products are in offer table discount the price to new one

	CategoryDiscountRate, CategoryOffer, err := u.offRepo.GetCategoryOfferDiscountPercentage(cartProduct.Category_id)
	if err != nil {
		return models.CartResponse{}, err
	}
	cartProduct.CategoryOffer = CategoryOffer
	ProductDiscountRate, ProductOffer, err := u.offRepo.GetInventoryOfferDiscountPercentage(cartProduct.ProductID)
	if err != nil {
		return models.CartResponse{}, err
	}
	cartProduct.ProductOffer = ProductOffer
	DiscountRate := CategoryDiscountRate + ProductDiscountRate

	//Discounted Price = Original Price - (Original Price * (Discount Percentage / 100))

	cartProduct.DiscountedPrice = cartProduct.TotalPrice - (cartProduct.TotalPrice * (DiscountRate / 100))

	var cartResponse models.CartResponse
	cartResponse.CartID = uint(cartID)
	cartResponse.CartData = append(cartResponse.CartData, cartProduct)

	return cartResponse, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func (u CartUseCase) GetCart(userID int) (models.CartResponse, error) {

	//get cartId
	cartID, err := u.CartRepo.GetCartID(userID)
	if err != nil {
		return models.CartResponse{}, err
	}
	//get all productsID corresponse to cartId

	pIDs, err := u.CartRepo.GetProductIDs(cartID)
	if err != nil {
		return models.CartResponse{}, err
	}

	//get all product_name corresponse to cartID
	var productNames []string
	for i := range pIDs {

		pName, err := u.CartRepo.GetProductNames(pIDs[i])
		if err != nil {
			return models.CartResponse{}, err
		}
		productNames = append(productNames, pName)

	}

	var images []string
	var categoryID []int

	for i := range pIDs {

		//get all product images in cart
		imageUrl, err := u.invRepo.GetProductImages(pIDs[i])
		if err != nil {
			return models.CartResponse{}, err
		}

		images = append(images, imageUrl)

		//get all product categories in cart

		cID, err := u.invRepo.GetCategoryID(pIDs[i])
		if err != nil {
			return models.CartResponse{}, err
		}
		categoryID = append(categoryID, cID)
	}

	var quantity []int
	var stock []int

	for i := range pIDs {

		// get all product quantity in caarttttt
		q, err := u.CartRepo.FindCartQuantity(pIDs[i], cartID)
		if err != nil {
			return models.CartResponse{}, err
		}
		quantity = append(quantity, q)

		// get all stock of products in cart
		s, err := u.invRepo.FindStock(pIDs[i])
		if err != nil {
			return models.CartResponse{}, err
		}

		stock = append(stock, s)

	}
	var price []float64
	for i := range pIDs {

		//get all product price in cart

		p, err := u.invRepo.FindPrice(pIDs[i])
		if err != nil {
			return models.CartResponse{}, err
		}
		price = append(price, p*float64(quantity[i]))

	}
	var cart []models.GetCart
	for i := range pIDs {
		var c models.GetCart
		c.ProductID = pIDs[i]
		c.ProductName = productNames[i]
		c.Image = images[i]
		c.Category_id = categoryID[i]
		c.Quantity = quantity[i]
		c.StockAvailable = stock[i] - quantity[i]
		c.TotalPrice = price[i]

		//check if any offers are there

		// if the category id of these products are in offer table discount the price to new one

		CategoryDiscountRate, CategoryOffer, err := u.offRepo.GetCategoryOfferDiscountPercentage(c.Category_id)
		if err != nil {
			return models.CartResponse{}, err
		}

		c.CategoryOffer = CategoryOffer
		ProductDiscountRate, ProductOffer, err := u.offRepo.GetInventoryOfferDiscountPercentage(c.ProductID)
		if err != nil {
			return models.CartResponse{}, err
		}
		c.ProductOffer = ProductOffer
		DiscountRate := CategoryDiscountRate + ProductDiscountRate
		//Discounted Price = Original Price - (Original Price * (Discount Percentage / 100))

		c.DiscountedPrice = c.TotalPrice - (c.TotalPrice * (DiscountRate / 100))

		cart = append(cart, c)
	}

	var response models.CartResponse

	response.CartID = uint(cartID)
	response.CartData = cart

	return response, nil
}

func (u CartUseCase) RemoveCart(userID, pid int) error {

	cartID, err := u.CartRepo.GetCartID(userID)
	if err != nil {
		return err
	}

	err = u.CartRepo.RemoveCartItems(pid, cartID)
	if err != nil {
		return err
	}

	return nil
}

func (u CartUseCase) DecreaseCartQuantity(userID, quantity, pID int) error {

	if quantity < 0 {
		return errors.New("quantity cant be negative")
	}

	cartID, err := u.CartRepo.GetCartID(userID)

	if err != nil {
		return err
	}

	cartQuantity, err := u.CartRepo.FindCartQuantity(pID, cartID)
	if err != nil {
		return err
	}
	fmt.Println("quantity", quantity)

	updatedQuantity := cartQuantity - quantity
	fmt.Println("updatedquantity", updatedQuantity)
	fmt.Println("cartquantity", cartQuantity)
	if quantity > cartQuantity {

		err := fmt.Errorf("cant decrease cartquantity by %d cartquantity is only %d", quantity, cartQuantity)

		return err
	}

	if updatedQuantity == 0 {

		err = u.CartRepo.RemoveCartItems(pID, cartID)
		if err != nil {
			errMsg := fmt.Errorf("product of id %d is removed from the cart", pID)
			return errMsg
		}
	}

	err = u.CartRepo.UpdateCartQuantity(cartID, pID, updatedQuantity)
	if err != nil {
		return err
	}

	return nil

}
