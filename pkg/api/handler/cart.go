package handler

import (
	"net/http"
	interfaces "project/pkg/usecase/interface"
	response "project/pkg/utils/Response"
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
