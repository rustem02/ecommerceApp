package models

import (
	"database/sql"

	"github.com/Krasav4ik01/ecommerceApp/config"
	"github.com/Krasav4ik01/ecommerceApp/entities"
)

type UserModel struct {
	db *sql.DB
}

// модель пользователся, который мы возьмем с БД

func NewUserModel() *UserModel {
	conn, err := config.DBConn()

	if err != nil {
		panic(err)
	}

	return &UserModel{
		db: conn,
	}
}

func (u UserModel) Where(user *entities.User, fieldName, fieldValue string) error {

	row, err := u.db.Query("select id, name, username, email, pass from users where "+fieldName+" = ? limit 1", fieldValue)

	if err != nil {
		return err
	}

	defer row.Close()

	for row.Next() {
		row.Scan(&user.Id, &user.Name, &user.Username, &user.Email, &user.Pass)
	}

	return nil
}

//функция для создание/insert пользователя

func (u UserModel) Create(user entities.User) (int64, error) {

	result, err := u.db.Exec("insert into users (name, username, email, pass) values(?,?,?,?)",
		user.Name, user.Username, user.Email, user.Pass)

	if err != nil {
		return 0, err
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId, nil

}
