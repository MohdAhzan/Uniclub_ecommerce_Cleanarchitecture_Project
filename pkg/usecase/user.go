package usecase

import (
	"errors"
	"fmt"
	config "project/pkg/config"
	helper_interface "project/pkg/helper/interface"
	interfaces "project/pkg/repository/interface"
	"project/pkg/utils/models"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
	cfg      config.Config
	helper   helper_interface.Helper
}

func NewUserUseCase(repo interfaces.UserRepository, cfg config.Config, h helper_interface.Helper) *userUseCase {

	return &userUseCase{
		userRepo: repo,
		cfg:      cfg,
		helper:   h,
	}
}

var ErrorHashingPassword = "Error In Hashing Password"

func (u *userUseCase) UserSignup(user models.UserDetails) (models.TokenUsers, error) {

	fmt.Println("<<<Add Users>>>")
	//check if user already exists
	userExist := u.userRepo.CheckUserAvailability(user.Email)
	if userExist {
		return models.TokenUsers{}, errors.New("user already exist, please Signin")
	}

	if user.Password != user.ConfirmPassword {
		return models.TokenUsers{}, errors.New("password does not match")
	}

	// password Hashing

	hashedPassword, err := u.helper.PasswordHashing(user.Password)
	if err != nil {
		return models.TokenUsers{}, errors.New(ErrorHashingPassword)
	}

	user.Password = hashedPassword

	// INSERT USER DETAILS TO DATABASE

	userdata, err := u.userRepo.UserSignup(user)
	if err != nil {
		return models.TokenUsers{}, errors.New("couldn't add the user")
	}

	// creating a jwt token for clients

	tokenString, err := u.helper.GenerateTokenClients(userdata)
	if err != nil {
		return models.TokenUsers{}, errors.New("could not create token due to some internal error")
	}

	return models.TokenUsers{
		Users: userdata,
		Token: tokenString,
	}, nil
}

func (u *userUseCase) UserLoginHandler(user models.UserLogin) (models.TokenUsers, error) {

	//check if user exist in this email
	ok := u.userRepo.CheckUserAvailability(user.Email)
	if !ok {
		return models.TokenUsers{}, errors.New("the user does not exist")
	}

	//check if user is blocked or banned
	isBlocked, err := u.userRepo.UserBlockStatus(user.Email)

	if err != nil {
		return models.TokenUsers{}, errors.New("error checking the userblockstatus")
	}
	if isBlocked {
		return models.TokenUsers{}, errors.New("the user is blocked by admin")
	}

	//fetching userdetails to check password
	user_details, err := u.userRepo.FindUserByEmail(user)
	if err != nil {
		return models.TokenUsers{}, errors.New("error fetching userdetails")
	}

	err = u.helper.CompareHashAndPassword(user_details.Password, user.Password)
	if err != nil {
		return models.TokenUsers{}, errors.New("incorrect password")
	}

	var userDetails models.UserDetailsResponse

	userDetails.Id = int(user_details.Id)
	userDetails.Name = user_details.Name
	userDetails.Email = user_details.Email
	userDetails.Phone = user_details.Phone

	tokenString, err := u.helper.GenerateTokenClients(userDetails)
	if err != nil {
		return models.TokenUsers{}, errors.New("couldn't generate token for client ")
	}

	return models.TokenUsers{
		Users: userDetails,
		Token: tokenString,
	}, nil
}
