package controllers

import (
	"github.com/Krasav4ik01/ecommerceApp/initializers"
	"github.com/Krasav4ik01/ecommerceApp/models"
	"github.com/gin-gonic/gin"
)

// Для создании поста
func PostsCreate(c *gin.Context) {

	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})

}

// Функция для просмотра всех постов
func PostsIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	//c.HTML(http.StatusOK, "posts.html", gin.H{
	//	"posts": posts,
	//})
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

// TODO: Начал вытаскивать данные на фронтенд
func PostsIndexParseTemplate() {
	app := gin.Default()

	app.LoadHTMLGlob("templates/posts/posts.html")
	//app.Static("/assets", "./assets")

	app.GET("/post", PostsIndex)

	app.Run()
}

// Для просмотра функции по title. или же можно назвать search items based on name
func PostShow(c *gin.Context) {
	title := c.Param("title")
	var post models.Post
	//initializers.DB.First(&post, id)
	initializers.DB.Where("title = ?", title).First(&post)

	//c.HTML(http.StatusOK, "detailPost.html", gin.H{
	//	"post": post,
	//})
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostShowParseTemplate() {
	app := gin.Default()

	app.LoadHTMLGlob("templates/posts/detailPost.html")
	//app.Static("/assets", "./assets")

	app.GET("/post/:id", PostShow)

	app.Run()
}

// Функция для Обнавлении поста по id
func PostsUpdate(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})
	c.JSON(200, gin.H{
		"post": post,
	})
}

// Для удалении поста по id
func PostsDelete(c *gin.Context) {
	id := c.Param("id")

	//var post models.Post
	initializers.DB.Delete(&models.Post{}, id)

	//c.JSON(200, gin.H{
	//	"post": post,
	//})
	c.Status(200)
}
