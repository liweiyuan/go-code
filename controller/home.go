package controller

import (
	"net/http"
	"github.com/liweiyuan/go-code/vm"
)

type home struct {
}

func (h home) registerRouters() {
	http.HandleFunc("/", midleAuth(indexHandler))
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", midleAuth(logoutHandler))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	tpName:="index.html"
	vop := vm.IndexViewModelOp{}
	username, _ :=getSessionUser(r)
	v:=vop.GetViewModel(username)
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
		//fmt.Fprintf(w, "Username:%s Password:%s", userName, password)
		if len(userName) < 3 {
			v.AddError("username must longer than 3")
		}

		if len(password) < 3 {
			v.AddError("password must longer than 6")
		}
		/*if !check(userName, password) {
			v.AddError("username password not correct, please input again")
		}*/

		if !vm.CheckLogin(userName, password) {
			v.AddError("username password not correct, please input again")
		}

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

//账号，密码验证
func check(userName string, password string) bool {

	if userName == "bonfy" && password == "abc123" {
		return true
	}
	return false
}
