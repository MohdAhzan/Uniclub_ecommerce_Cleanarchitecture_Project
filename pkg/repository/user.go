package repository

import (
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
