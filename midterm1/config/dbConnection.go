package db

import (
	"fmt"
	"github.com/Krasav4ik01/ecommerceApp/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	//godotenv.Load()
	//DbHost := os.Getenv("localhost")
	//DbName := os.Getenv("golang")
	//DbUsername := os.Getenv("rustem")
	//DbPassword := os.Getenv("123")
	//
	//connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbUsername, DbPassword, DbHost, DbName)
	dbConnection, err := gorm.Open(sqlite.Open("golang.db"), &gorm.Config{})

	if err != nil {
		panic("connection failed to the database ")
	}
	DB = dbConnection
	fmt.Println("db connected successfully")

	AutoMigrate(dbConnection)
	//if err := DB.AutoMigrate(&models.Cashier{}, &models.Category{}, &models.Payment{}, &models.PaymentType{}, &models.Product{}, &models.Discount{}, &models.Order{}).Error; err != nil {
	//	log.Fatalf("Migration failed %v", err)
	//}

}

func AutoMigrate(connection *gorm.DB) {
	connection.Debug().AutoMigrate(&models.Cashier{}, &models.Category{}, &models.Payment{}, &models.PaymentType{}, &models.Product{}, &models.Discount{}, &models.Order{}, &models.Comment{}, &models.Rating{})
}
