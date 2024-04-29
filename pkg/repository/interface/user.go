package interfaces

import (
	"project/pkg/utils/domain"
	"project/pkg/utils/models"
)

type UserRepository interface {
	UserSignup(user models.UserDetails, referallID string) (models.UserDetailsResponse, error)
	CheckUserAvailability(email string) bool
	FindUserByEmail(user models.UserLogin) (models.UserSignInResponse, error)
	UserBlockStatus(email string) (bool, error)
	GetUserDetails(user_id int) (models.UserDetailsResponse, error)
	CheckifDefaultAddress(id int) (bool, error)
	AddAddress(id int, address models.AddAddress, defAddress bool) error
	GetAddressess(id int) ([]domain.Address, error)
	EditAddress(id int, userid uint, address models.EditAddress) error
	GetHashedPassword(id int) (string, error)
	EditUserDetails(id int, details models.EditUserDetails) error
	ChangePassword(id int, newHashedPass string) error
	DeleteAddress(addressID, userID int) error
	GetUserByReferralCode(refcode string) (int, error)
	CreateWallet(userID int) error
	AddMoneytoWallet(model models.AddMoneytoWallet) error
	GetWallet(userID int) (models.GetWallet, error)
	
}
