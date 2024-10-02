package usecase_test

import (
	"errors"
	"fmt"
	mocks "project/pkg/mocks/admin"
	"project/pkg/utils/domain"
	"project/pkg/utils/models"
	"testing"
	"time"

	"project/pkg/usecase"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {

  mockctrl := gomock.NewController(t)
  mockadminRepo := mocks.NewMockAdminRepository(mockctrl)
  mockhelper := mocks.NewMockHelper(mockctrl)
  mockorderRepo := mocks.NewMockOrderRepository(mockctrl)
  mockuserRepo := mocks.NewMockUserRepository(mockctrl)

  mockadminUsecase := usecase.NewAdminUsecase(mockadminRepo, mockhelper, mockorderRepo, mockuserRepo)

  // Define what the mock should return when GetUsers is called
  t.Parallel() 

  mockUsers := []models.UserDetailsAtAdmin{
    {Id: 1, Name: "Demouser1", Email: "demouser1@gmail.com", Phone: "9878362828", Blocked: false},
    {Id: 2, Name: "Demouser2", Email: "demouser2@gmail.com", Phone: "9999999999", Blocked: true},
  }

  testCases := map[string]struct {
    stub      func(mocks.MockAdminRepository)
    expect    []models.UserDetailsAtAdmin
    expectErr error
  }{
    "success": {
      stub: func(repo mocks.MockAdminRepository) {
        repo.EXPECT().GetUsers().Return(mockUsers, nil).Times(1)
      },
      expect:    mockUsers,
      expectErr: nil,
    },

    "failed": {
      stub: func(repo mocks.MockAdminRepository) {
        repo.EXPECT().GetUsers().Return([]models.UserDetailsAtAdmin{}, errors.New(" Error fetching UserDetails")).Times(1)
      },
      expect:    []models.UserDetailsAtAdmin{},
      expectErr: errors.New("Error fetching UserDetails"),
    },
  }

  for _, testcase := range testCases {

    testcase.stub(*mockadminRepo)
    response, err := mockadminUsecase.GetUsers()

    assert.Equal(t, testcase.expect, response, "got expected test result")
    assert.Equal(t, testcase.expectErr, err, "got expected test error")

  }

}




func TestAddCategory(t *testing.T) {

  ctrl:=gomock.NewController(t)

  catRepo:=mocks.NewMockCategoryRepository(ctrl)
  catusecase:=usecase.NewCategoryUseCase(catRepo)

  mockData:=[]domain.Category{
    {
      ID: 1,
      Category: "Tshirts",
    },{
      ID: 2,
      Category: "Shirts",

    },{
      ID: 3,
      Category: "Jackets",
    },
  }

  // type Category struct {
  // 	ID       uint   `json:"id" gorm:"unique;not null"`
  // 	Category string `json:"category"`
  // }
  //
  testCases:=map[string]struct{
    arg string
    stub func(mocks.MockCategoryRepository)
    expect  domain.Category 
    expectErr error
  }{

    "successCase1":{
      arg:mockData[0].Category,
      stub : func(repo mocks.MockCategoryRepository){
   
         fmt.Println(mockData[0].Category) 
          
        repo.EXPECT().CheckCategory(mockData[0].Category).Return(false,nil).Times(1)
        repo.EXPECT().AddCategory(mockData[0].Category).Return(mockData[0] ,nil).Times(1)
      },
      expect  :   mockData[0] ,
      expectErr: nil,
    },

    "successCase2":{

      arg:mockData[1].Category,

      stub : func(repo mocks.MockCategoryRepository){

        repo.EXPECT().CheckCategory(mockData[1].Category).Return(false,nil).Times(1)
        repo.EXPECT().AddCategory(mockData[1].Category).Return(mockData[1] ,nil).Times(1)
      },
      expect  :   mockData[1] ,
      expectErr: nil,
    },

    "successCase3":{

      arg:mockData[2].Category,
      stub : func(repo mocks.MockCategoryRepository){

        repo.EXPECT().CheckCategory(mockData[2].Category).Return(false,nil).Times(1)
        repo.EXPECT().AddCategory(mockData[2].Category).Return(mockData[2] ,nil).Times(1)
      },
      expect  :   mockData[2] ,
      expectErr: nil,
    },
    "alreadyAddedcase4":{

      arg:mockData[0].Category,
      stub : func(repo mocks.MockCategoryRepository){
   
         fmt.Println(mockData[0].Category) 
          
        repo.EXPECT().CheckCategory(mockData[0].Category).Return(true,errors.New("category already exists")).Times(1)
      },
      expectErr: errors.New("category already exists"),
    },
  }

    t.Parallel()

  for _,testcase := range testCases {
  

    testcase.stub(*catRepo)
    response,err:=catusecase.AddCategory(testcase.arg)
    if err!=nil{
    
      if assert.Error(t, err) {
        assert.Equal(t, testcase.expectErr, err)
      }else{

        t.Errorf("expected %v error but got %v error",testcase.expectErr,err)
      }
      
    }

    t.Log(response)
      
      hm:=assert.Equal(t,testcase.expect,response)
      if !hm{
        t.Errorf("exepected %v but Got %v",testcase.expect,response)
      }
  
  } 



}

// mockgen -source=./pkg/usecase/interface/admin.go -destination=./pkg/mocks/admin/mock_admin_usecase.go -package=mocks
//
// mockgen -source=./pkg/repository/interface/admin.go -destination=./pkg/mocks/admin/mock_admin_repository.go  -package=mocks
