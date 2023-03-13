package models

import (
	"database/sql"

	"github.com/Krasav4ik01/ecommerceApp/config"
	"github.com/Krasav4ik01/ecommerceApp/entities"
)

type ProductModel struct {
	db *sql.DB
}

// модель продукта, которую мы возьмем с БД

func NewProductModel() *ProductModel {
	conn, err := config.DBConn()

	if err != nil {
		panic(err)
	}

	return &ProductModel{
		db: conn,
	}
}

func (u ProductModel) WhereProduct(product *entities.Product, fieldName, fieldValue string) error {

	row, err := u.db.Query("select id, title, model, price, rating from products where "+fieldName+" = ? limit 1", fieldValue)

	if err != nil {
		return err
	}

	defer row.Close()

	for row.Next() {
		row.Scan(&product.Id, &product.Title, &product.Model, &product.Price, &product.Rating)
	}

	return nil
}

//функция для создание/insert продукта

func (u ProductModel) CreateProduct(product entities.Product) (int64, error) {

	result, err := u.db.Exec("insert into products (title, model, price, rating) values(?,?,?,?)",
		product.Title, product.Model, product.Price, product.Rating)

	if err != nil {
		return 0, err
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId, nil

}
