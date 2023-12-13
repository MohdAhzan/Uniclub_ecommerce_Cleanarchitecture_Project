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

var InternalError = "Internal Server Error"
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

	tokenString, _, err := u.helper.GenerateTokenClients(userdata)
	if err != nil {
		return models.TokenUsers{}, errors.New("could not create token due to some internal error")
	}

	return models.TokenUsers{
		Users: userdata,
		Token: tokenString,
	}, nil
}
