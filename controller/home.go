package controller

import (
	"net/http"
	"github.com/liweiyuan/go-code/vm"
	"log"
)

type home struct {
}

func (h home) registerRouters() {
	http.HandleFunc("/", midleAuth(indexHandler))
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/logout", midleAuth(logoutHandler))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	tpName := "index.html"
	vop := vm.IndexViewModelOp{}
	username, _ := getSessionUser(r)
	v := vop.GetViewModel(username)
	templates[tpName].Execute(w, &v)

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "login.html"
	vop := vm.LoginViewModelOp{}
	v := vop.GetViewModel()

	//根据请求类型处理不同的请求
	if r.Method == http.MethodGet {
		templates[tpName].Execute(w, &v)
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		userName := r.Form.Get("username")
		password := r.Form.Get("password")

		errs := checkLogin(userName, password)
		v.AddError(errs...)

		if len(v.Errs) > 0 {
			templates[tpName].Execute(w, &v)
		} else {
			setSessionUser(w, r, userName)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}

}

//logout Func
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	clearSession(w, r)
	http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
}


//registerHandler Func
func registerHandler(w http.ResponseWriter, r *http.Request) {
	tpName:="register.html"

	vop:=vm.RegisterViewModelOp{}
	v:=vop.GetViewModel()

	if r.Method == http.MethodGet {
		templates[tpName].Execute(w, &v)
	}

	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.Form.Get("username")
		email := r.Form.Get("email")
		pwd1 := r.Form.Get("pwd1")
		pwd2 := r.Form.Get("pwd2")

		errs := checkRegister(username, email, pwd1, pwd2)
		v.AddError(errs...)

		if len(v.Errs) > 0 {
			templates[tpName].Execute(w, &v)
		} else {
			if err := addUser(username, pwd1, email); err != nil {
				log.Println("add User error:", err)
				w.Write([]byte("Error insert database"))
				return
			}
			setSessionUser(w, r, username)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}
