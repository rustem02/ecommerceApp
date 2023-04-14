package main

import (
	db "github.com/Krasav4ik01/ecommerceApp/config"
	"github.com/Krasav4ik01/ecommerceApp/initializers"
	"github.com/Krasav4ik01/ecommerceApp/models"
)

// Инициализация БД
func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	//config.DBConn()
	db.Connect()
}

// Для миграции go run migrate/migrate.go
func main() {
	initializers.DB.AutoMigrate(&models.Post{}, &models.User{})
	//config.DB.AutoMigrate(&entities.User{}, &entities.Product{})
	db.DB.AutoMigrate(&models.Cashier{}, &models.Category{}, &models.Payment{}, &models.PaymentType{}, &models.Product{}, &models.Discount{}, &models.Order{}, &models.Comment{}, &models.Rating{})

}
