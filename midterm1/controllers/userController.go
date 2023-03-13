package controllers

import (
	"fmt"
	"github.com/Krasav4ik01/ecommerceApp/initializers"
	"github.com/Krasav4ik01/ecommerceApp/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

// var SECRET = "safg73wg7r48wfgosbdfvndusbvusb8"
// Регистрация
// Создаем пользователя спомощью initializers.DB.Create(&user)
func SignUp(c *gin.Context) {
	var data struct {
		FirstName string
		LastName  string
		Address   string
		Email     string `gorm:"unique"`
		Pass      string
	}

	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read data",
		})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Pass), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to hash Password",
		})
		return
	}

	user := models.User{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Address:   data.Address,
		Email:     data.Email,
		Pass:      string(hash),
	}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to register user",
		})
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

// Проверяем есть ли пользователь с таким email и делаем расщифровку пароли, затем генерируем токен, создаем куки файл и входим в систему
func SignIn(c *gin.Context) {
	var data struct {
		Email string `gorm:"unique"`
		Pass  string
	}

	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read data",
		})
		return
	}
	var user models.User
	initializers.DB.First(&user, "email = ?", data.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid email or password",
		})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Pass), []byte(data.Pass))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid email or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create token",
		})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Login", tokenString, 3600*24*30, "", "", false, true)
	fmt.Println(tokenString, err)

	c.JSON(200, gin.H{})
}

// Промежуточная, если удалить куки файл, то пользователь соответвенно выйдет из системы
// Пока есть куки файл, пользователь будет в системе
func Validate(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "You are logged in",
	})
}
