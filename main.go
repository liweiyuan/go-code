package main

import (
	"net/http"
	"html/template"
)

type User struct {
	Username string
}

type IndexViewModel struct {
	Title string
	User  User
}

//加入模板引擎

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		user := User{Username: "wade"}
		v := IndexViewModel{Title: "HomePage", User: user}
		tpl, _ := template.New("").ParseFiles("template/index.html")
		tpl.Execute(writer, &v)
	})
	http.ListenAndServe(":8080", nil)
}
