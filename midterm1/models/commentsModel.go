package models

import "time"

type Comment struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	CashierId int    `json:"cashierId"` //Cashier
	ProductId int    `json:"productId"` //Product
	Content   string `json:"content"`
	//Cashier   Cashier  `json:"cashier"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updated_at"`
}
