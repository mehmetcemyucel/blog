package migrate

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.uber.org/zap"
	"helloWorld/pkg/config/config"
)

type Mig struct {
	Migration *migrate.Migrate
}

func NewMig(c config.PostgreConfig, log *zap.Logger) *Mig {

	fullConnStr := "postgres://" + c.Username + ":" + c.Password + "@" + c.Host + ":" + c.Port + "/" + c.Name
	if !c.Ssl {
		fullConnStr = fullConnStr + "?sslmode=disable"
	}
	log.Info("Migration has started with; ", zap.String("connUrl", fullConnStr))

	var m, err = migrate.New("file://db", fullConnStr)
	if err != nil {
		log.Fatal("Migration create has an error", zap.Error(err))
	}
	if err := m.Up(); err != nil {
		log.Warn("Migration has an error ", zap.Error(err))
	}
	log.Info("Migration finished successfully")
	return &Mig{m}
}
