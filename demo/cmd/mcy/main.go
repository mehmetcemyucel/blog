package main

import (
	"go.uber.org/zap"
	"helloWorld/pkg/api"
	"helloWorld/pkg/config/config"
	"helloWorld/pkg/config/migrate"
	"helloWorld/pkg/config/postgre"
	"helloWorld/pkg/logging"
	"helloWorld/pkg/router"
)

// @title Golang Example API
// @version 1.0
// @description This is an example microservice with Golang
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	config := config.NewConfig()
	log := logging.NewLog(config.Log)
	migrate := migrate.NewMig(config.Postgre, log)
	db := postgre.NewDB(config.Postgre, log)
	api := api.NewApi(config, db, log)
	server := router.NewServer(api, config, log)
	app := App{config, server, db, api, migrate}
	log.Info("App Server Starting...")
	app.run(log)
}

type App struct {
	Config *config.Config
	Server *router.Echo
	Db     *postgre.DB
	Api    *api.Api
	Mig    *migrate.Mig
}

func (a *App) run(log *zap.Logger) {
	a.Server.RunEcho()
}
