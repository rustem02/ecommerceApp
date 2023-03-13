package models

import (
	"google.golang.org/protobuf/types/known/wrapperspb"
	"gorm.io/gorm"
	"image"
)

// Модель Пользователя
type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Address   string
	Email     string `gorm:"unique"`
	Pass      string
	//Role      int
}

type Product struct {
	gorm.Model
	Title        string
	Desc         string
	UserId       uint16 // доработать
	User         User   `gorm:"association_foreignkey:Refer"`
	CategoriesId int
	Price        uint
	Rating       wrapperspb.DoubleValue // доработать
	Count        uint
	Image        image.Image // доработать

}

type Cart struct {
	gorm.Model
	ProductId uint    // доработать
	Product   Product `gorm:"association_foreignkey:Refer"`
	UserId    uint    // доработать
	User      User    `gorm:"association_foreignkey:Refer"`
	Quantity  uint
}
