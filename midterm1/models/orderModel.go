package models

import (
	"time"
)

type Order struct {
	//gorm.Model
	Id             int       `json:"Id" gorm:"primaryKey"`
	CashierID      int       `json:"cashierId"`
	PaymentTypesId int       `json:"paymentTypesId"`
	TotalPrice     int       `json:"totalPrice"`
	TotalPaid      int       `json:"totalPaid"`
	TotalReturn    int       `json:"totalReturn"`
	ReceiptId      string    `json:"receiptId"`
	IsDownload     int       `json:"is_download"`
	ProductId      string    `json:"productId"`
	Quantities     string    `json:"quantities"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	DeletedAt      time.Time `json:"deletedAt"`
}

type ProductResponseOrder struct {
	//gorm.Model
	ProductId        int      `json:"Id" gorm:"primaryKey"`
	Name             string   `json:"name"`
	Price            int      `json:"price"`
	Qty              int      `json:"qty"`
	Discount         Discount `json:"discount"`
	TotalNormalPrice int      `json:"totalNormalPrice"`
	TotalFinalPrice  int      `json:"totalFinalPrice"`
}

type ProductOrder struct {
	Id         int    `json:"Id" gorm:"primaryKey"`
	Sku        string `json:"sku"`
	Name       string `json:"name"`
	Stock      int    `json:"stock"`
	Price      int    `json:"price"`
	Image      string `json:"image"`
	CategoryId int    `json:"categoryId"`
	DiscountId int    `json:"discountId"`
}

type RevenueResponse struct {
	PaymentTypeId int    `json:"paymentTypeId"`
	Name          string `json:"name"`
	Logo          string `json:"logo"`
	TotalAmount   int    `json:"totalAmount"`
}

type SoldResponse struct {
	ProductId   int    `json:"productId"`
	Name        string `json:"name"`
	TotalQty    int    `json:"totalQty"`
	TotalAmount int    `json:"totalAmount"`
}
