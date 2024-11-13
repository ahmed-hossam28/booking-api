package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var (
	postgres_ *gorm.DB
)

func configPostgres() {
	var err error
	postgres_, err = gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		logger.Fatal("Postgres: ", err.Error())
	}

}

func GetPostgres() *gorm.DB {
	configPostgres()
	return postgres_
}
