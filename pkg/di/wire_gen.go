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
	"project/pkg/repository"
	"project/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	helper := helper.NewHelper(cfg)

	adminRepository := repository.NewAdminRepository(gormDB)
	adminUseCase := usecase.NewAdminUsecase(adminRepository, helper)
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

	serverHTTP := http.NewServerHTTP(userHandler, adminHandler, otpHandler, categoryHandler)

	return serverHTTP, nil

}
