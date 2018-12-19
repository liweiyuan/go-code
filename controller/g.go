package controller

import (
	"html/template"
	"github.com/gorilla/sessions"
)

var (
	homeController home
	templates      map[string]*template.Template
	sessionName    string
	store          *sessions.CookieStore
)

func init() {
	templates = PopulateTemplate()
	store = sessions.NewCookieStore([]byte("something-very-secret"))
	sessionName = "go-code"
}

//Startup func
func Startup() {
	homeController.registerRouters()
}
