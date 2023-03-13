package main

import (
	"fmt"
	"github.com/Krasav4ik01/ecommerceApp/controllers"
	"github.com/Krasav4ik01/ecommerceApp/initializers"
	"github.com/Krasav4ik01/ecommerceApp/middleware"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"net/http"
)

//это типа urls.py как на django. здесь хранятся все адресса страниц

func handleRequest() {

	//новые функции, пока на доработке
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/logout", controllers.Logout)
	http.HandleFunc("/register", controllers.Register)
	http.HandleFunc("/search", controllers.SearchUsers)

	//http.HandleFunc("/publishTemplate", authcontroller.PublishTemplate)
	//http.HandleFunc("/publish", authcontroller.PublishItem)
	http.HandleFunc("/publish", controllers.CreateProduct)
	http.HandleFunc("/searchProduct", controllers.SearchProducts)

	http.ListenAndServe(":8080", nil)
	fmt.Println("http://localhost:8080")

}

func init() {
	initializers.ConnectDB()
}

// TODO:Основные функции
// Глобалные изменения.
func handleJSONRequests() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:title", controllers.PostShow)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.POST("/register", controllers.SignUp)
	r.POST("/login", controllers.SignIn)
	r.GET("/validate", middleware.AuthRequire, controllers.Validate)
	r.Run()
}

func main() {
	//controllers.PostShowParseTemplate()
	//controllers.PostsIndexParseTemplate()

	handleJSONRequests()

	//handleRequest()
}
