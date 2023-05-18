package models

import (
	"time"
)

//var rating Rating
//var averageRating float64
//
////var ProductRating = product.ProductRating
//
//func getProductRating() float64 {
//	var products []Product
//	for i := 0; i < len(products); i++ {
//		db.DB.Model(&rating).Select("AVG(product_rating)").Where("product_id = ?", products[i].Id).Scan(&averageRating)
//		products[i].ProductRating = averageRating
//	}
//
//	// возвращаем среднее значение ProductRating для всех продуктов
//	var totalRating float64 = 0
//	for i := 0; i < len(products); i++ {
//		totalRating += products[i].ProductRating
//	}
//	return totalRating / float64(len(products))
//}

type Product struct {
	//gorm.Model
	Id               int       `json:"id" gorm:"primaryKey"`
	Sku              string    `json:"sku"`
	Name             string    `json:"name"`
	Stock            int       `json:"stock"`
	Price            int       `json:"price"`
	Image            string    `json:"image"`
	TotalFinalPrice  int       `json:"total_final_price"`
	TotalNormalPrice int       `json:"total_normal_price"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	CategoryId       int       `json:"categoryId"`
	DiscountId       int       `json:"discountId"`
	ProductRating    float64   `json:"productRating"`
	RatingId         int       `json:"ratingId"`
	//CountRating int `json:"countRating"`
}

type ProductResult struct {
	//gorm.Model
	Id            int      `json:"productId" gorm:"primaryKey"`
	Sku           string   `json:"sku"`
	Name          string   `json:"name"`
	Stock         int      `json:"stock"`
	Price         int      `json:"price"`
	Image         string   `json:"image"`
	Category      Category `json:"category"`
	Discount      Discount `json:"discount"`
	Rating        Rating   `json:"rating"`
	ProductRating float64  `json:"productRating"`
}
