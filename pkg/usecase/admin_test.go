package usecase_test

import (
	"errors"
	mocks "project/pkg/mocks/admin"
	"project/pkg/utils/models"
	"testing"

	"project/pkg/usecase"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)


func TestGetUser (t *testing.T){

  mockctrl:=gomock.NewController(t)    
  mockadminRepo:=mocks.NewMockAdminRepository(mockctrl) 
  mockhelper:=mocks.NewMockHelper(mockctrl)  
  mockorderRepo:=mocks.NewMockOrderRepository(mockctrl)  
  mockuserRepo:=mocks.NewMockUserRepository(mockctrl) 

  mockadminUsecase:=usecase.NewAdminUsecase(mockadminRepo,mockhelper,mockorderRepo,mockuserRepo)

  // Define what the mock should return when GetUsers is called

  mockUsers := []models.UserDetailsAtAdmin{
    {	Id : 1, Name    :"Demouser1", Email  :"demouser1@gmail.com", Phone   :"9878362828", Blocked : false, },
    {	Id : 2, Name    :"Demouser2", Email  :"demouser2@gmail.com", Phone   :"9999999999", Blocked : true, },
  }

  testCases:=map[string]struct{
    stub func(mocks.MockAdminRepository)
    expect []models.UserDetailsAtAdmin
    expectErr error
  }{
    "success" : {
      stub: func(repo mocks.MockAdminRepository){
        repo.EXPECT().GetUsers().Return(mockUsers,nil).Times(1)
      },
      expect: mockUsers,
      expectErr: nil,
    },

    "failed" : {
      stub: func(repo mocks.MockAdminRepository){
        repo.EXPECT().GetUsers().Return([]models.UserDetailsAtAdmin{},errors.New(" Error fetching UserDetails")).Times(1)
      },
      expect:[]models.UserDetailsAtAdmin{} ,
      expectErr:errors.New("Error fetching UserDetails"),
    },

  }

      for _,testcase:=range testCases{

        testcase.stub(*mockadminRepo)
        response,err:=mockadminUsecase.GetUsers()

        assert.Equal(t,testcase.expect,response,"got expected test result")
        assert.Equal(t,testcase.expectErr,err,"got expected test error")
        
    }  


}

// mockgen -source=./pkg/usecase/interface/admin.go -destination=./pkg/mocks/admin/mock_admin_usecase.go -package=mocks
//
// mockgen -source=./pkg/repository/interface/admin.go -destination=./pkg/mocks/admin/mock_admin_repository.go  -package=mocks
