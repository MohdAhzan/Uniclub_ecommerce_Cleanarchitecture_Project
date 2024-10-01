package usecase_test

import (
  mocks "project/pkg/mocks/admin"
  "project/pkg/utils/models"
  "testing"

  "project/pkg/usecase"

  "github.com/golang/mock/gomock"
  "github.com/stretchr/testify/assert"
)


func TestAddtoCart(t *testing.T) {
  mockCtrl := gomock.NewController(t)
  // defer mockCtrl.Finish()

  mockInvRepo := mocks.NewMockInventoryRepository(mockCtrl)
  mockCartRepo := mocks.NewMockCartRepository(mockCtrl)
  mockOfferRepo := mocks.NewMockOfferRepository(mockCtrl)

  cartUsecase := usecase.NewCartUseCase(mockCartRepo, mockInvRepo, mockOfferRepo)



  mockData := []struct {
    pid,pdtStock, userID, quantity,cartId,cartQuantity,category_id,updatedCartQuantity int
    ProductName,Image,CategoryOffer,ProductOffer string
    CategoryDiscountRate,ProductDiscountRate,price float64
  }{
    {1,10, 1, 2,1,2,2,4,"404 zero Tshirts","path/someloc/some/image.png","CATOFF5","PDTOFF10",5,10,2000},
    {1,10, 1, 2,0,0,1,2,"404 zero Tshirts","path/someloc/some/image2.png","CATOFF5","PDTOFF10",5,10,1500},
  }

  testCases := map[string]struct {
    stub      func(mockCartRepo *mocks.MockCartRepository, mockInvRepo *mocks.MockInventoryRepository, mockOfferRepo *mocks.MockOfferRepository)
    expect    models.CartResponse
    expectErr error
  }{
    "success": {
      stub: func(cartRepo *mocks.MockCartRepository, invRepo *mocks.MockInventoryRepository, offRepo *mocks.MockOfferRepository) {
        invRepo.EXPECT().CheckStock(mockData[0].pid).Return(mockData[0].pdtStock, nil).Times(1)
        cartRepo.EXPECT().GetCartID(mockData[0].userID).Return(mockData[0].cartId, nil).Times(1)
        cartRepo.EXPECT().CheckIfItemIsAlreadyAdded(mockData[0].cartId,mockData[0].pid).Return(true, nil).Times(1)
        cartRepo.EXPECT().FindCartQuantity(mockData[0].pid,mockData[0].cartId).Return(mockData[0].cartQuantity, nil).Times(1)
        cartRepo.EXPECT().UpdateCartQuantity(mockData[0].cartId, mockData[0].pid, mockData[0].updatedCartQuantity).Return(nil).Times(1)
        cartRepo.EXPECT().GetProductNames(mockData[0].pid).Return(mockData[0].ProductName,nil).Times(1)
        invRepo.EXPECT().GetProductImages(mockData[0].pid).Return(mockData[0].Image, nil).Times(1)
        invRepo.EXPECT().GetCategoryID(mockData[0].pid).Return(mockData[0].category_id,nil).Times(1)

        cartRepo.EXPECT().FindCartQuantity(mockData[0].pid,mockData[0].cartId).Return(mockData[0].updatedCartQuantity, nil).Times(1)

        invRepo.EXPECT().FindPrice(mockData[0].pid).Return(mockData[0].price,nil).Times(1)
        offRepo.EXPECT().GetCategoryOfferDiscountPercentage(mockData[0].category_id).Return(mockData[0].CategoryDiscountRate,mockData[0].CategoryOffer,nil).Times(1)
        offRepo.EXPECT().GetInventoryOfferDiscountPercentage(mockData[0].pid).Return(mockData[0].ProductDiscountRate,mockData[0].ProductOffer,nil).Times(1)
      },
      expect: models.CartResponse{
        CartID: 1,
        CartData: []models.GetCart{
          {
            ProductID:       1,
            ProductName:     mockData[0].ProductName,
            Image:           mockData[0].Image,
            Category_id:    mockData[0].category_id,
            Quantity:        mockData[0].updatedCartQuantity,
            StockAvailable:  6,//10-4
            TotalPrice:      8000.0, // 2000 * 4
            CategoryOffer:  mockData[0].CategoryOffer,
            ProductOffer:   mockData[0].ProductOffer,
            DiscountedPrice: 6800, // (8000 - (8000 * 15%))
          },
        },
      },
      expectErr: nil,
    },
    "user_withnoitemsinCart": {
      stub: func(cartRepo *mocks.MockCartRepository, invRepo *mocks.MockInventoryRepository, offRepo *mocks.MockOfferRepository) {
        invRepo.EXPECT().CheckStock(mockData[1].pid).Return(mockData[1].pdtStock, nil).Times(1)
        cartRepo.EXPECT().GetCartID(mockData[1].userID).Return(mockData[1].cartId, nil).Times(1)
        //created and simulated the new cartId creation
        mockData[1].cartId=2

        cartRepo.EXPECT().CreateNewCart(mockData[1].userID).Return(mockData[1].cartId,nil).Times(1)
        cartRepo.EXPECT().CheckIfItemIsAlreadyAdded(mockData[1].cartId,mockData[1].pid).Return(false, nil).Times(1)
        cartRepo.EXPECT().AddtoCartItems(mockData[1].cartId, mockData[1].pid, mockData[1].updatedCartQuantity).Return(nil).Times(1)
        cartRepo.EXPECT().GetProductNames(mockData[1].pid).Return(mockData[1].ProductName,nil).Times(1)
        invRepo.EXPECT().GetProductImages(mockData[1].pid).Return(mockData[1].Image, nil).Times(1)
        invRepo.EXPECT().GetCategoryID(mockData[1].pid).Return(mockData[1].category_id,nil).Times(1)

        cartRepo.EXPECT().FindCartQuantity(mockData[1].pid,mockData[1].cartId).Return(mockData[1].updatedCartQuantity, nil).Times(1)

        invRepo.EXPECT().FindPrice(mockData[1].pid).Return(mockData[1].price,nil).Times(1)
        offRepo.EXPECT().GetCategoryOfferDiscountPercentage(mockData[1].category_id).Return(mockData[1].CategoryDiscountRate,mockData[1].CategoryOffer,nil).Times(1)
        offRepo.EXPECT().GetInventoryOfferDiscountPercentage(mockData[1].pid).Return(mockData[1].ProductDiscountRate,mockData[1].ProductOffer,nil).Times(1)
      },
      expect: models.CartResponse{
        CartID: 2,
        CartData: []models.GetCart{
          {
            ProductID:       1,
            ProductName:     mockData[1].ProductName,
            Image:           mockData[1].Image,
            Category_id:    mockData[1].category_id,
            Quantity:        mockData[1].updatedCartQuantity,
            StockAvailable:  8,//10-2
            TotalPrice:      3000.0, // 1500 * 2
            CategoryOffer:  mockData[1].CategoryOffer,
            ProductOffer:   mockData[1].ProductOffer,
            DiscountedPrice: 2550, // (3000 - (3000 * 15%))
          },
        },
      },
      expectErr: nil,
    },

  }

  for name, testcase := range testCases {
    t.Run(name, func(t *testing.T) {
      testcase.stub(mockCartRepo, mockInvRepo, mockOfferRepo)

      // Call the actual function
      cartresponse, err := cartUsecase.AddtoCart(1, 1, 2)

      // Check if the error is what we expect
      if err != nil {
        if testcase.expectErr != nil && err.Error() != testcase.expectErr.Error() {
          t.Errorf("expected error %v, got %v", testcase.expectErr, err)
        }
      }
      t.Logf("TestCartresponse is %v",cartresponse)
      assert.Equal(t,testcase.expect,cartresponse)
      //    {
      // 	t.Errorf("expected %v, got %v", tc.expect, result)
      // }
    })
  }
}


//
// mockgen -source=./pkg/repository/interface/admin.go -destination=./pkg/mocks/admin/mock_admin_repository.go  -package=mocks
