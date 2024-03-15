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
		errRes := response.ClientResponse(http.StatusBadRequest, "Error string Conversion", nil, err.Error)
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	UserID, _ := c.Get("id")

	cartResponse, err := u.Cartusecase.AddtoCart(pid, UserID.(int))

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
