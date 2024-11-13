package database

import (
	"github.com/ahmed-hossam28/config"
	"github.com/joho/godotenv"
)

var logger = config.NewSingletonLogger()

func init() {
	err := godotenv.Load()
	if err != nil {
		logger.Fatal(err.Error())
	}
}
