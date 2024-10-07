package tests

import (
	"errors"
	"fmt"
	"log"
	"os"
	"project/pkg/config"
	"project/pkg/helper"
	"project/pkg/repository"
	"project/pkg/usecase"
	"project/pkg/utils/domain"
	"project/pkg/utils/models"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var cfg Config

type Config struct{
  TEST_DBHost     string `mapstructure:"TEST_DB_HOST"`
  TEST_DBName     string `mapstructure:"TEST_DB_NAME"`
  TEST_DBUser      string `mapstructure:"TEST_DB_USER"`
  TEST_DBPassword string `mapstructure:"TEST_DB_PASSWORD"` 
  TEST_DBPort string `mapstructure:"TEST_DB_PORT"` 
}

type Users struct {
  ID         uint   `json:"id" gorm:"unique;not null"`
  Name       string `json:"name"`
  Email      string `json:"email" validate:"email"`
  Password   string `json:"password" validate:"min=8,max=20"`
  Phone      string `json:"phone"`
  Blocked    bool   `json:"blocked" gorm:"default:false"`
  ReferralID string `json:"referral_id" gorm:"unique"`
}

func TestChangePassword(t *testing.T){

  testUsers:=[]Users{
    {

      ID: 1,
      Name: "ahzan",
      Email: "gsirkahzanpkm@gmail.com",
      Phone: "7034968543",
      Password: "user123",
      Blocked: false,
      ReferralID: "" ,

    },

  }

  db,dbErr:=testDBconnection(t)
  if dbErr!=nil{
    t.Errorf("error connecting the testDatabase \n %v",dbErr)
    return 
  } 


  defer func() {
    err:=truncateAllTables(db)
    if err!=nil{
      t.Error("error truncating the tables",err)
      return
    }


    DB,err:=db.DB()
    if err!=nil{
      t.Error(err)
      return
    }
    err= DB.Close()
    if err!=nil{
      t.Errorf("error closing testDB %v",err)
      return
    }
    t.Log("Closed testDB")
  }()


  h:=helper.NewHelper(config.Config{})
  userRepo:=repository.NewUserRepository(db) 
  userUsecase:=usecase.NewUserUseCase(userRepo,config.Config{},h)

  err:=db.AutoMigrate(&testUsers)
  if err!=nil{
    t.Errorf("error creating testUsers %v",err)
    return
  }

  for _,testUser:=range testUsers{

    hashedPass,err:= h.PasswordHashing(testUser.Password)
    if err!=nil{
      t.Errorf("error hashing testuser password %v",err)
      return
    }
    testUser.Password=hashedPass

    err= db.Exec(`INSERT into users (id,name,email,phone,password,blocked,referral_id) values(?,?,?,?,?,?,?)`,testUser.ID,testUser.Name,testUser.Email, testUser.Phone,testUser.Password,testUser.Blocked,testUser.ReferralID).Error
  }

  changePassModel:=[]models.ChangePassword{
    { CurrentPassword: "user123",NewPassword: "new123",ConfirmPassword: "new123"} ,
    { CurrentPassword: "incorrectPass",NewPassword: "foo",ConfirmPassword: "foo"} ,
    { CurrentPassword: "new123",NewPassword: "foo",ConfirmPassword: "bar"} ,
  }

  t.Run("success",func(t *testing.T) {

    err=userUsecase.ChangePassword(1,changePassModel[0])
    if err!=nil{
      t.Error("Error changing Pass",err)
      return
    }
    t.Log("Successfully changed Password")

  })

  t.Run("caseIncorrectPass",func(t *testing.T) {


    err=userUsecase.ChangePassword(1,changePassModel[1])
    if assert.Error(t,err){

      if assert.Equal(t,errors.New("incorrect PassWord !! Try Again"), err){

      }else{
        t.Error(err)
      }

    }


  })


  t.Run("casePassworddontMatch",func(t *testing.T) {

    err=userUsecase.ChangePassword(1,changePassModel[2])
    if assert.Error(t,err){

      if assert.Equal(t,errors.New("passwords doesn't match"), err){

      }else{
        t.Errorf("failed Incorrect Pass test expected %v got %v",errors.New("passwords doesn't match"),err)
        return
      }
    }
  })
}


var envs = []string{ "TEST_DB_HOST","TEST_DB_NAME","TEST_DB_USER","TEST_DB_PASSWORD","TEST_DB_PORT"}

var i int = 0 

//tesDBconnectoin func will loadup testdatabase for the whole integrated test
func testDBconnection(t *testing.T)(*gorm.DB ,error){


  if i==0{

    err := os.Chdir("../../../")
    if err != nil {
      log.Fatalf("Failed to set working directory: %v", err)
    }

    i++
  }else{
    err := os.Chdir(".")
    if err != nil {
      log.Fatalf("Failed to set working directory: %v", err)
    }
  }


  viper.AddConfigPath(".")
  viper.SetConfigFile(".env")

  if err := viper.ReadInConfig(); err != nil {
    return &gorm.DB{},err
  }



  for _, env := range envs {
    if err := viper.BindEnv(env); err != nil {
      return &gorm.DB{}, err
    }
  }


  if err := viper.Unmarshal(&cfg); err != nil {

    return &gorm.DB{}, err
  }
  if err := validator.New().Struct(&cfg); err != nil {
    return &gorm.DB{}, err
  }


  dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.TEST_DBHost, cfg.TEST_DBUser, cfg.TEST_DBName, cfg.TEST_DBPort, cfg.TEST_DBPassword)

  db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
  if dbErr !=nil{
    t.Errorf("error connecting the testdb_uniclub_project %s",dbErr)
    return db,dbErr 
  }

  return db,nil
}

func truncateAllTables(db *gorm.DB) error {
  err := db.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE").Error
  if err != nil {
    return err
  }
  return nil
}

func truncateWallet(db *gorm.DB) error {
  err := db.Exec("TRUNCATE TABLE wallets RESTART IDENTITY CASCADE").Error
  if err != nil {
    return err
  }
  return nil
}



// t.Log("jffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff\njjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj\njjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj")


func TestUserSignup(t *testing.T){

  db,dbErr:=testDBconnection(t)
  h:=helper.NewHelper(config.Config{})
  if dbErr!=nil{
    t.Error("error connecting Database",dbErr)
    return 
  }
  defer func() {
    err:=truncateAllTables(db)
    if err!=nil{
      t.Error("error truncating the tables",err)
      return
    }
    err=truncateWallet(db)
    if err!=nil{
      t.Error("error truncating the tables",err)
      return
    }


    DB,err:=db.DB()
    if err!=nil{
      t.Error(err)
      return
    }
    err= DB.Close()
    if err!=nil{
      t.Errorf("error closing testDB %v",err)
      return
    }
    t.Log("Closed testDB")
  }()



  userRepo:=repository.NewUserRepository(db) 
  userUsecase:=usecase.NewUserUseCase(userRepo,config.Config{},h)

  type arg struct{
    data models.UserDetails
    refCode string
  }

  mockData:=[]models.UserDetails{
    {

      Name: "ahzan",
      Email: "gsirkahzanpkm@gmail.com",
      Phone: "7034968543",
      Password:"user123",
      ConfirmPassword :"user123",
    },{
      Name: "Demo2",
      Email: "demo@gmail.com",
      Phone: "9999999999",
      Password:"foo",
      ConfirmPassword :"barrr"},

  }

  err:=db.AutoMigrate(&domain.Users{})
  if err!=nil{
    t.Errorf("error creating testUsers %v",err)
    return
  }
  err=db.AutoMigrate(&domain.Wallet{})
  if err!=nil{
    t.Errorf("error creating testUsers %v",err)
    return
  }


  testCases:=map[string]struct{
    input  arg 
    expect models.TokenUsers 
    expectErr error
  }{

    "successfullSignup":{
      input: arg{data: mockData[0],refCode: ""} ,

      expect: models.TokenUsers{Users:models.UserDetailsResponse{Id:1,Name: "ahzan",Email: "gsirkahzanpkm@gmail.com", Phone: "7034968543",ReferralID: ""},
      },
      expectErr: nil,
    },
    "failedPassIncorrect":{
      input: arg{data: mockData[1],refCode: ""} ,

      expect: models.TokenUsers{Users:models.UserDetailsResponse{Id:2,Name: "ahzan",Email: "gsirkahzanpkm@gmail.com", Phone: "7034968543",ReferralID: ""},
      },
      expectErr: errors.New("password does not match"),

    },

  }

  for key,testCase:= range testCases {


    t.Run(key,func(t *testing.T) {


      response,err:=userUsecase.UserSignup(testCase.input.data,testCase.input.refCode)

      // assert.Error(t,testCase.expectErr,err)

      if err!=nil{
        if assert.Error(t, err) {
          assert.Equal(t, testCase.expectErr, err)
          return
        }
      }
      //  if err!=nil{
      //   assert.Error()("error in the case %s \n %v  ",key,err)
      //
      // } 




      // assert.Equal(t,testCase.expect.AccessToken,response.AccessToken)
      // assert.Equal(t,testCase.expect.RefreshToken,response.RefreshToken)
      // assert.Equal(t,testCase.expect.Users.Id,response.Users.Id)
      assert.Equal(t,testCase.expect.Users.Email,response.Users.Email)
      assert.Equal(t,testCase.expect.Users.Phone,response.Users.Phone)

      //since refcode can be random its testing to be two random values
      if !assert.NotEqual(t,testCase.expect.Users.ReferralID,response.Users.ReferralID){

        t.Errorf("%v",errors.New("expected two different referall codes but got same"))

      }

    })

  }



}


// accessToken,refreshToken,err :=  mockhelper.GenerateTokenClients(mockData[1])
// if err!=nil{
//   t.Errorf(err.Error())
// }
//
// t.Run("TestONE",func(t *testing.T) {
//
//   t.Log(accessToken)
//   t.Log(refreshToken)
//
// })
//
// }

