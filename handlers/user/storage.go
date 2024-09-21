package user

import (
	"api-dev/database"
	"api-dev/schemas"
	"log"
)

var (
	//db = database.GetPostgres()
	//db = database.GetMysql()
	//db = database.GetInMemory()
	db = database.GetSqlite()
)

func init() {
	err := db.AutoMigrate(&schemas.User{})
	if err != nil {
		log.Fatal("User Storage: ", err.Error())
	}
}
