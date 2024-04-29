package handler

import (
	"net/http"
	interfaces "project/pkg/usecase/interface"
	response "project/pkg/utils/Response"
	"project/pkg/utils/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OfferHandler struct {
	offerUsecase interfaces.OfferUsecase
}

func NewOfferHandler(offUsecase interfaces.OfferUsecase) *OfferHandler {
	return &OfferHandler{
		offerUsecase: offUsecase,
	}
}

func (o *OfferHandler) AddCategoryOffer(c *gin.Context) {

	var offerModel models.AddCategoryOffer

	err := c.BindJSON(&offerModel)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error parsing json", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err = o.offerUsecase.AddCategoryOffer(offerModel)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error adding category offer", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return

	}

	successRes := response.ClientResponse(http.StatusOK, "successfully added category Offer", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

func (o *OfferHandler) GetAllCategoryOffers(c *gin.Context) {

	catOffers, err := o.offerUsecase.GetAllCategoryOffers()
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error fetching category offers", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return

	}

	successRes := response.ClientResponse(http.StatusOK, "successfully fetched category Offers", catOffers, nil)
	c.JSON(http.StatusOK, successRes)

}

func (o *OfferHandler) EditCategoryOffer(c *gin.Context) {

	newDis := c.Query("new_discount")
	newDiscount, err := strconv.ParseFloat(newDis, 64)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in string conversion into float64", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	cid := c.Query("category_id")
	categoryID, err := strconv.Atoi(cid)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in string conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err = o.offerUsecase.EditCategoryOffer(newDiscount, categoryID)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error editing category offer", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully edited category Offer", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

func (o *OfferHandler) ValidorInvalidCategoryOffers(c *gin.Context) {

	statusString := c.Query("status")
	status, err := strconv.ParseBool(statusString)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in string conversion into bool enter valid query parameter", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	cid := c.Query("category_id")
	categoryID, err := strconv.Atoi(cid)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error in string conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err = o.offerUsecase.ValidorInvalidCategoryOffers(status, categoryID)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error editing category offer", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully edited category Offer", nil, nil)
	c.JSON(http.StatusOK, successRes)

}
