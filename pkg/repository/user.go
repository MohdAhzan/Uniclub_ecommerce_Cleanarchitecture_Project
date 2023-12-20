package repository

import (
	"errors"
	"fmt"
	interfaces "project/pkg/repository/interface"
	"project/pkg/utils/models"

	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}

func (c *userDatabase) CheckUserAvailability(email string) bool {
	var count int
	query := fmt.Sprintf("select count(*) from users where email='%s'", email)
	if err := c.DB.Raw(query).Scan(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (u *userDatabase) UserSignup(user models.UserDetails) (models.UserDetailsResponse, error) {
	var userDetails models.UserDetailsResponse
	err := u.DB.Raw("insert into users (name,email,password,phone) values (?,?,?,?) RETURNING id,name,email,phone", user.Name, user.Email, user.Password, user.Phone).Scan(&userDetails).Error
	if err != nil {
		return models.UserDetailsResponse{}, err
	}
	fmt.Println("USER DETAILS eNTEREDD")
	return userDetails, nil
}

func (u *userDatabase) FindUserByEmail(user models.UserLogin) (models.UserSignInResponse, error) {
	var user_details models.UserSignInResponse

	err := u.DB.Raw("select * from users where email = ? and blocked = false", user.Email).Scan(&user_details).Error

	if err != nil {
		return models.UserSignInResponse{}, errors.New("error getting user details")

	}
	return user_details, nil
}

func (u *userDatabase) UserBlockStatus(email string) (bool, error) {
	var isBlocked bool
	err := u.DB.Raw("select blocked from users where email = ?", email).Scan(&isBlocked).Error
	if err != nil {
		return false, err
	}
	return isBlocked, nil
}
