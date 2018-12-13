package main

import (
	"net/http"
	"html/template"
)

type User struct {
	Username string
}

type Post struct {
	User User
	Body string
}

type IndexViewModel struct {
	Title string
	User  User

	Posts []Post
}

//加入模板引擎

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		user1 := User{Username: "wade"}
		user2 := User{Username: "bosh"}

		posts := []Post{
			Post{User: user1, Body: "Beautiful day in Portland!"},
			Post{User: user2, Body: "The Avengers movie was so cool!"},
		}
		v := IndexViewModel{Title: "HomePage", User: user1, Posts: posts}
		tpl, _ := template.ParseFiles("template/index.html")
		tpl.Execute(writer, &v)
	})
	http.ListenAndServe(":8080", nil)
}
