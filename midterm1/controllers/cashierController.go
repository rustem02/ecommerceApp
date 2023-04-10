package controllers

import (
	db "github.com/Krasav4ik01/ecommerceApp/config"
	"github.com/Krasav4ik01/ecommerceApp/models"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

//func CreateCashier(ctx *gin.Context) {
//	var body struct {
//		Name     string
//		Email    string `gorm:"unique"`
//		Passcode string
//	}
//
//	if ctx.Bind(&body) != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{
//			"message": "Failed to read data",
//		})
//		return
//	}
//
//	cashier := models.Cashier{
//		Name:     body.Name,
//		Email:    body.Email,
//		Passcode: body.Passcode,
//	}
//	result := db.DB.Create(&cashier)
//
//	if result.Error != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{
//			"message": "Failed to create cashier",
//		})
//		return
//	}
//
//	ctx.JSON(200, gin.H{
//		"cashier": cashier,
//	})
//
//	return
//}

func CreateCashier(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		log.Fatalf("registeration error in post request %v", err)
	}

	if data["firstName"] == "" || data["lastName"] == "" || data["email"] == "" || data["passcode"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Cashier Name is required",
			"error":   map[string]interface{}{},
		})
	}
	//passCode := strconv.Itoa(rand.Intn(1000000))
	//fmt.Println("passCode:::", passCode
	cashier := models.Cashier{
		FirstName: data["firstName"],
		LastName:  data["lastName"],
		Email:     data["email"],
		Passcode:  data["passcode"],
	}
	db.DB.Create(&cashier)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    cashier,
	})
}

func GetCashierDetails(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")

	var cashier models.Cashier
	db.DB.Select("id ,firstName, lastName, email").Where("id=?", cashierId).First(&cashier)
	cashierData := make(map[string]interface{})
	cashierData["cashierId"] = cashier.Id
	cashierData["firstName"] = cashier.FirstName
	cashierData["lastName"] = cashier.LastName
	cashierData["email"] = cashier.Email

	if cashier.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Cashier Not Found",
			"error":   map[string]interface{}{},
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    cashierData,
	})
}

func DeleteCashier(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")
	var cashier models.Cashier
	db.DB.Where("id=?", cashierId).First(&cashier)

	if cashier.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Cashier Not Found",
			"error":   map[string]interface{}{},
		})
	}
	db.DB.Where("id = ?", cashierId).Delete(&cashier)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
	})
}

func UpdateCashier(c *fiber.Ctx) error {

	cashierId := c.Params("cashierId")
	var cashier models.Cashier

	db.DB.Find(&cashier, "id = ?", cashierId)
	if cashier.FirstName == "" {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Cashier Not Found",
		})
	}

	var updateCashierData models.Cashier
	c.BodyParser(&updateCashierData)
	if updateCashierData.FirstName == "" || updateCashierData.LastName == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Cashier name is required",
			"error":   map[string]interface{}{},
		})
	}

	cashier.FirstName = updateCashierData.FirstName
	cashier.LastName = updateCashierData.LastName
	db.DB.Save(&cashier)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    cashier,
	})

}

// структура Cashiers
type Cashiers struct {
	Id        int    `json:"cashierId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func CashiersList(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit"))
	skip, _ := strconv.Atoi(c.Query("skip"))
	var count int64
	var cashier []Cashiers
	db.DB.Select("*").Limit(limit).Offset(skip).Find(&cashier).Count(&count)
	metaMap := map[string]interface{}{
		"total": count,
		"limit": limit,
		"skip":  skip,
	}
	cashiersData := map[string]interface{}{
		"cashiers": cashier,
		"meta":     metaMap,
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"Message": "Success",
		"data":    cashiersData,
	})

}
