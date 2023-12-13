package handler

import (
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
	c.Set("Refresh", admin.RefreshToken)

	successRes := response.ClientResponse(http.StatusBadRequest, "Admin authenticated succesfully", admin, nil)
	c.JSON(http.StatusOK, successRes)
}
