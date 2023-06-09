// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/prasanth-pn/Movies-onito-test-task/pkg/api"
	"github.com/prasanth-pn/Movies-onito-test-task/pkg/api/handler"
	"github.com/prasanth-pn/Movies-onito-test-task/pkg/config"
	"github.com/prasanth-pn/Movies-onito-test-task/pkg/db"
	"github.com/prasanth-pn/Movies-onito-test-task/pkg/repository"
	"github.com/prasanth-pn/Movies-onito-test-task/pkg/usecase"
)

// Injectors from wire.go:

func InitializeEvent(cfg config.Config) (*http.ServerHttp, error) {
	sqlDB := db.ConnectDB(cfg)
	moviesRepoInterfaces := repository.NewMovieRepository(sqlDB)
	movieusecaseInterface := usecase.NewMovieUseCase(moviesRepoInterfaces)
	moviesHandler := handler.NewMovieHandler(movieusecaseInterface)
	serverHttp := http.NewServerHttp(moviesHandler)
	return serverHttp, nil
}
