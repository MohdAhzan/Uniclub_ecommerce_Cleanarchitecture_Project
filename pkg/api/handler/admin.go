package handler

import (
	"fmt"
	"net/http"
	services "project/pkg/usecase/interface"
	response "project/pkg/utils/Response"
	models "project/pkg/utils/models"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminUseCase services.AdminUseCase
}

func NewAdminHandler(usecase services.AdminUseCase) *AdminHandler {
	return &AdminHandler{
		adminUseCase: usecase,
	}
}

// ADMIN_LOGIN

func (ad *AdminHandler) LoginHandler(c *gin.Context) {
	var adminDetails models.AdminLogin
	if err := c.BindJSON(&adminDetails); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "details not in the correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	admin, err := ad.adminUseCase.LoginHandler(adminDetails)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	c.Set("Access", admin.AccessToken)
	// c.Set("Refresh", admin.RefreshToken)

	successRes := response.ClientResponse(http.StatusBadRequest, "Admin authenticated succesfully", admin, nil)
	c.JSON(http.StatusOK, successRes)
}

// DISPLAY USERS

func (ad *AdminHandler) GetUsers(c *gin.Context) {

	users, err := ad.adminUseCase.GetUsers()
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't retrieve details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "successfully retrived the users", users, nil)
	c.JSON(http.StatusOK, successRes)
}

func (ad *AdminHandler) BlockUser(c *gin.Context) {
	id := c.Query("id")
	fmt.Println("\nTHE IN URL", id)
	err := ad.adminUseCase.BlockUser(id)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't block user", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return

	}

	successRes := response.ClientResponse(http.StatusOK, "successfully blocked the user ", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (ad *AdminHandler) UnBlockUser(c *gin.Context) {
	id := c.Query("id")
	err := ad.adminUseCase.UnBlockUser(id)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "couldn't block user", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRess := response.ClientResponse(http.StatusOK, "successfully unblocked the user", nil, nil)
	c.JSON(http.StatusOK, successRess)

}
