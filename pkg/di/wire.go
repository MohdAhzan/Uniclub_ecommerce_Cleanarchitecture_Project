//go:build wireinject
// +build wireinject

package di

import (
	"fmt"
	http "project/pkg/api"
	"project/pkg/api/handler"
	"project/pkg/config"
	"project/pkg/db"
	"project/pkg/repository"
	"project/pkg/usecase"

	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	fmt.Println("initializing the API")
	wire.Build(db.ConnectDatabase, repository.NewUserRepository, usecase.NewUserUseCase, handler.NewUserHandler, http.NewServerHTTP)
	return &http.ServerHTTP{}, nil
}
