package controller

import (
	"net/http"
	"github.com/liweiyuan/go-code/vm"
)

type home struct {
}

func (h home) registerRouters() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	vop := vm.IndexViewModelOp{}
	v := vop.GetViewModel()
	templates["index.html"].Execute(w, &v)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "login.html"
	vop := vm.LoginViewModelOp{}
	v := vop.GetViewModel()
	templates[tpName].Execute(w, &v)
}
