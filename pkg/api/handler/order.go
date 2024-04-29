package handler

import (
	"net/http"
	interfaces "project/pkg/usecase/interface"
	response "project/pkg/utils/Response"
	"project/pkg/utils/models"
	"strconv"

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

func (o *OrderHandler) Checkout(c *gin.Context) {

	userID, _ := c.Get("id")

	orderDetails, err := o.orderUseCase.Checkout(userID.(int))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error checkout", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully fetched Checkout Details", orderDetails, nil)

	c.JSON(http.StatusOK, successRes)

}

func (o *OrderHandler) GetOrders(c *gin.Context) {
	id, _ := c.Get("id")

	OrderDetails, err := o.orderUseCase.GetOrders(id.(int))
	if err != nil {
		errRes := response.ClientResponse(400, "Couldnt fetch order details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(200, "successfully fetched orderDetails", OrderDetails, nil)
	c.JSON(200, successRes)

}

func (o *OrderHandler) GetOrderDetailsByOrderID(c *gin.Context) {

	userID, _ := c.Get("id")

	idString := c.Query("order_id")

	orderID, err := strconv.Atoi(idString)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error string conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	// var order domain.OrderDetails

	orderDetails, err := o.orderUseCase.GetOrderDetailsByOrderID(orderID, userID.(int))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error fetching orderDetails", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully fetched orderdetails", orderDetails, nil)
	c.JSON(http.StatusOK, successRes)

}

func (o *OrderHandler) CancelOrder(c *gin.Context) {

	idString := c.Query("order_id")

	orderID, err := strconv.Atoi(idString)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in string conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	user_id, _ := c.Get("id")

	err = o.orderUseCase.CancelOrder(orderID, user_id.(int))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "failed to cancel the order ", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	succesRes := response.ClientResponse(http.StatusOK, "This order has been cancelled", nil, nil)
	c.JSON(200, succesRes)
}

func (o *OrderHandler) ReturnOrder(c *gin.Context) {
	uID, _ := c.Get("id")

	idString := c.Query("order_id")

	orderID, err := strconv.Atoi(idString)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in string conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	userID := uID.(int)

	err = o.orderUseCase.ReturnOrder(orderID, userID)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "failed to return the order", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "This Order has been Requested for Return", nil, nil)
	c.JSON(http.StatusOK, successRes)

}
