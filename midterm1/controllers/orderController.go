package controllers

import (
	db "github.com/Krasav4ik01/ecommerceApp/config"
	"github.com/Krasav4ik01/ecommerceApp/middleware"
	"github.com/Krasav4ik01/ecommerceApp/models"
	"github.com/gofiber/fiber/v2"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func CreateOrder(c *fiber.Ctx) error {
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
		Quantity  int `json:"qty"`
	}

	body := struct {
		CashierId int        `json:"cashierId"`
		PaymentId int        `json:"paymentId"`
		TotalPaid int        `json:"totalPaid"`
		Products  []products `json:"products"`
	}{}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Empty Body",
			"error":   map[string]interface{}{},
		})
	}

	Prodresponse := make([]*models.ProductResponseOrder, 0)

	var TotalInvoicePrice = struct {
		ttprice int
	}{}

	productsIds := ""
	quantities := ""
	for _, v := range body.Products {
		totalPrice := 0
		productsIds = productsIds + "," + strconv.Itoa(v.ProductId)
		quantities = quantities + "," + strconv.Itoa(v.Quantity)

		prods := models.ProductOrder{}
		var discount models.Discount
		db.DB.Table("products").Where("id=?", v.ProductId).First(&prods)
		db.DB.Where("id = ?", prods.DiscountId).Find(&discount)
		discCount := 0

		if discount.Type == "BUY_N" {
			totalPrice = prods.Price * v.Quantity

			discCount = totalPrice - discount.Result
			TotalInvoicePrice.ttprice = TotalInvoicePrice.ttprice + discCount

		}

		if discount.Type == "PERCENT" {
			totalPrice = prods.Price * v.Quantity
			percentage := totalPrice * discount.Result / 100
			discCount = totalPrice - percentage
			TotalInvoicePrice.ttprice = TotalInvoicePrice.ttprice + discCount
		}

		Prodresponse = append(Prodresponse,
			&models.ProductResponseOrder{
				ProductId:        prods.Id,
				Name:             prods.Name,
				Price:            prods.Price,
				Discount:         discount,
				Qty:              v.Quantity,
				TotalNormalPrice: prods.Price,
				TotalFinalPrice:  discCount,
			},
		)

	}
	orderResp := models.Order{
		CashierID:      body.CashierId,
		PaymentTypesId: body.PaymentId,
		TotalPrice:     TotalInvoicePrice.ttprice,
		TotalPaid:      body.TotalPaid,
		TotalReturn:    body.TotalPaid - TotalInvoicePrice.ttprice,
		ReceiptId:      "R000" + strconv.Itoa(rand.Intn(1000)),
		ProductId:      productsIds,
		Quantities:     quantities,
		UpdatedAt:      time.Now().UTC(),
		CreatedAt:      time.Now().UTC(),
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

func OrderDetail(c *fiber.Ctx) error {
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

	param := c.Params("orderId")

	var order models.Order
	db.DB.Where("id=?", param).First(&order)

	if order.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"Message": "Not Found",
			"error":   map[string]interface{}{},
		})
	}
	productIds := strings.Split(order.ProductId, ",")
	TotalProducts := make([]*models.Product, 0)

	for i := 1; i < len(productIds); i++ {
		prods := models.Product{}
		db.DB.Where("id = ?", productIds[i]).Find(&prods)
		TotalProducts = append(TotalProducts, &prods)
	}
	cashier := models.Cashier{}
	db.DB.Where("id = ?", order.CashierID).Find(&cashier)

	paymentType := models.PaymentType{}
	db.DB.Where("id = ?", order.PaymentTypesId).Find(&paymentType)

	orderTable := models.Order{}
	db.DB.Where("id = ?", order.Id).Find(&orderTable)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"data": map[string]interface{}{
			"order": map[string]interface{}{
				"orderId":        order.Id,
				"cashiersId":     order.CashierID,
				"paymentTypesId": order.PaymentTypesId,
				"totalPrice":     order.TotalPrice,
				"totalPaid":      order.TotalPaid,
				"totalReturn":    order.TotalReturn,
				"receiptId":      order.ReceiptId,
				"createdAt":      order.CreatedAt,
				"cashier":        cashier,
				"payment_type":   paymentType,
			},
			"products": TotalProducts,
		},
		"Message": "Success",
	})

}

func OrdersList(c *fiber.Ctx) error {

	limit, _ := strconv.Atoi(c.Query("limit"))
	skip, _ := strconv.Atoi(c.Query("skip"))
	var count int64
	var order []models.Order

	db.DB.Select("*").Limit(limit).Offset(skip).Find(&order).Count(&count)

	type OrderList struct {
		OrderId        int                `json:"orderId"`
		CashierID      int                `json:"cashiersId"`
		PaymentTypesId int                `json:"paymentTypesId"`
		TotalPrice     int                `json:"totalPrice"`
		TotalPaid      int                `json:"totalPaid"`
		TotalReturn    int                `json:"totalReturn"`
		ReceiptId      string             `json:"receiptId"`
		CreatedAt      time.Time          `json:"createdAt"`
		Payments       models.PaymentType `json:"payment_type"`
		Cashiers       models.Cashier     `json:"cashier"`
	}
	OrderResponse := make([]*OrderList, 0)

	for _, v := range order {
		cashier := models.Cashier{}
		db.DB.Where("id = ?", v.CashierID).Find(&cashier)
		paymentType := models.PaymentType{}
		db.DB.Where("id = ?", v.PaymentTypesId).Find(&paymentType)

		OrderResponse = append(OrderResponse, &OrderList{
			OrderId:        v.Id,
			CashierID:      v.CashierID,
			PaymentTypesId: v.PaymentTypesId,
			TotalPrice:     v.TotalPrice,
			TotalPaid:      v.TotalPaid,
			TotalReturn:    v.TotalReturn,
			ReceiptId:      v.ReceiptId,
			CreatedAt:      v.CreatedAt,
			Payments:       paymentType,
			Cashiers:       cashier,
		})

	}

	return c.Status(404).JSON(fiber.Map{
		"success": true,
		"message": "Sucess",
		"data":    OrderResponse,
		"meta": map[string]interface{}{
			"total": count,
			"limit": limit,
			"skip":  skip,
		},
	})
}
