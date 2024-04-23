package models

type UserDetails struct {
	Name            string `json:"name"`
	Email           string `json:"email" validate:"email"`
	Phone           string `json:"phone"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmpassword"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password"`
}

// user details along with embedded token which can be used by the user to access protected routes
type TokenUsers struct {
	Users UserDetailsResponse
	Token string
}
type TokenUsersID struct {
	UserID uint `json:"id"`
}

// userdetails shown after logging in
type UserDetailsResponse struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	ReferralID string `json:"referral_id"`
}

type UserSignInResponse struct {
	Id         uint   `json:"id"`
	UserID     uint   `json:"user_id"`
	Name       string `json:"name"`
	Email      string `json:"email" validate:"email"`
	Phone      string `json:"phone"`
	Password   string `json:"password"`
	ReferralID string `json:"referal_id"`
}

type UserDetailsAtAdmin struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Blocked bool   `json:"blocked"`
}

type EditUserDetails struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
type AddAddress struct {
	Name     string `json:"name" validate:"required"`
	Address  string `json:"address"  validate:"required"`
	LandMark string `json:"landmark"`
	City     string `json:"city"  validate:"required"`
	Pincode  string `json:"pincode"   validate:"required,len=6"`
	State    string `json:"state" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
}

type EditAddress struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	LandMark string `json:"landmark"`
	City     string `json:"city"`
	Pincode  string `json:"pincode"`
	State    string `json:"state"`
	Phone    string `json:"phone"`
	Default  bool   `json:"default"`
}

type ChangePassword struct {
	CurrentPassword string `json:"currentPass"`
	NewPassword     string `json:"newPass"`
	ConfirmPassword string `json:"confirmPass"`
}

type Address struct {
	UserID   int
	Name     string
	LandMark string
	City     string
	Pincode  string
	State    string
	Phone    string
}

type CheckOut struct {
	CartID     int
	Addresses  Address
	Products   []GetCart
	TotalPrice float64
}

type AddMoneytoWallet struct {
	UserID          int
	Amount          float64
	TranscationType string
}

type GetWallet struct {
	UserID    int
	Username    string
	TotalAmount float64
}
