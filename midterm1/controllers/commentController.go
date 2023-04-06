package controllers

import (
	db "github.com/Krasav4ik01/ecommerceApp/config"
	"github.com/Krasav4ik01/ecommerceApp/middleware"
	"github.com/Krasav4ik01/ecommerceApp/models"
	"github.com/gofiber/fiber/v2"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type Comment struct {
	Comment   models.Comment
	CashierId string `json:"cashierId"`
}

type NewComment struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	CashierId uint   `json:"cashierId"`
	ProductId uint   `json:"productId"`
	Content   string `json:"content"`
}

func CreateComment(c *fiber.Ctx) error {
	//Token authenticate
	headerToken := c.Get("Authorization")
	if headerToken == "" {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized",
			"error":   map[string]interface{}{},
		})
	}
	if err := middleware.AuthenticateToken(middleware.SplitToken(headerToken)); err != nil {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized",
			"error":   map[string]interface{}{},
		})
	}
	//Token authenticate

	type products struct {
		ProductId int `json:"productId"`
	}
	type comments struct {
		CashierId uint   `json:"cashierId"`
		ProductId uint   `json:"productId"`
		Content   string `json:"content"`
		//Cashier   *Cashier  `json:"cashier"`

	}
	body := struct {
		Products []products `json:"products"`
		Comments []comments `json:"comments"`
	}{}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Empty Body",
			"error":   map[string]interface{}{},
		})
	}

	Prodresponse := make([]*models.Comment, 0)

	productsIds := ""
	for _, v := range body.Products {
		productsIds = productsIds + "," + strconv.Itoa(v.ProductId)

		//prods := models.Product{}

		//Prodresponse = append(Prodresponse,
		//	&models.ProductResponseOrder{
		//		ProductId:        prods.Id,
		//		Name:             prods.Name,
		//		Price:            prods.Price,
		//		Discount:         discount,
		//		//Qty:              v.Quantity,
		//		TotalNormalPrice: prods.Price,
		//		TotalFinalPrice:  discCount,
		//	},
		//)

		//Prodresponse = append(Prodresponse,
		//	&models.Comment{
		//		CashierId: 1,
		//		ProductId: uint(prods.Id),
		//		Content:  ,
		//		Cashier:   nil,
		//		CreatedAt: time.Time{},
		//	})

	}
	orderResp := models.Order{
		CashierID: 1,
		//PaymentTypesId: body.PaymentId,
		//TotalPrice:     TotalInvoicePrice.ttprice,
		//TotalPaid:      body.TotalPaid,
		//TotalReturn:    body.TotalPaid - TotalInvoicePrice.ttprice,
		ReceiptId: "R000" + strconv.Itoa(rand.Intn(1000)),
		ProductId: productsIds,
		UpdatedAt: time.Now().UTC(),
		CreatedAt: time.Now().UTC(),
	}

	db.DB.Create(&orderResp)

	return c.Status(200).JSON(fiber.Map{

		"message": "success",
		"success": true,
		"data": map[string]interface{}{
			"order":    orderResp,
			"products": Prodresponse,
		},
	})
}

func CreateAnotherComment(c *fiber.Ctx) error {
	var data NewComment
	err := c.BodyParser(&data)
	if err != nil {
		log.Fatalf("Product error in post request %v", err)
	}
	var p []models.Product
	db.DB.Find(&p)

	comment := models.Comment{
		CashierId: data.CashierId,
		ProductId: data.ProductId,
		Content:   data.Content,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	db.DB.Create(&comment)

	//db.DB.Table("comments").Where("id = ?", comment.Id).Update("sku", "SK00"+strconv.Itoa(comment.Id))

	Response := map[string]interface{}{
		"success": true,
		"message": "Success",
		"data":    comment,
	}
	return (c.JSON(Response))

}

func CommentsList(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit"))
	skip, _ := strconv.Atoi(c.Query("skip"))
	var count int64
	var comment []models.Comment
	db.DB.Select("*").Limit(limit).Offset(skip).Find(&comment).Count(&count)
	metaMap := map[string]interface{}{
		"total": count,
		"limit": limit,
		"skip":  skip,
	}
	cashiersData := map[string]interface{}{
		"comments": comment,
		"meta":     metaMap,
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"Message": "Success",
		"data":    cashiersData,
	})

}
