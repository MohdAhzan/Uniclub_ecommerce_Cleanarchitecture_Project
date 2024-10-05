[1mdiff --git a/pkg/helper/helper.go b/pkg/helper/helper.go[m
[1mindex 209b562..66191ec 100644[m
[1m--- a/pkg/helper/helper.go[m
[1m+++ b/pkg/helper/helper.go[m
[36m@@ -79,7 +79,7 @@[m [mfunc (h *helper) GenerateTokenAdmin(admin models.AdminDetailsResponse) (string,[m
 	// 		IssuedAt:  jwt.NewNumericDate(time.Now()),[m
 	// 	},[m
 	// }[m
[31m-[m
[32m+[m	[32m//[m
 	// refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)[m
 	// refreshTokenString, err := refreshToken.SignedString([]byte("adminrefreshToken988243rwcfsdsjfyf74cysf38"))[m
 	// if err != nil {[m
[36m@@ -90,40 +90,40 @@[m [mfunc (h *helper) GenerateTokenAdmin(admin models.AdminDetailsResponse) (string,[m
 [m
 }[m
 [m
[31m-func (h *helper) GenerateTokenClients(user models.UserDetailsResponse) (string, error) {[m
[32m+[m[32mfunc (h *helper) GenerateTokenClients(user models.UserDetailsResponse) (string,string, error) {[m
 	accessTokenClaims := &AuthCustomClaims{[m
 		Id:    user.Id,[m
 		Email: user.Email,[m
 		Role:  "client",[m
 		RegisteredClaims: jwt.RegisteredClaims{[m
[31m-			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),[m
[32m+[m			[32mExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 1)),[m
 			IssuedAt:  jwt.NewNumericDate(time.Now()),[m
 		},[m
 	}[m
 [m
[31m-	// refreshTokenClaims := &AuthCustomClaims{[m
[31m-	// 	Id:    user.Id,[m
[31m-	// 	Email: user.Email,[m
[31m-	// 	Role:  "client",[m
[31m-	// 	RegisteredClaims: jwt.RegisteredClaims{[m
[31m-	// 		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),[m
[31m-	// 		IssuedAt:  jwt.NewNumericDate(time.Now()),[m
[31m-	// 	},[m
[31m-	// }[m
[32m+[m	[32mrefreshTokenClaims := &AuthCustomClaims{[m
[32m+[m		[32mId:    user.Id,[m
[32m+[m		[32mEmail: user.Email,[m
[32m+[m		[32mRole:  "client",[m
[32m+[m		[32mRegisteredClaims: jwt.RegisteredClaims{[m
[32m+[m			[32mExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),[m
[32m+[m			[32mIssuedAt:  jwt.NewNumericDate(time.Now()),[m
[32m+[m		[32m},[m
[32m+[m	[32m}[m
 [m
 	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)[m
 	accessTokenString, err := accessToken.SignedString([]byte("useraccesstokenasdioufou23854284jsdf9823jsdfh"))[m
 	if err != nil {[m
[31m-		return "", err[m
[32m+[m		[32mreturn "","", err[m
 	}[m
 [m
[31m-	// refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)[m
[31m-	// refreshTokenString, err := refreshToken.SignedString([]byte("userrefreshtokenasdgfr23788h23cy86qnw3dr367d4ye2"))[m
[31m-	// if err != nil {[m
[31m-	// 	return "", "", err[m
[31m-	// }[m
[32m+[m	[32mrefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)[m
[32m+[m	[32mrefreshTokenString, err := refreshToken.SignedString([]byte("userrefreshtokenasdgfr23788h23cy86qnw3dr367d4ye2"))[m
[32m+[m	[32mif err != nil {[m
[32m+[m		[32mreturn "", "", err[m
[32m+[m	[32m}[m
 [m
[31m-	return accessTokenString, nil[m
[32m+[m	[32mreturn accessTokenString,refreshTokenString ,nil[m
 [m
 }[m
 [m
[1mdiff --git a/pkg/helper/interface/helper.go b/pkg/helper/interface/helper.go[m
[1mindex 403f306..1bd03bd 100644[m
[1m--- a/pkg/helper/interface/helper.go[m
[1m+++ b/pkg/helper/interface/helper.go[m
[36m@@ -10,7 +10,7 @@[m [mimport ([m
 [m
 type Helper interface {[m
 	PasswordHashing(string) (string, error)[m
[31m-	GenerateTokenClients(user models.UserDetailsResponse) (string, error)[m
[32m+[m	[32mGenerateTokenClients(user models.UserDetailsResponse) (string,string, error)[m
 	GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, error)[m
 	CompareHashAndPassword(hashedPassword string, password string) error[m
 [m
[1mdiff --git a/pkg/mocks/admin/mock_admin_helper.go b/pkg/mocks/admin/mock_admin_helper.go[m
[1mindex dd22c5e..4cb732d 100644[m
[1m--- a/pkg/mocks/admin/mock_admin_helper.go[m
[1m+++ b/pkg/mocks/admin/mock_admin_helper.go[m
[36m@@ -112,12 +112,13 @@[m [mfunc (mr *MockHelperMockRecorder) GenerateTokenAdmin(admin interface{}) *gomock.[m
 }[m
 [m
 // GenerateTokenClients mocks base method.[m
[31m-func (m *MockHelper) GenerateTokenClients(user models.UserDetailsResponse) (string, error) {[m
[32m+[m[32mfunc (m *MockHelper) GenerateTokenClients(user models.UserDetailsResponse) (string, string, error) {[m
 	m.ctrl.T.Helper()[m
 	ret := m.ctrl.Call(m, "GenerateTokenClients", user)[m
 	ret0, _ := ret[0].(string)[m
[31m-	ret1, _ := ret[1].(error)[m
[31m-	return ret0, ret1[m
[32m+[m	[32mret1, _ := ret[1].(string)[m
[32m+[m	[32mret2, _ := ret[2].(error)[m
[32m+[m	[32mreturn ret0, ret1, ret2[m
 }[m
 [m
 // GenerateTokenClients indicates an expected call of GenerateTokenClients.[m
[1mdiff --git a/pkg/usecase/admin_test.go b/pkg/usecase/admin_test.go[m
[1mindex e494626..4f6ee00 100644[m
[1m--- a/pkg/usecase/admin_test.go[m
[1m+++ b/pkg/usecase/admin_test.go[m
[36m@@ -7,7 +7,6 @@[m [mimport ([m
 	"project/pkg/utils/domain"[m
 	"project/pkg/utils/models"[m
 	"testing"[m
[31m-	"time"[m
 [m
 	"project/pkg/usecase"[m
 [m
[1mdiff --git a/pkg/usecase/otp.go b/pkg/usecase/otp.go[m
[1mindex cb5fad6..37be2b2 100644[m
[1m--- a/pkg/usecase/otp.go[m
[1m+++ b/pkg/usecase/otp.go[m
[36m@@ -59,7 +59,8 @@[m [mfunc (otp *otpUseCase) VerifyOTP(code models.VerifyData) (models.TokenUsers, err[m
 		return models.TokenUsers{}, err[m
 	}[m
 [m
[31m-	tokenString, err := otp.helper.GenerateTokenClients(userDetails)[m
[32m+[m	[32mtokenString,refreshtokenString, err := otp.helper.GenerateTokenClients(userDetails)[m
[32m+[m
 	if err != nil {[m
 		return models.TokenUsers{}, err[m
 	}[m
[36m@@ -73,7 +74,8 @@[m [mfunc (otp *otpUseCase) VerifyOTP(code models.VerifyData) (models.TokenUsers, err[m
 [m
 	return models.TokenUsers{[m
 		Users: user,[m
[31m-		Token: tokenString,[m
[32m+[m		[32mAccessToken: tokenString,[m
[32m+[m		[32mRefreshToken: refreshtokenString,[m
 	}, nil[m
 [m
 }[m
[1mdiff --git a/pkg/usecase/user.go b/pkg/usecase/user.go[m
[1mindex b26d0ab..433b211 100644[m
[1m--- a/pkg/usecase/user.go[m
[1m+++ b/pkg/usecase/user.go[m
[36m@@ -106,14 +106,15 @@[m [mfunc (u *userUseCase) UserSignup(user models.UserDetails, refCode string) (model[m
 [m
 	// creating a jwt token for clients[m
 [m
[31m-	tokenString, err := u.helper.GenerateTokenClients(userdata)[m
[32m+[m	[32mtokenString,refreshString, err := u.helper.GenerateTokenClients(userdata)[m
 	if err != nil {[m
 		return models.TokenUsers{}, errors.New("could not create token due to some internal error")[m
 	}[m
 [m
 	return models.TokenUsers{[m
 		Users: userdata,[m
[31m-		Token: tokenString,[m
[32m+[m[32m    AccessToken :tokenString ,[m
[32m+[m[32m    RefreshToken :refreshString ,[m
 	}, nil[m
 }[m
 [m
[36m@@ -154,14 +155,15 @@[m [mfunc (u *userUseCase) UserLoginHandler(user models.UserLogin) (models.TokenUsers[m
 	userDetails.Phone = user_details.Phone[m
 	userDetails.ReferralID = user_details.ReferralID[m
 [m
[31m-	tokenString, err := u.helper.GenerateTokenClients(userDetails)[m
[32m+[m	[32mtokenString,refreshString, err := u.helper.GenerateTokenClients(userDetails)[m
 	if err != nil {[m
 		return models.TokenUsers{}, errors.New("couldn't generate token for client ")[m
 	}[m
 [m
 	return models.TokenUsers{[m
 		Users: userDetails,[m
[31m-		Token: tokenString,[m
[32m+[m		[32mAccessToken:tokenString,[m
[32m+[m[32m    RefreshToken: refreshString,[m
 	}, nil[m
 }[m
 [m
[1mdiff --git a/pkg/utils/models/user.go b/pkg/utils/models/user.go[m
[1mindex b1b1a3a..c65570c 100644[m
[1m--- a/pkg/utils/models/user.go[m
[1m+++ b/pkg/utils/models/user.go[m
[36m@@ -1,113 +1,114 @@[m
 package models[m
 [m
 type UserDetails struct {[m
[31m-	Name            string `json:"name"`[m
[31m-	Email           string `json:"email" validate:"email"`[m
[31m-	Phone           string `json:"phone"`[m
[31m-	Password        string `json:"password"`[m
[31m-	ConfirmPassword string `json:"confirmpassword"`[m
[32m+[m[32m  Name            string `json:"name"`[m
[32m+[m[32m  Email           string `json:"email" validate:"email"`[m
[32m+[m[32m  Phone           string `json:"phone"`[m
[32m+[m[32m  Password        string `json:"password"`[m
[32m+[m[32m  ConfirmPassword string `json:"confirmpassword"`[m
 }[m
 [m
 type UserLogin struct {[m
[31m-	Email    string `json:"email" validate:"email"`[m
[31m-	Password string `json:"password"`[m
[32m+[m[32m  Email    string `json:"email" validate:"email"`[m
[32m+[m[32m  Password string `json:"password"`[m
 }[m
 [m
 // user details along with embedded token which can be used by the user to access protected routes[m
 type TokenUsers struct {[m
[31m-	Users UserDetailsResponse[m
[31m-	Token string[m
[32m+[m[32m  Users UserDetailsResponse[m
[32m+[m[32m  AccessToken string[m
[32m+[m[32m  RefreshToken string[m
 }[m
 type TokenUsersID struct {[m
[31m-	UserID uint `json:"id"`[m
[32m+[m[32m  UserID uint `json:"id"`[m
 }[m
 [m
 // userdetails shown after logging in[m
 type UserDetailsResponse struct {[m
[31m-	Id         int    `json:"id"`[m
[31m-	Name       string `json:"name"`[m
[31m-	Email      string `json:"email"`[m
[31m-	Phone      string `json:"phone"`[m
[31m-	ReferralID string `json:"referral_id"`[m
[32m+[m[32m  Id         int    `json:"id"`[m
[32m+[m[32m  Name       string `json:"name"`[m
[32m+[m[32m  Email      string `json:"email"`[m
[32m+[m[32m  Phone      string `json:"phone"`[m
[32m+[m[32m  ReferralID string `json:"referral_id"`[m
 }[m
 [m
 type UserSignInResponse struct {[m
[31m-	Id         uint   `json:"id"`[m
[31m-	UserID     uint   `json:"user_id"`[m
[31m-	Name       string `json:"name"`[m
[31m-	Email      string `json:"email" validate:"email"`[m
[31m-	Phone      string `json:"phone"`[m
[31m-	Password   string `json:"password"`[m
[31m-	ReferralID string `json:"referal_id"`[m
[32m+[m[32m  Id         uint   `json:"id"`[m
[32m+[m[32m  UserID     uint   `json:"user_id"`[m
[32m+[m[32m  Name       string `json:"name"`[m
[32m+[m[32m  Email      string `json:"email" validate:"email"`[m
[32m+[m[32m  Phone      string `json:"phone"`[m
[32m+[m[32m  Password   string `json:"password"`[m
[32m+[m[32m  ReferralID string `json:"referal_id"`[m
 }[m
 [m
 type UserDetailsAtAdmin struct {[m
[31m-	Id      uint   `json:"id"`[m
[31m-	Name    string `json:"name"`[m
[31m-	Email   string `json:"email"`[m
[31m-	Phone   string `json:"phone"`[m
[31m-	Blocked bool   `json:"blocked"`[m
[32m+[m[32m  Id      uint   `json:"id"`[m
[32m+[m[32m  Name    string `json:"name"`[m
[32m+[m[32m  Email   string `json:"email"`[m
[32m+[m[32m  Phone   string `json:"phone"`[m
[32m+[m[32m  Blocked bool   `json:"blocked"`[m
 }[m
 [m
 type EditUserDetails struct {[m
[31m-	Name     string `json:"name"`[m
[31m-	Email    string `json:"email"`[m
[31m-	Phone    string `json:"phone"`[m
[31m-	Password string `json:"password"`[m
[32m+[m[32m  Name     string `json:"name"`[m
[32m+[m[32m  Email    string `json:"email"`[m
[32m+[m[32m  Phone    string `json:"phone"`[m
[32m+[m[32m  Password string `json:"password"`[m
 }[m
 type AddAddress struct {[m
[31m-	Name     string `json:"name" validate:"required"`[m
[31m-	Address  string `json:"address"  validate:"required"`[m
[31m-	LandMark string `json:"landmark"`[m
[31m-	City     string `json:"city"  validate:"required"`[m
[31m-	Pincode  string `json:"pincode"   validate:"required,len=6"`[m
[31m-	State    string `json:"state" validate:"required"`[m
[31m-	Phone    string `json:"phone" validate:"required"`[m
[32m+[m[32m  Name     string `json:"name" validate:"required"`[m
[32m+[m[32m  Address  string `json:"address"  validate:"required"`[m
[32m+[m[32m  LandMark string `json:"landmark"`[m
[32m+[m[32m  City     string `json:"city"  validate:"required"`[m
[32m+[m[32m  Pincode  string `json:"pincode"   validate:"required,len=6"`[m
[32m+[m[32m  State    string `json:"state" validate:"required"`[m
[32m+[m[32m  Phone    string `json:"phone" validate:"required"`[m
 }[m
 [m
 type EditAddress struct {[m
[31m-	Name     string `json:"name"`[m
[31m-	Address  string `json:"address"`[m
[31m-	LandMark string `json:"landmark"`[m
[31m-	City     string `json:"city"`[m
[31m-	Pincode  string `json:"pincode"`[m
[31m-	State    string `json:"state"`[m
[31m-	Phone    string `json:"phone"`[m
[31m-	Default  bool   `json:"default"`[m
[32m+[m[32m  Name     string `json:"name"`[m
[32m+[m[32m  Address  string `json:"address"`[m
[32m+[m[32m  LandMark string `json:"landmark"`[m
[32m+[m[32m  City     string `json:"city"`[m
[32m+[m[32m  Pincode  string `json:"pincode"`[m
[32m+[m[32m  State    string `json:"state"`[m
[32m+[m[32m  Phone    string `json:"phone"`[m
[32m+[m[32m  Default  bool   `json:"default"`[m
 }[m
 [m
 type ChangePassword struct {[m
[31m-	CurrentPassword string `json:"currentPass"`[m
[31m-	NewPassword     string `json:"newPass"`[m
[31m-	ConfirmPassword string `json:"confirmPass"`[m
[32m+[m[32m  CurrentPassword string `json:"currentPass"`[m
[32m+[m[32m  NewPassword     string `json:"newPass"`[m
[32m+[m[32m  ConfirmPassword string `json:"confirmPass"`[m
 }[m
 [m
 type Address struct {[m
[31m-	UserID   int[m
[31m-	Name     string[m
[31m-	LandMark string[m
[31m-	City     string[m
[31m-	Pincode  string[m
[31m-	State    string[m
[31m-	Phone    string[m
[32m+[m[32m  UserID   int[m
[32m+[m[32m  Name     string[m
[32m+[m[32m  LandMark string[m
[32m+[m[32m  City     string[m
[32m+[m[32m  Pincode  string[m
[32m+[m[32m  State    string[m
[32m+[m[32m  Phone    string[m
 }[m
 [m
 type CheckOut struct {[m
[31m-	CartID     int[m
[31m-	Addresses  Address[m
[31m-	Products   []GetCart[m
[31m-	TotalPrice float64[m
[31m-	Final_Price float64[m
[32m+[m[32m  CartID     int[m
[32m+[m[32m  Addresses  Address[m
[32m+[m[32m  Products   []GetCart[m
[32m+[m[32m  TotalPrice float64[m
[32m+[m[32m  Final_Price float64[m
 }[m
 [m
 type AddMoneytoWallet struct {[m
[31m-	UserID          int[m
[31m-	Amount          float64[m
[31m-	TranscationType string[m
[32m+[m[32m  UserID          int[m
[32m+[m[32m  Amount          float64[m
[32m+[m[32m  TranscationType string[m
 }[m
 [m
 type GetWallet struct {[m
[31m-	UserID    int[m
[31m-	Username    string[m
[31m-	TotalAmount float64[m
[32m+[m[32m  UserID    int[m
[32m+[m[32m  Username    string[m
[32m+[m[32m  TotalAmount float64[m
 }[m
