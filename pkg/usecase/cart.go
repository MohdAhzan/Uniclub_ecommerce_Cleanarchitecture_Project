package usecase

import (
	"errors"
	"fmt"
	interfaces "project/pkg/repository/interface"
	"project/pkg/utils/models"
)

type CartUseCase struct {
	CartRepo interfaces.CartRepository
	invRepo  interfaces.InventoryRepository
}

func NewCartUseCase(CartRepo interfaces.CartRepository, InvRepo interfaces.InventoryRepository) *CartUseCase {
	return &CartUseCase{
		CartRepo: CartRepo,
		invRepo:  InvRepo,
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
	fmt.Println("boolllll", exist)
	if exist {
		cartQuantity, err = u.CartRepo.FindCartQuantity(pid, cartID)
		if err != nil {
			return models.CartResponse{}, errors.New("error adding to cart")
		}

	}

	// creating cartItems for user
	newQuantity := cartQuantity + quantity
	fmt.Println("Updatedcart values", cartID, cartQuantity, quantity, newQuantity)
	errMsg := fmt.Sprintf("Can't add %d products as stock is %d", newQuantity, stock)
	if cartQuantity+quantity > stock {
		return models.CartResponse{}, errors.New(errMsg)
	}

	if cartQuantity == 0 {
		err = u.CartRepo.AddtoCartItems(cartID, pid)
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
	cartProduct.ProductName, _ = u.CartRepo.GetProductNames(pid)
	cartProduct.Image, _ = u.invRepo.GetProductImages(pid)
	cartProduct.Category_id, _ = u.invRepo.GetCategoryID(pid)
	cartProduct.Quantity, _ = u.CartRepo.FindCartQuantity(pid, cartID)

	cartProduct.StockAvailable = stock - cartProduct.Quantity
	price, _ := u.invRepo.FindPrice(pid)

	cartProduct.TotalPrice = price * float64(cartProduct.Quantity)

	var cartResponse models.CartResponse
	cartResponse.CartID = uint(cartID)
	cartResponse.CartData = append(cartResponse.CartData, cartProduct)

	return cartResponse, nil
}

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

		cart = append(cart, c)
	}

	var response models.CartResponse

	response.CartID = uint(cartID)
	response.CartData = cart

	return response, nil
}

func (u *CartUseCase) RemoveCart(userID, pid int) error {

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
