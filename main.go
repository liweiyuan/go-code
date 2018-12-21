package main

import (
	"github.com/liweiyuan/go-code/model"
	"github.com/liweiyuan/go-code/controller"
	"net/http"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/gorilla/context"
)

func main() {

	// Setup DB
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	// Setup Controller
	controller.Startup()
	http.ListenAndServe(":8888", context.ClearHandler(http.DefaultServeMux))
}
