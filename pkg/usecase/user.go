package usecase

import (
	"errors"
	"fmt"
	config "project/pkg/config"
	helper_interface "project/pkg/helper/interface"
	interfaces "project/pkg/repository/interface"
	"project/pkg/utils/domain"
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

func (u *userUseCase) UserSignup(user models.UserDetails, refCode string) (models.TokenUsers, error) {

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

	//create an referral ID for user
	referalID, err := u.helper.GenerateReferralCode()
	if err != nil {
		return models.TokenUsers{}, err
	}

	// INSERT USER DETAILS HIS REFERRAL ID TO DATABASE
	userdata, err := u.userRepo.UserSignup(user, referalID)
	if err != nil {
		return models.TokenUsers{}, errors.New("failed to add the user")
	}

	//CREATE A WALLET FOR NEW USER
	err = u.userRepo.CreateWallet(userdata.Id)
	if err != nil {
		return models.TokenUsers{}, err
	}

	//CHECKS IF NEW USER HAS ANY REFERAL CODES
	if len(refCode) != 0 {
		//GET REFERRED USER DETAILS FOR REFERAL CASHBACK
		userID, err := u.userRepo.GetUserByReferralCode(refCode)
		if err != nil {
			return models.TokenUsers{}, errors.New("invalid referall code please enter and valid referal code to earn cashback")
		}
		//get reffered users detail for checking and crediting amount
		refferedUser, err := u.userRepo.GetUserDetails(userID)
		if err != nil {
			return models.TokenUsers{}, err
		}
		
    exists := u.userRepo.CheckUserAvailability(refferedUser.Email)
		if exists {
			// credit 100rs to referred User
			var model models.AddMoneytoWallet
			model.UserID = userID
			model.Amount = 100 //100rs  for referred user
			model.TranscationType = "REFERAL"

			err = u.userRepo.AddMoneytoWallet(model)
			if err != nil {
				return models.TokenUsers{}, err
			}
			//and 50 to new signed User
			model.UserID = userdata.Id
			model.Amount = 50 //50rs for referred user
			model.TranscationType = "REFERAL"
			err = u.userRepo.AddMoneytoWallet(model)
			if err != nil {
				return models.TokenUsers{}, err
			}

		}
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
	userDetails.ReferralID = user_details.ReferralID

	tokenString, err := u.helper.GenerateTokenClients(userDetails)
	if err != nil {
		return models.TokenUsers{}, errors.New("couldn't generate token for client ")
	}

	return models.TokenUsers{
		Users: userDetails,
		Token: tokenString,
	}, nil
}

func (u userUseCase) GetUserDetails(id int) (models.UserDetailsResponse, error) {

	userDetails, err := u.userRepo.GetUserDetails(id)
	if err != nil {
		return models.UserDetailsResponse{}, err
	}

	return userDetails, nil
}

func (u *userUseCase) EditUserDetails(id int, details models.EditUserDetails) error {

	// exist := u.userRepo.CheckUserAvailability(details.Email)
	// if !exist {
	// 	return errors.New("INvalid userid check user")
	// }

	hashedPassword, err := u.userRepo.GetHashedPassword(id)
	if err != nil {
		return errors.New("error fetching EncryptedPassword")
	}
	Err := u.helper.CompareHashAndPassword(hashedPassword, details.Password)

	if Err != nil {
		return errors.New("incorrect PassWord !! Try Again")

	}

	err = u.userRepo.EditUserDetails(id, details)
	if err != nil {
		return err
	}

	return nil
}

func (u userUseCase) AddAddress(id int, address models.AddAddress) error {

	isDefault, err := u.userRepo.CheckifDefaultAddress(id)
	if err != nil {
		return err
	}

	err = u.userRepo.AddAddress(id, address, isDefault)
	if err != nil {
		return err
	}

	return nil

}

func (u userUseCase) GetAddressess(id int) ([]domain.Address, error) {

	notexist, err := u.userRepo.CheckifDefaultAddress(id)
	if err != nil {
		return []domain.Address{}, err
	}
	if notexist {
		return []domain.Address{}, errors.New("no address for this User")
	}

	addressess, err := u.userRepo.GetAddressess(id)
	if err != nil {
		return []domain.Address{}, err
	}

	return addressess, nil
}

func (u userUseCase) EditAddress(id int, userId uint, address models.EditAddress) error {

	err := u.userRepo.EditAddress(id, userId, address)
	if err != nil {
		return err
	}

	return nil
}

func (u userUseCase) DeleteAddress(addressID int, userID int) error {
	err := u.userRepo.DeleteAddress(addressID, userID)
	if err != nil {
		return err
	}
	return nil
}

func (u userUseCase) ChangePassword(id int, changePass models.ChangePassword) error {

	hashedPassword, err := u.userRepo.GetHashedPassword(id)
	if err != nil {
		return errors.New("error fetching EncryptedPassword")
	}
	Err := u.helper.CompareHashAndPassword(hashedPassword, changePass.CurrentPassword)

	if Err != nil {
		return errors.New("incorrect PassWord !! Try Again")

	}

	if changePass.NewPassword != changePass.ConfirmPassword {
		return errors.New("passwords doesn't match")
	}

	newHashedPass, err := u.helper.PasswordHashing(changePass.NewPassword)
	if err != nil {
		return errors.New("failed to hash newPassword")
	}

	err = u.userRepo.ChangePassword(id, newHashedPass)
	if err != nil {
		return err
	}
	return nil
}

func (u userUseCase) GetWallet(userID int) (models.GetWallet, error) {

	userdata, err := u.userRepo.GetUserDetails(userID)
	if err != nil {
		return models.GetWallet{}, err
	}

	walletData, err := u.userRepo.GetWallet(userID)
	if err != nil {
		return models.GetWallet{}, err
	}
	walletData.UserID = userID
	walletData.Username = userdata.Name
	return walletData, nil
}
