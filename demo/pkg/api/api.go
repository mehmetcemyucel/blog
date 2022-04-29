package api

import (
	"go.uber.org/zap"
	"helloWorld/pkg/config/config"
	"helloWorld/pkg/config/postgre"
	"helloWorld/pkg/repository"
	"helloWorld/pkg/service"
)

type Api struct {
	ExampleApi ExampleApi
	config     *config.Config
	log        *zap.Logger
}

func NewApi(config *config.Config, postgre *postgre.DB, log *zap.Logger) *Api {
	exampleApi := initExampleApi(postgre, log)
	return &Api{ExampleApi: exampleApi, config: config, log: log}
}

func initExampleApi(pDb *postgre.DB, log *zap.Logger) ExampleApi {
	r := repository.NewRepository(pDb.Db)
	s := service.NewExampleService(r, log)
	a := NewExampleApi(s, log)
	return a
}
