package handler

import (
	"fmt"
	"net/http"
	services "github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/usecase/interface"
	response "github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/Response"
	"github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/models"
	"strconv"

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

	// fetching entered referall
	ref := c.Query("referall_code")

	// sending the struct and ref from client to UsecAse

	userCreated, err := u.userUseCase.UserSignup(user, ref)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "user couldnt signup", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	succesRes := response.ClientResponse(http.StatusCreated, "User successfully created", userCreated, nil)
	c.JSON(http.StatusCreated, succesRes)

}

func (u *UserHandler) UserLoginHandler(c *gin.Context) {

	var user models.UserLogin

	if err := c.BindJSON(&user); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provide are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	if err := validator.New().Struct(user); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "constraints not satisfied ", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	user_details, err := u.userUseCase.UserLoginHandler(user)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "User could not be logged in", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successResponse := response.ClientResponse(http.StatusOK, "User successfully logged in", user_details, nil)
	c.JSON(http.StatusOK, successResponse)

}

func (u *UserHandler) GetUserDetails(c *gin.Context) {

	id, _ := c.Get("id")

	userDetails, err := u.userUseCase.GetUserDetails(id.(int))
	if err != nil {
		errMsg := response.ClientResponse(http.StatusBadRequest, "Error Fetching Userdetails", nil, err.Error())
		c.JSON(http.StatusBadRequest, errMsg)
	}

	succesRes := response.ClientResponse(http.StatusOK, "successfully fetched Userdetails", userDetails, nil)
	c.JSON(http.StatusOK, succesRes)
}

func (u *UserHandler) EditUserDetails(c *gin.Context) {

	userID, exist := c.Get("id")

	id := userID.(int)

	if !exist {
		errRes := response.ClientResponse(http.StatusBadRequest, "empty userid", nil, nil)
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	var editUserdetails models.EditUserDetails

	err := c.BindJSON(&editUserdetails)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Error BInding Json ", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err = u.userUseCase.EditUserDetails(id, editUserdetails)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Error Editing Account Details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully edited Account Details", nil, nil)
	c.JSON(200, successRes)
}

func (u *UserHandler) AddAddressess(c *gin.Context) {

	idany, exist := c.Get("id")

	if !exist {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Empty userid", nil, nil)
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	id := idany.(int)

	var address models.AddAddress

	err := c.BindJSON(&address)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Invalid Json", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	err = u.userUseCase.AddAddress(id, address)
	if err != nil {
		errMsg := response.ClientResponse(http.StatusBadRequest, "Error Adding Address", nil, err.Error())
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added Address", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

func (u *UserHandler) GetAddressess(c *gin.Context) {

	id, exist := c.Get("id")
	if !exist {
		errRes := response.ClientResponse(http.StatusBadRequest, "empty userID", nil, nil)
		c.JSON(400, errRes)
		return
	}
	fmt.Println("jwtIDDD", id)
	uID := (id).(int)

	addresses, err := u.userUseCase.GetAddressess(uID)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Error fetching User addresses", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	succesRes := response.ClientResponse(http.StatusOK, "successfully fetched User addresses", addresses, nil)
	c.JSON(http.StatusOK, succesRes)
}

func (u *UserHandler) EditAddress(c *gin.Context) {
	//get address id
	idString := c.Query("addressID")

	id, err := strconv.Atoi(idString)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error converting param", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	// get userid
	uid, exists := c.Get("id")
	if !exists {
		errRes := response.ClientResponse(http.StatusBadRequest, "empty userID", nil, nil)
		c.JSON(400, errRes)
		return
	}
	userID := uid.(int)
	fmt.Println("addressID", id)
	fmt.Println("userJwtID", userID)

	var address models.EditAddress

	Err := c.BindJSON(&address)
	if Err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Error binding json model", nil, Err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err = u.userUseCase.EditAddress(id, uint(userID), address)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error editing Address", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully edited address", nil, nil)
	c.JSON(200, successRes)

}

func (u *UserHandler) DeleteAddress(c *gin.Context) {
	aID := c.Query("addressid")

	addressID, err := strconv.Atoi(aID)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error converting param", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	uID, exist := c.Get("id")
	if !exist {
		errRes := response.ClientResponse(http.StatusBadRequest, "empty userID", nil, nil)
		c.JSON(400, errRes)
		return
	}
	userID := uID.(int)

	err = u.userUseCase.DeleteAddress(addressID, userID)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Failed to delete address", nil, err.Error())
		c.JSON(400, errRes)
		return
	}
	succesMessage := fmt.Sprintf("successfully deleted address of id %d of UserID %d ", addressID, userID)
	successRes := response.ClientResponse(http.StatusOK, succesMessage, nil, nil)
	c.JSON(http.StatusOK, successRes)

}

func (u *UserHandler) ChangePassword(c *gin.Context) {

	id, _ := c.Get("id")
	userID := id.(int)

	var changePass models.ChangePassword

	err := c.BindJSON(&changePass)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Error binding json", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	err = u.userUseCase.ChangePassword(userID, changePass)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Error changing Password", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully changed Password", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

func (u *UserHandler) GetWallet(c *gin.Context) {

	user_id, _ := c.Get("id")

	wallet, err := u.userUseCase.GetWallet(user_id.(int))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "error fetching Wallet", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully Fetched Wallet", wallet, nil)
	c.JSON(http.StatusOK, successRes)
}
