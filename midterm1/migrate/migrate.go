package main

import (
	"github.com/Krasav4ik01/ecommerceApp/initializers"
	"github.com/Krasav4ik01/ecommerceApp/models"
)

// Инициализация БД
func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	//config.DBConn()
}

// Для миграции go run migrate/migrate.go
func main() {
	initializers.DB.AutoMigrate(&models.Post{}, &models.User{})
	//config.DB.AutoMigrate(&entities.User{}, &entities.Product{})
}
