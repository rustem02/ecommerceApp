package main

import (
	"fmt"
	db "github.com/Krasav4ik01/ecommerceApp/config"
	"github.com/Krasav4ik01/ecommerceApp/initializers"
	"github.com/Krasav4ik01/ecommerceApp/routes"
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
)

func init() {
	initializers.ConnectDB()
}

// Основные функции
// Глобалные изменения.

func main() {
	//handleJSONRequests()
	fmt.Println("Запуск проекта...")
	//db connection
	db.Connect()

	app := fiber.New()
	app.Use(cors.New())
	//routing
	routes.Setup(app)
	app.Listen(":8080")

}
