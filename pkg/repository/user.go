package repository

import (
	"errors"
	"fmt"
	interfaces "project/pkg/repository/interface"
	"project/pkg/utils/domain"
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

func (u *userDatabase) GetUserDetails(user_id int) (models.UserDetailsResponse, error) {

	var userDetails models.UserDetailsResponse

	err := u.DB.Raw("select id,name,email,phone from users where id = ? ", user_id).Scan(&userDetails).Error
	if err != nil {
		return models.UserDetailsResponse{}, err
	}

	return userDetails, err
}

func (u *userDatabase) EditUserDetails(id int, details models.EditUserDetails) error {

	fmt.Println(details)

	err := u.DB.Exec("UPDATE users SET name= $1,email = $2,phone=$3 where id = $4", details.Name, details.Email, details.Phone, id).Error
	if err != nil {
		return errors.New("error updating User accountDetails")
	}

	return nil
}

func (u *userDatabase) CheckifDefaultAddress(userID int) (bool, error) {
	var count int
	if err := u.DB.Raw("select count(*) from addresses where user_id = ?", userID).Scan(&count).Error; err != nil {
		return false, err
	}
	if count > 0 {
		return false, nil
	}

	return true, nil
}

func (u *userDatabase) AddAddress(userID int, address models.AddAddress, defAddress bool) error {

	err := u.DB.Exec(`INSERT INTO addresses (user_id, name, address,land_mark,  city,pincode ,state, phone, "default")
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8 ,$9)`, userID, address.Name, address.Address, address.LandMark,
		address.City, address.Pincode, address.State, address.Phone, defAddress).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *userDatabase) GetAddressess(id int) ([]domain.Address, error) {

	var addressess []domain.Address
	err := u.DB.Raw("select * from addresses where user_id=?", id).Scan(&addressess).Error
	if err != nil {
		return []domain.Address{}, errors.New("no addressess in user")
	}
	return addressess, nil

}

func (u *userDatabase) EditAddress(id int, userid uint, address models.EditAddress) error {

	var count int

	err := u.DB.Raw("select count(*) from addresses where id= $1 and user_id = $2", id, userid).Scan(&count).Error
	if err != nil {
		return err
	}

	if count < 1 {
		return errors.New("enter valid address_id")
	}

	// if we want to make address default change rest to not default
	if address.Default {
		u.DB.Exec(`UPDATE addresses SET "default" = false where user_id = ? `, userid)
	}

	err = u.DB.Exec(`UPDATE addresses SET name = $1, address = $2, land_mark = $3, city = $4, pincode = $5, state = $6,phone=$7, "default" = $8 where user_id =$9 and id = $10`,
		address.Name, address.Address, address.LandMark, address.City, address.Pincode, address.State, address.Phone, address.Default, userid, id).Error

	if err != nil {
		return err
	}
	return nil
}

func (u *userDatabase) GetHashedPassword(id int) (string, error) {

	var hashedPassword string

	err := u.DB.Raw("select password from users where id = ?", id).Scan(&hashedPassword).Error
	if err != nil {
		return "", errors.New("failed to get hashedPass")
	}

	return hashedPassword, nil
}

func (u *userDatabase) ChangePassword(id int, newHashedPass string) error {

	err := u.DB.Raw("UPDATE users SET password = $1 where id = $2", newHashedPass, id).Error
	if err != nil {
		return err
	}


	return nil
}
