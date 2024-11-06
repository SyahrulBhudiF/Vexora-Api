package core

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

func NewDB(viper *viper.Viper) (*gorm.DB, error) {
	username := viper.GetString("database.user")
	password := viper.GetString("database.password")
	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")
	dbname := viper.GetString("database.dbname")

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s prepareThreshold=0", host, port, username, password, dbname)

	dbLogger := logger.New(log.New(os.Stdout, "\r\n", 0), logger.Config{
		SlowThreshold:             0,
		LogLevel:                  logger.Silent,
		IgnoreRecordNotFoundError: true,
		ParameterizedQueries:      true,
		Colorful:                  false,
	})

	db, err := gorm.Open(
		postgres.New(postgres.Config{DSN: dsn, PreferSimpleProtocol: true}),
		&gorm.Config{Logger: dbLogger},
	)

	if err != nil {
		logrus.Fatal(err)
	}

	return db, nil
}
