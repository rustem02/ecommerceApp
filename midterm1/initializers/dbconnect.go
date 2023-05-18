package initializers

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

// в dsn пишите свой username, password и dbname

func ConnectDB() {
	var err error

	//host := "localhost"
	//user := "postgres"
	//password := "123"
	//dbname := "golang"
	//port := "5432"
	//
	//dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable"
	DB, err = gorm.Open(sqlite.Open("golang.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Couldn't connect to DB")
	}
}
