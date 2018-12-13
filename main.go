package main

import (
	"net/http"
	"html/template"
	"os"
	"io/ioutil"
)

type User struct {
	Username string
}

type Post struct {
	User
	Body string
}

//匿名组合
type IndexViewModel struct {
	Title string
	User
	Posts []Post
}

// PopulateTemplates func
// Create map templates name to templates.Template
func PopulateTemplate() map[string]*template.Template {
	const basePath = "templates"
	result := make(map[string]*template.Template)

	layout := template.Must(template.ParseFiles(basePath + "/_base.html"))
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}
	for _, fi := range fis {
		f, err := os.Open(basePath + "/content/" + fi.Name())
		if err != nil {
			panic("Failed to open template '" + fi.Name() + "'")
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read content from file '" + fi.Name() + "'")
		}
		f.Close()
		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + fi.Name() + "' as template")
		}
		result[fi.Name()] = tmpl
	}
	return result
}

//加入模板引擎
//模板填充
func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		user1 := User{Username: "wade"}
		user2 := User{Username: "bosh"}

		posts := []Post{
			Post{User: user1, Body: "Beautiful day in Portland!"},
			Post{User: user2, Body: "The Avengers movie was so cool!"},
		}
		v := IndexViewModel{Title: "HomePage", User: user1, Posts: posts}

		templates := PopulateTemplate()
		templates["index.html"].Execute(writer, &v)
		//tpl, _ := templates.ParseFiles("templates/index.html")
		//tpl.Execute(writer, &v)
	})
	http.ListenAndServe(":8080", nil)
}
