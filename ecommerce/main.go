package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      uint16 `json:"age"`
	IsActive bool   `json:"is_active"`
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/golang")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	//insert, err := db.Query("INSERT INTO `users` (`name`, `age`, `email`, `is_active`) VALUES('Robin', 20, 'robin@mail.ru', '1')")
	//
	//if err != nil {
	//	panic(err)
	//}
	//defer insert.Close()

	res, err := db.Query("SELECT `name`, `age`, `email`, `is_active` FROM `users`")
	if err != nil {
		panic(err)
	}

	for res.Next() {
		var user User
		err = res.Scan(&user.Name, &user.Age, &user.Email, &user.IsActive)
		if err != nil {
			panic(err)
		}
		fmt.Println(fmt.Sprintf("User: %s with age: %d , email: %s , is_active: %t", user.Name, user.Age, user.Email, user.IsActive))
	}

}
