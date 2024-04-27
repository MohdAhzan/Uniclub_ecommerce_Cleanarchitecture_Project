package handler

import (
	"net/http"
	interfaces "project/pkg/usecase/interface"
	response "project/pkg/utils/Response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WishlistHandler struct {
	wishusecase interfaces.WishlistUsecase
}

func NewWishlistHandler(w interfaces.WishlistUsecase) *WishlistHandler {
	return &WishlistHandler{
		wishusecase: w,
	}
}

func (w *WishlistHandler) AddToWishlist(c *gin.Context) {

	pidStrg := c.Query("pid")

	pID, err := strconv.Atoi(pidStrg)
	if err != nil {
		errREs := response.ClientResponse(http.StatusBadRequest, "error converting string ", nil, err.Error())
		c.JSON(http.StatusBadRequest, errREs)
		return
	}

	user_id, _ := c.Get("id")

	err = w.wishusecase.AddToWishlist(user_id.(int), pID)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error adding product to wishlist", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "succesfully added to wishlist", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

func (w *WishlistHandler) GetWishlist(c *gin.Context) {

	user_id, _ := c.Get("id")

	wishlistData, err := w.wishusecase.GetWishlist(user_id.(int))
	if err != nil {
		errREs := response.ClientResponse(http.StatusBadRequest, "error fetching wishlist details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errREs)
		return

	}

	successREs := response.ClientResponse(http.StatusOK, "successfully fetched wishlist product details", wishlistData, nil)
	c.JSON(http.StatusOK, successREs)

}

func (w *WishlistHandler) RemoveFromWishlist(c *gin.Context) {

	user_id, _ := c.Get("id")

	pdtString := c.Query("pid")

	pID, err := strconv.Atoi(pdtString)
	if err != nil {
		errREs := response.ClientResponse(http.StatusBadRequest, "error in string conversion", nil, err.Error())
		c.JSON(http.StatusBadRequest, errREs)
		return
	}

	err = w.wishusecase.RemoveFromWishlist(user_id.(int), pID)
	if err != nil {
		errREs := response.ClientResponse(http.StatusBadRequest, "error removing item from wishlist", nil, err.Error())
		c.JSON(http.StatusBadRequest, errREs)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully removed wishlisht", nil, nil)
	c.JSON(http.StatusOK, successRes)

}
