package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var inMemory *gorm.DB

func configInMemoryDB() {
	var err error
	inMemory, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatal("InMemory:", err)
	}
}

func GetInMemory() *gorm.DB {
	configInMemoryDB()
	return inMemory
}
