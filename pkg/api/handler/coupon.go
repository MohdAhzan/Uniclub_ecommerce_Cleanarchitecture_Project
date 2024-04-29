package handler

import (
	"net/http"
	interfaces "project/pkg/usecase/interface"
	response "project/pkg/utils/Response"
	"project/pkg/utils/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CouponHandler struct {
	couponUseCase interfaces.CouponUseCase
}

func NewCouponHandler(coupon interfaces.CouponUseCase) *CouponHandler {
	return &CouponHandler{
		couponUseCase: coupon,
	}
}

func (coup *CouponHandler) CreateNewCoupon(c *gin.Context) {
	var coupon models.Coupons
	if err := c.BindJSON(&coupon); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	err := coup.couponUseCase.CreateNewCoupon(coupon)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not add the Coupon", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added Coupon", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (coup *CouponHandler) GetAllCoupons(c *gin.Context) {

	couponData, err := coup.couponUseCase.GetAllCoupons()
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could fetch Coupon Details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully fetched all coupon details", couponData, nil)
	c.JSON(http.StatusOK, successRes)
}

func (coup *CouponHandler) MakeCouponInvalid(c *gin.Context) {

	idStr := c.Query("coupon_id")
	couponID, err := strconv.Atoi(idStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "error in coupon Query conversion please input a valid query ", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	err = coup.couponUseCase.MakeCouponInvalid(couponID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "failed to make Coupon Invalid", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully make Coupon Invalid", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (coup *CouponHandler) MakeCouponValid(c *gin.Context) {

	idStr := c.Query("coupon_id")
	couponID, err := strconv.Atoi(idStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "error in coupon Query conversion please input a valid query ", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	err = coup.couponUseCase.MakeCouponValid(couponID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "failed to make Coupon Valid", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully make Coupon Valid", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
