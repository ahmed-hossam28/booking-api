package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var (
	mysql_ *gorm.DB
)

func configMysql() {
	var err error
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	password := os.Getenv("MYSQL_PASSWORD")
	user := os.Getenv("MYSQL_USER")
	dbname := os.Getenv("MYSQL_DB")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)
	mysql_, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("Mysql:", err.Error())
	}
}

func GetMysql() *gorm.DB {
	configMysql()
	return mysql_
}
