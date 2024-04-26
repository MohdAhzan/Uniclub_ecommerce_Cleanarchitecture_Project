//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	http "project/pkg/api"
	"project/pkg/api/handler"
	"project/pkg/config"
	"project/pkg/db"
	"project/pkg/helper"
	"project/pkg/redis"
	"project/pkg/repository"
	"project/pkg/usecase"
	// "project/pkg/redis"
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

	helper := helper.NewHelper(cfg)

	adminRepository := repository.NewAdminRepository(gormDB)
	orderRepository := repository.NewOrderRepository(gormDB)
	adminUseCase := usecase.NewAdminUsecase(adminRepository, helper, orderRepository)
	adminHandler := handler.NewAdminHandler(adminUseCase)

	userRepository := repository.NewUserRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository, cfg, helper)
	userHandler := handler.NewUserHandler(userUseCase)

	otpRepository := repository.NewOtpRepository(gormDB)
	otpUseCase := usecase.NewOtpUseCase(cfg, otpRepository, helper)
	otpHandler := handler.NewOtpHandler(otpUseCase)

	categoryRepository := repository.NewCategoryRepository(gormDB)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryUseCase)

	inventoryRepository := repository.NewInventoryRepository(gormDB, redisClient)
	inventoryUsecase := usecase.NewInventoryUseCase(inventoryRepository, helper)
	inventoryHandler := handler.NewInventoryHandler(inventoryUsecase, redisClient)

	cartRepository := repository.NewCartRepository(gormDB)
	cartUseCase := usecase.NewCartUseCase(cartRepository, inventoryRepository)
	cartHandler := handler.NewCartHandler(cartUseCase)

	orderRepository = repository.NewOrderRepository(gormDB)
	orderUseCase := usecase.NewOrderUseCase(orderRepository, cartRepository, cartUseCase, userRepository, helper)
	orderHandler := handler.NewOrderHandler(orderUseCase)

	paymentRepository := repository.NewPaymentRepository(gormDB)
	paymentUseCase := usecase.NewPaymentUseCase(paymentRepository, cfg, orderRepository)
	paymentHandler := handler.NewPaymentHandler(paymentUseCase)

	serverHTTP := http.NewServerHTTP(userHandler, adminHandler, otpHandler,
		categoryHandler, inventoryHandler, cartHandler, orderHandler, paymentHandler)

	return serverHTTP, nil

}
