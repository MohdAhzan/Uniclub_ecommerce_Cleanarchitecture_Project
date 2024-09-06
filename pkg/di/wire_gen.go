//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	http "project/pkg/api"
	"project/pkg/api/handler"
	"project/pkg/api/middleware"
	"project/pkg/config"
	"project/pkg/db"
	"project/pkg/helper"
	"project/pkg/redis"
	"project/pkg/repository"
	"project/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	redisClient, err := redis.InitializeClient()
	if err != nil {
		return nil, err
	}
  
  middleware.CfgHelper(cfg)


	helper := helper.NewHelper(cfg)

	adminRepository := repository.NewAdminRepository(gormDB)
	orderRepository := repository.NewOrderRepository(gormDB)
	userRepository := repository.NewUserRepository(gormDB)
	adminUseCase := usecase.NewAdminUsecase(adminRepository, helper, orderRepository, userRepository)
	adminHandler := handler.NewAdminHandler(adminUseCase, helper)

	userRepository = repository.NewUserRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository, cfg, helper)
	userHandler := handler.NewUserHandler(userUseCase)

	otpRepository := repository.NewOtpRepository(gormDB)
	otpUseCase := usecase.NewOtpUseCase(cfg, otpRepository, helper)
	otpHandler := handler.NewOtpHandler(otpUseCase)

	categoryRepository := repository.NewCategoryRepository(gormDB)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryUseCase)

	couponRepository := repository.NewCouponRepository(gormDB)
	couponUsecase := usecase.NewCouponUseCase(couponRepository)
	couponHandler := handler.NewCouponHandler(couponUsecase)

	inventoryRepository := repository.NewInventoryRepository(gormDB, redisClient)
	offerRepository := repository.NewOfferRepository(gormDB)
	inventoryUsecase := usecase.NewInventoryUseCase(inventoryRepository, helper, offerRepository)
	inventoryHandler := handler.NewInventoryHandler(inventoryUsecase, redisClient)

	cartRepository := repository.NewCartRepository(gormDB)
	cartUseCase := usecase.NewCartUseCase(cartRepository, inventoryRepository, offerRepository)
	cartHandler := handler.NewCartHandler(cartUseCase)

	orderRepository = repository.NewOrderRepository(gormDB)
	orderUseCase := usecase.NewOrderUseCase(orderRepository, cartRepository, cartUseCase, userRepository, helper, couponRepository, offerRepository, inventoryRepository)
	orderHandler := handler.NewOrderHandler(orderUseCase)

	paymentRepository := repository.NewPaymentRepository(gormDB)
	paymentUseCase := usecase.NewPaymentUseCase(paymentRepository, cfg, orderRepository, userRepository)
	paymentHandler := handler.NewPaymentHandler(paymentUseCase)

	wishlistRepository := repository.NewWishlistRepository(gormDB)
	wishlistUsecase := usecase.NewWishlistUsecase(wishlistRepository, inventoryRepository)
	wishlistHandler := handler.NewWishlistHandler(wishlistUsecase)

	offerRepository = repository.NewOfferRepository(gormDB)
	offerUseCase := usecase.NewOfferUseCase(offerRepository, categoryRepository, inventoryRepository)
	offerHandler := handler.NewOfferHandler(offerUseCase)

	serverHTTP := http.NewServerHTTP(userHandler, adminHandler, otpHandler,
		categoryHandler, inventoryHandler, cartHandler, orderHandler, paymentHandler, wishlistHandler, offerHandler, couponHandler)

	return serverHTTP, nil

}
