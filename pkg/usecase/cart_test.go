package usecase_test

import (
	mocks "project/pkg/mocks/admin"
	"project/pkg/utils/models"
	"testing"

	"project/pkg/usecase"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// var mockData struct{
//   pid int
//   userId int
//   quantity int
// }

// func TestAddtoCart(t *testing.T){
//   mockCtrl:=gomock.NewController(t)
//
//   mockInvRepo:=mocks.NewMockInventoryRepository(mockCtrl)
//   mockCartRepo:=mocks.NewMockCartRepository(mockCtrl)
//   mockOfferRepo:=mocks.NewMockOfferRepository(mockCtrl)
//
//   cartUsecase:=usecase.NewCartUseCase(mockCartRepo,mockInvRepo,mockOfferRepo)
//
// 	// ProductID       int     `json:"product_id"`
// 	// ProductName     string  `json:"product_name"`
// 	// Image           string  `json:"image"`
// 	// Category_id     int     `json:"category_id"`
// 	// Quantity        int     `json:"quantity"`
// 	// StockAvailable  int     `json:"stock_available"`
// 	// TotalPrice      float64 `json:"total_price"`
// 	// CategoryOffer   string  `json:"category_offer"`
// 	// ProductOffer    string  `json:"product_offer"`
// 	// DiscountedPrice float64 `json:"discounted_price"`
//
//   mockUsers := []models.GetCart{
//     {ProductID: 1,ProductName: "404 zero Tshirts ",Image:"path/someloc/some/image.png",Category_id:k},
//     {},
//   }
//
//   mockData:=[]struct{pid,userId,quantity int }{
//     {1,1,2},{2,1,1},
//   }
//
//   testCases:=map[string]struct{
//
//     stub func(mocks.MockCartRepository,mocks.MockInventoryRepository,mocks.MockOfferRepository)
//     expect []models.CartResponse
//     expectErr error
//   }{
//     "success":{
//       stub: func (cartRepo mocks.MockCartRepository,invRepo mocks.MockInventoryRepository,offRepo mocks.MockOfferRepository){
//         invRepo.EXPECT().CheckStock(1).Return(10,nil).Times(1)
//         cartRepo.EXPECT().GetCartID(1).Return(1,nil).Times(1)
//         cartRepo.EXPECT().CheckIfItemIsAlreadyAdded(1,1).Return(true,nil).Times(1)
//         cartRepo.EXPECT().FindCartQuantity(1,1).Return(2,nil).Times(1)
//         cartRepo.EXPECT().UpdateCartQuantity(1,1,4).Return(nil).Times(1)
//       },
//       //
//       // expect : // hmm i need something to test with like an expected result model but currently have no clue on it
//       //
//       //
//
//
//     },
//   }
//
//
// }

func TestAddtoCart(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockInvRepo := mocks.NewMockInventoryRepository(mockCtrl)
	mockCartRepo := mocks.NewMockCartRepository(mockCtrl)
	mockOfferRepo := mocks.NewMockOfferRepository(mockCtrl)

	cartUsecase := usecase.NewCartUseCase(mockCartRepo, mockInvRepo, mockOfferRepo)
	//
	// mockData := []struct {
	// 	pid, userID, quantity int
	// }{
	// 	{1, 1, 2},
	// 	{2, 1, 1},
	// }

	testCases := map[string]struct {
		stub      func(mockCartRepo *mocks.MockCartRepository, mockInvRepo *mocks.MockInventoryRepository, mockOfferRepo *mocks.MockOfferRepository)
		expect    models.CartResponse
		expectErr error
	}{
		"success": {
			stub: func(cartRepo *mocks.MockCartRepository, invRepo *mocks.MockInventoryRepository, offRepo *mocks.MockOfferRepository) {
				invRepo.EXPECT().CheckStock(1).Return(10, nil).Times(1)
				cartRepo.EXPECT().GetCartID(1).Return(1, nil).Times(1)
				cartRepo.EXPECT().CheckIfItemIsAlreadyAdded(1, 1).Return(true, nil).Times(1)
				cartRepo.EXPECT().FindCartQuantity(1, 1).Return(2, nil).Times(1)
				cartRepo.EXPECT().UpdateCartQuantity(1, 1, 4).Return(nil).Times(1)

				// Mock offer and price calculations
				invRepo.EXPECT().FindPrice(1).Return(100.0, nil).Times(1)
				invRepo.EXPECT().GetProductImages(1).Return("path/someloc/some/image.png", nil).Times(1)
				offRepo.EXPECT().GetCategoryOfferDiscountPercentage(1).Return(10.0, "Category Offer", nil).Times(1)
				offRepo.EXPECT().GetInventoryOfferDiscountPercentage(1).Return(5.0, "Product Offer", nil).Times(1)
			},
			expect: models.CartResponse{
				CartID: 1,
				CartData: []models.GetCart{
					{
						ProductID:       1,
						ProductName:     "404 zero Tshirts",
						Image:           "path/someloc/some/image.png",
						Category_id:     1,
						Quantity:        4,
						StockAvailable:  6,
						TotalPrice:      400.0, // 100 * 4
						CategoryOffer:   "Category Offer",
						ProductOffer:    "Product Offer",
						DiscountedPrice: 340.0, // (400 - (400 * 15%))
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
			result, err := cartUsecase.AddtoCart(1, 1, 2)

			// Check if the error is what we expect
			if err != nil {
				if testcase.expectErr != nil && err.Error() != testcase.expectErr.Error() {
					t.Errorf("expected error %v, got %v", testcase.expectErr, err)
				}
			}

        assert.Equal(t,testcase.expect,result)
			//    {
			// 	t.Errorf("expected %v, got %v", tc.expect, result)
			// }
		})
	}
}


//
// mockgen -source=./pkg/repository/interface/admin.go -destination=./pkg/mocks/admin/mock_admin_repository.go  -package=mocks
