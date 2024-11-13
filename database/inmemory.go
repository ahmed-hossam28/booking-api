package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var inMemory *gorm.DB

func configInMemoryDB() {
	var err error
	inMemory, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		logger.Fatal("InMemory:", err.Error())
	}
}

func GetInMemory() *gorm.DB {
	configInMemoryDB()
	return inMemory
}
