package models

import "time"

type Rating struct {
	Id            int       `json:"id" gorm:"primaryKey"`
	CashierId     int       `json:"cashierId"`
	ProductId     int       `json:"productId"`
	ProductRating int       `json:"productRating"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
