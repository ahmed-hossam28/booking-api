package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	sqlite_ *gorm.DB
)

func configSQLite() {
	var err error
	path := "./sqlite-files"
	if _, err = os.Stat(path); err != nil {
		log.Println("creating sqlite database....")
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Fatalln(err)
		}
	}
	sqlite_, err = gorm.Open(sqlite.Open(path+"/main.db"), &gorm.Config{})
	if err != nil {
		logger.Fatal("Sqlite:", err.Error())
	}
}

func GetSqlite() *gorm.DB {
	configSQLite()
	return sqlite_
}
