package handler

import (
	"fmt"
	"net/http"
	services "project/pkg/usecase/interface"
	response "project/pkg/utils/Response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	payUseCase services.PaymentUseCase
}

func NewPaymentHandler(payUsecase services.PaymentUseCase) *PaymentHandler {
	return &PaymentHandler{
		payUseCase: payUsecase,
	}
}

func (p *PaymentHandler) MakePaymentFromRazorPay(c *gin.Context) {

	order := c.Query("order_id")
	userid := c.Query("user_id")

	orderID, err := strconv.Atoi(order)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error string conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	userID, err := strconv.Atoi(userid)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error string conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	body, err := p.payUseCase.MakePaymentFromRazorPay(userID, orderID)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	// var d models.OrderPaymentDetails
	fmt.Println("checkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkk", body, "\n order", orderID, userID)

	c.HTML(http.StatusOK, "razorpay.html", body)

}
func (p *PaymentHandler) VerifyPaymentFromRazorPay(c *gin.Context) {

	// fmt.Println("verifaction workdnig")
	razorOrderID := c.Query("order_id")

	razorPaymentID := c.Query("payment_id")

	razorID := c.Query("razor_id")

	fmt.Println("and this is order ID", razorOrderID, "paymnt id", razorPaymentID, "razorID", razorID)

	err := p.payUseCase.VerifyPaymentFromRazorPay(razorPaymentID, razorOrderID, razorID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not update payment details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully updated payment details", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

func (p *PaymentHandler) PaymentFromWallet(c *gin.Context) {

	useridStr := c.Query("user_id")
	orderIdString := c.Query("order_id")

	OrderID, err := strconv.Atoi(orderIdString)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "error converting orderID to string please enter in valid format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	userID, err := strconv.Atoi(useridStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "error converting orderID to string please enter in valid format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	body, err := p.payUseCase.PaymentFromWallet(OrderID, userID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "error while paying from wallet", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	iS := 0
	if body.FinalPrice == float64(iS) {
		successRes := response.ClientResponse(http.StatusOK, "Successfully paid using wallet", nil, nil)
		c.JSON(http.StatusOK, successRes)
		return
	}

	c.HTML(http.StatusOK, "razorpay.html", body)
}
