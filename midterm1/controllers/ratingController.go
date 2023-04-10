package controllers

import (
	db "github.com/Krasav4ik01/ecommerceApp/config"
	"github.com/Krasav4ik01/ecommerceApp/models"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
	"time"
)

type RatingStruct struct {
	RatingId   int `json:"ratingId"`
	CashierId  int `json:"cashierId"`
	ProductId  int `json:"productId"`
	ProdRating int `json:"prodRating"`
}

func CreateRating(c *fiber.Ctx) error {
	var data RatingStruct
	err := c.BodyParser(&data)
	if err != nil {
		log.Fatalf("Product error in post request %v", err)
	}

	rating := models.Rating{
		CashierId:     data.CashierId,
		ProductId:     data.ProductId,
		ProductRating: data.ProdRating,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
	}
	db.DB.Create(&rating)

	Response := map[string]interface{}{
		"success": true,
		"message": "Success",
		"data":    rating,
	}
	return (c.JSON(Response))

}

func RatingList(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit"))
	skip, _ := strconv.Atoi(c.Query("skip"))
	var count int64
	var rating []models.Rating

	db.DB.Select("*").Limit(limit).Offset(skip).Find(&rating).Count(&count)

	type RatingList struct {
		RatingId   int            `json:"ratingId"`
		CashierID  int            `json:"cashiersId"`
		ProductID  int            `json:"productId"`
		ProdRating float32        `json:"prodRating"`
		CreatedAt  time.Time      `json:"createdAt"`
		Cashiers   models.Cashier `json:"cashier"`
		Product    models.Product `json:"product"`
	}
	RatingsResponse := make([]*RatingList, 0)

	for _, v := range rating {
		cashier := models.Cashier{}
		db.DB.Where("id = ?", v.CashierId).Find(&cashier)
		product := models.Product{}
		db.DB.Where("id = ?", v.ProductId).Find(&product)

		RatingsResponse = append(RatingsResponse, &RatingList{
			RatingId:   v.Id,
			CashierID:  v.CashierId,
			ProductID:  v.ProductId,
			ProdRating: float32(v.ProductRating),
			CreatedAt:  v.CreatedAt,
			Cashiers:   cashier,
			Product:    product,
		})

	}

	return c.Status(404).JSON(fiber.Map{
		"success": true,
		"message": "Sucess",
		"data":    RatingsResponse,
		"meta": map[string]interface{}{
			"total": count,
			"limit": limit,
			"skip":  skip,
		},
	})
}
