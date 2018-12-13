package main

import (
	"github.com/liweiyuan/go-code/controller"
	"net/http"
)

func main() {
	controller.Startup()
	http.ListenAndServe(":8888", nil)
}
