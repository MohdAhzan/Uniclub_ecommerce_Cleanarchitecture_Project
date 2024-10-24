package handler

import (
	"fmt"
	"net/http"
	services "github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/usecase/interface"
	response "github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/Response"
	"github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type InventaryHandler struct {
	InventoryUseCase services.InventoryUseCase
	Cache            *redis.Client
}

func NewInventoryHandler(usecase services.InventoryUseCase, redis *redis.Client) *InventaryHandler {

	return &InventaryHandler{
		InventoryUseCase: usecase,
		Cache:            redis,
	}
}

func (Inv *InventaryHandler) AddInventory(c *gin.Context) {

	var inventory models.AddInventory

	err := Inv.Cache.Del(Inv.Cache.Context(), "List_Home_products").Err()
	if err != nil {
		errMsg := response.ClientResponse(http.StatusBadRequest, "err deleting cache", nil, err.Error())
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	CategoryID, err := strconv.Atoi(c.Request.FormValue("category_id"))

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "form value error", nil, err.Error())
		c.JSON(400, errRes)
		return
	}

	ProductName := c.Request.FormValue("product_name")
	Size := c.Request.FormValue("size")
	Stock, err := strconv.Atoi(c.Request.FormValue("stock"))
	if err != nil {
		errRes := response.ClientResponse(400, "form value errror", nil, err.Error())
		c.JSON(400, errRes)
		return
	}
	Price, err := strconv.Atoi(c.Request.FormValue("price"))

	fmt.Println("price", Price)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "form value error", nil, err.Error())
		c.JSON(400, errRes)
		return
	}
	if Price < 0 {
		errRes := response.ClientResponse(400, "form value error", nil, "Invalid Price")
		c.JSON(400, errRes)
		return
	}

	inventory.CategoryID = CategoryID
	inventory.ProductName = ProductName
	inventory.Size = Size
	inventory.Stock = Stock
	inventory.Price = float64(Price)

	file, err := c.FormFile("image")
	if err != nil {
		errRes := response.ClientResponse(400, "error retrieving image", nil, err.Error())
		c.JSON(400, errRes)
		return
	}

	inventoryResponse, Err := Inv.InventoryUseCase.AddInventory(inventory, file)
	if Err != nil {
		errRes := response.ClientResponse(400, "error adding products to inventory", nil, Err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(200, "successfully added Inventory", inventoryResponse, nil)
	c.JSON(200, successRes)

}

func (Inv *InventaryHandler) GetProductsForAdmin(c *gin.Context) {

	productDetails, err := Inv.InventoryUseCase.GetProductsForAdmin()

	if err != nil {
		errRes := response.ClientResponse(400, "couldnt get product details for admin", nil, err.Error())
		c.JSON(400, errRes)
		return
	}
	successRes := response.ClientResponse(200, "successfully retrieved product details", productDetails, nil)
	c.JSON(200, successRes)
}

func (Inv *InventaryHandler) GetProductsForUsers(c *gin.Context) {

	///

	///
	productDetails, err := Inv.InventoryUseCase.GetProductsForUsers()

	if err != nil {
		errRes := response.ClientResponse(400, "couldnt get product details for admin", nil, err.Error())
		c.JSON(400, errRes)
		return
	}
	successRes := response.ClientResponse(200, "successfully retrieved product details", productDetails, nil)
	c.JSON(200, successRes)
}

func (Inv *InventaryHandler) DeleteInventory(c *gin.Context) {

	product_id := c.Query("id")

	pid, err := strconv.Atoi(product_id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error converting to int", nil, err.Error())
		c.JSON(400, errRes)
		return
	}
	delCmd := Inv.Cache.Del(Inv.Cache.Context(), "List_Home_products")
	res, _ := delCmd.Result()
	fmt.Println(res, "checkinggg delteeeeetekjlekjtsdkljfkasjflasjfkluasldkfjaklsjfijsduxcm,fuhzyzdxukfhjkjsdfxjlidfsdl")
	if err := delCmd.Err(); err != nil {
		errMsg := response.ClientResponse(http.StatusBadRequest, "err deleting cache", nil, err.Error())
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}
	Err := Inv.InventoryUseCase.DeleteInventory(pid)
	if Err != nil {
		errRes := response.ClientResponse(400, "fields provided are in wrong format", nil, Err.Error())
		c.JSON(400, errRes)
		return
	}
	successRes := response.ClientResponse(200, "successfully deleted the inventory", nil, nil)
	c.JSON(200, successRes)
}

func (inv *InventaryHandler) EditInventoryDetails(c *gin.Context) {

	productID := c.Query("id")

	pid, err := strconv.Atoi(productID)
	if err != nil {
		errRes := response.ClientResponse(400, "error converting the id", nil, err.Error())
		c.JSON(400, errRes)
		return
	}
	err = inv.Cache.Del(inv.Cache.Context(), "List_Home_products").Err()
	if err != nil {
		errMsg := response.ClientResponse(http.StatusBadRequest, "err deleting cache", nil, err.Error())
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}
	var model models.EditInventory

	Err := c.BindJSON(&model)
	if Err != nil {
		errRes := response.ClientResponse(400, "error binding model", nil, Err.Error())
		c.JSON(400, errRes)
		return
	}

	eRR := inv.InventoryUseCase.EditInventory(pid, model)
	if eRR != nil {
		errRes := response.ClientResponse(400, "error editing product check category id", nil, eRR.Error())
		c.JSON(400, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully edited product", nil, nil)
	c.JSON(200, successRes)
}

func (inv *InventaryHandler) SearchProducts(c *gin.Context) {

	pdtName := c.Query("search")

	searchedPdts, err := inv.InventoryUseCase.SearchProducts(pdtName)

	if err != nil {
		errRes := response.ClientResponse(400, "failed to Search Products", nil, err.Error())
		c.JSON(400, errRes)
		return
	}

	successRes := response.ClientResponse(200, "successfully searched the product", searchedPdts, nil)
	c.JSON(200, successRes)
}
