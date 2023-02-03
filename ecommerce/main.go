package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
)

type User struct {
	//Name  string `json:"name"`
	//Email string `json:"email"`
	//Pass  string `json:"pass"`
	Id                uint16
	Name, Email, Pass string
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
		err = res.Scan(&user.Id, &user.Name, &user.Email, &user.Pass)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
		//fmt.Println(fmt.Sprintf("User: %s with email: %s", user.Name, user.Pass))
	}

	t.ExecuteTemplate(w, "home", users)
}

func login(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/login.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "login", nil)
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
	pass := r.FormValue("pass")

	if name == "" || email == "" || pass == "" {
		fmt.Fprintf(w, "Please fill the fileds")
	} else {

		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/golang")

		if err != nil {
			panic(err)
		}
		defer db.Close()

		insert, err := db.Query(fmt.Sprintf("INSERT INTO `users` (`name`, `email`, `pass`) VALUES('%s', '%s', '%s')", name, email, pass))

		if err != nil {
			panic(err)
		}
		defer insert.Close()

		http.Redirect(w, r, "/home/", http.StatusSeeOther)
	}
}

func handleRequest() {
	http.HandleFunc("/home/", home)
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	http.HandleFunc("/save_data", save_data)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
