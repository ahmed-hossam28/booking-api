package user

import (
	"github.com/ahmed-hossam28/database"
	"github.com/ahmed-hossam28/schemas"
	"log"
)

var (
	//db = database.GetPostgres()
	db = database.GetMysql()
	//db = database.GetInMemory()
	//db = database.GetSqlite()
)

func init() {
	err := db.AutoMigrate(&schemas.User{})
	if err != nil {
		log.Fatal("User Storage: ", err.Error())
	}
}
