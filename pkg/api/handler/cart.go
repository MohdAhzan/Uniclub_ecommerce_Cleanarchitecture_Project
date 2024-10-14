package handler

import (
	"fmt"
	"net/http"
	interfaces "github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/usecase/interface"
	response "github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/Response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	Cartusecase interfaces.CartUseCase
}

func NewCartHandler(usecase interfaces.CartUseCase) *CartHandler {
	return &CartHandler{
		Cartusecase: usecase,
	}
}

func (u *CartHandler) AddtoCart(c *gin.Context) {

	idString := c.Query("pid")
	pid, err := strconv.Atoi(idString)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Error string Conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	qtyString := c.Query("quantity")
	quantity, err := strconv.Atoi(qtyString)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Error string Conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if quantity <= 0 {
		Err := fmt.Errorf("incorrect quantity entered in the field")
		errRes := response.ClientResponse(http.StatusBadRequest, "Enter valid quantity", nil, Err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	UserID, _ := c.Get("id")

	cartResponse, err := u.Cartusecase.AddtoCart(pid, UserID.(int), quantity)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Error adding to cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added product to cart", cartResponse, nil)
	c.JSON(http.StatusOK, successRes)
}

func (u *CartHandler) GetCart(c *gin.Context) {
	userID, _ := c.Get("id")

	cartResponse, err := u.Cartusecase.GetCart(userID.(int))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Error fetching cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully fetched cart and products", cartResponse, nil)
	c.JSON(200, successRes)

}

func (u *CartHandler) RemoveCart(c *gin.Context) {

	pidString := c.Query("pid")

	pid, err := strconv.Atoi(pidString)
	if err != nil {
		errMsg := response.ClientResponse(http.StatusBadRequest, "ERror String conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	userID, _ := c.Get("id")

	err = u.Cartusecase.RemoveCart(userID.(int), pid)
	if err != nil {
		errMsg := response.ClientResponse(400, "Error  removing cart", nil, err.Error())
		c.JSON(400, errMsg)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully removed cart", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

func (u *CartHandler) DecreaseCartQuantity(c *gin.Context) {

	user_id, _ := c.Get("id")

	qtyString := c.Query("quantity")
	pdtString := c.Query("pid")
	quantity, err := strconv.Atoi(qtyString)
	if err != nil {

		errRes := response.ClientResponse(http.StatusBadRequest, "error converting into Integer", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	pID, err := strconv.Atoi(pdtString)
	if err != nil {

		errRes := response.ClientResponse(http.StatusBadRequest, "error converting into Integer", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	// if quantity > 0 {
	// 	err = fmt.Errorf("can't add quantity in DecreaseCartQuantity")
	// 	errREs := response.ClientResponse(http.StatusBadRequest, "Use AddtoCart for adding more quantity and Products :)", nil, err.Error())
	// 	c.JSON(http.StatusBadRequest, errREs)
	// 	return
	// }
	err = u.Cartusecase.DecreaseCartQuantity(user_id.(int), quantity, pID)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error Decreasing quantity", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusBadRequest, "successfully decreased quantity", nil, nil)
	c.JSON(http.StatusBadRequest, successRes)

}
