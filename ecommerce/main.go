package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"

	authcontroller "github.com/jeypc/go-auth/controllers"
)

type User struct {
	Id                          uint16
	Name, Username, Email, Pass string
}

var users = []User{}

func home(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/homePage.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/golang")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	res, err := db.Query("SELECT * FROM `users`")

	if err != nil {
		panic(err)
	}

	for res.Next() {
		var user User
		err = res.Scan(&user.Id, &user.Name, &user.Email, &user.Username, &user.Pass)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
		//fmt.Println(fmt.Sprintf("User: %s with email: %s", user.Name, user.Pass))
	}

	t.ExecuteTemplate(w, "home", users)
}

func login(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/login1.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "login", nil)
}

// TODO: Пока функция authorization не работает. Надо доработать

func authorization(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	pass := r.FormValue("pass")

	if email == "" || pass == "" {
		fmt.Fprintf(w, "Please fill all fileds")
	} else {

		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/golang")

		if err != nil {
			panic(err)
		}
		defer db.Close()

		res, err := db.Query("SELECT * FROM `users`")

		if err != nil {
			panic(err)
		}

		for res.Next() {
			var user User
			err = res.Scan(&user.Id, &user.Name, &user.Email, &user.Pass)
			if err != nil {
				panic(err)
			}
			users = append(users, user)
			if email != user.Email || pass != user.Pass {
				fmt.Fprintf(w, "You are not registered ")

			} else if email != user.Email && pass != user.Pass {
				fmt.Fprintf(w, "You are not registered ")
			} else {
				http.Redirect(w, r, "/home/", http.StatusSeeOther)
			}

		}
		http.Redirect(w, r, "/home/", http.StatusSeeOther)

	}
}

func register(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/register.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "register", nil)
}

func save_data(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	username := r.FormValue("username")
	pass := r.FormValue("pass")

	if name == "" || email == "" || pass == "" {
		fmt.Fprintf(w, "Please fill all fileds")
	} else {

		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/golang")

		if err != nil {
			panic(err)
		}
		defer db.Close()

		insert, err := db.Query(fmt.Sprintf("INSERT INTO `users` (`name`, `email`, `username`, `pass`) VALUES('%s', '%s', '%s', '%s')", name, username, email, pass))

		if err != nil {
			panic(err)
		}
		defer insert.Close()

		http.Redirect(w, r, "/home/", http.StatusSeeOther)
	}
}

func handleRequest() {
	http.HandleFunc("/home/", home)
	//http.HandleFunc("/login", login)
	http.HandleFunc("/authorization", authorization)
	http.HandleFunc("/register", register)
	http.HandleFunc("/save_data", save_data)

	//новые функции, пока на доработке
	http.HandleFunc("/", authcontroller.Index)
	http.HandleFunc("/login", authcontroller.Login)
	http.HandleFunc("/logout", authcontroller.Logout)

	http.ListenAndServe(":8080", nil)

}

//TODO:Testing new functions

func main() {
	handleRequest()
}
