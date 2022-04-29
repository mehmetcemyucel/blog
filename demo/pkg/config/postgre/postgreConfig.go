package postgre

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"helloWorld/pkg/config/config"
)

type DB struct {
	Db  *gorm.DB
	log *zap.Logger
}

func NewDB(c config.PostgreConfig, log *zap.Logger) *DB {
	db, err := initDB(c.Host, c.Port, c.Username, c.Password, c.Name, c.TablePrefix, c.Ssl)
	if err != nil {
		log.Fatal("PostgreSQL couldn't initiate with GORM.", zap.Error(err))
	}
	return &DB{db, log}
}

func initDB(host string, port string, username string, password string, name string, tablePrefix string, ssl bool) (*gorm.DB, error) {
	var dsn string
	if ssl {
		dsn = fmt.Sprintf("host=%s port=%s, user=%s password=%s dbname=%s", host, port, username, password, name)
	} else {
		dsn = fmt.Sprintf("host=%s port=%s, user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, name)
	}
	return gorm.Open(postgres.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		TablePrefix:   tablePrefix,
		SingularTable: false,
	}})
}
