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

	err=u.orderUseCase.OrderFromCart(order)

}
