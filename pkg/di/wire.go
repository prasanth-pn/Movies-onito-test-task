//go:build wireinject
//+build wireinject



package di

import (
	"github.com/google/wire"
	"github.com/prasanth-pn/Movies-onito-test-task/pkg/api/handler"
	http "github.com/prasanth-pn/Movies-onito-test-task/pkg/api"
	"github.com/prasanth-pn/Movies-onito-test-task/pkg/config"
	"github.com/prasanth-pn/Movies-onito-test-task/pkg/db"
	"github.com/prasanth-pn/Movies-onito-test-task/pkg/repository"
	"github.com/prasanth-pn/Movies-onito-test-task/pkg/usecase"
)

func InitializeEvent(cfg config.Config)( *http.ServerHttp,error){
	wire.Build(
		db.ConnectDB,
		repository.NewMovieRepository,
		usecase.NewMovieUseCase,
		handler.NewMovieHandler,
		http.NewServerHttp,
		
	)
	return &http.ServerHttp{},nil
}
