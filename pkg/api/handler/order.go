package handler

import (
	"net/http"
	interfaces "project/pkg/usecase/interface"
	response "project/pkg/utils/Response"
	"project/pkg/utils/models"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderUseCase interfaces.OrderUseCase
}

func NewOrderHandler(orderUseCase interfaces.OrderUseCase) *OrderHandler {
	return &OrderHandler{

		orderUseCase: orderUseCase,
	}
}

func (u *OrderHandler) OrderFromCart(c *gin.Context) {

	var order models.Order

	err := c.BindJSON(&order)
	if err != nil {
		errMsg := response.ClientResponse(http.StatusBadRequest, "error binding JSON", nil, err.Error())
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	if err = u.orderUseCase.OrderFromCart(order); err != nil {
		errMsg := response.ClientResponse(400, "error placing order", nil, err.Error())
		c.JSON(400, errMsg)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully placed order", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
