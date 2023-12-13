package handler

import (
	"net/http"
	services "project/pkg/usecase/interface"
	response "project/pkg/utils/Response"
	"project/pkg/utils/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	userUseCase services.UserUseCase
}
type Response struct {
	ID      uint   `copier:"must"`
	Name    string `copier:"must"`
	Surname string `copier:"must"`
}

func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

func (u *UserHandler) UserSignUp(c *gin.Context) {
	var user models.UserDetails

	//binding the userDetails from Client to Struct

	if err := c.BindJSON(&user); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format blahblahblah", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	//checking if the details sent by the client has correct constraints

	err := validator.New().Struct(user)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "constraints provided are not correct as of server", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	// sending the struct from client to the

	userCreated, err := u.userUseCase.UserSignup(user)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "user couldnt signup", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	succesRes := response.ClientResponse(http.StatusCreated, "User successfully created", userCreated, nil)
	c.JSON(http.StatusCreated, succesRes)

}
