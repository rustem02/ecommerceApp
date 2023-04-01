package models

import (
	"time"
)

type Cashier struct {
	//gorm.Model
	Id uint `json:"id" gorm:"primarykey"`
	//Name      string    `json:"name"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email" gorm:"unique"`
	Passcode  string    `json:"passcode"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
