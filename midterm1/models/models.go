package models

import (
	"gorm.io/gorm"
)

// Модель Пользователя
type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Address   string
	Email     string `gorm:"unique"`
	Pass      string
	Role      uint `gorm:"NULL"`
}

type Cart struct {
	gorm.Model
	ProductId uint    // доработать
	Product   Product `gorm:"association_foreignkey:cart_product;"`
	UserId    uint    // доработать
	User      User    `gorm:"association_foreignkey:cart_user;"`
	Quantity  uint
}

// Переделать модель Роли
//type Roles struct {
//	gorm.Model
//	RoleName string
//}

//type Orders struct {
//	gorm.Model
//	CartId      uint
//	OrderStatus string
//	Quantity    int
//	TotalPrice  uint
//}

//type Comments struct {
//	gorm.Model
//	UserId    User    `gorm:"association_foreignkey:comments_user;"`
//	ProductId Product `gorm:"association_foreignkey:comments_product;"`
//	Message   string
//	Rating    int
//}

//type Categodies struct {
//	gorm.Model
//	Title string
//}
