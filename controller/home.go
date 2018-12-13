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

	//根据请求类型处理不同的请求
	if r.Method==http.MethodGet{
		templates[tpName].Execute(w, &v)
	}
	if r.Method==http.MethodPost{
		r.ParseForm()
		userName:=r.Form.Get("username")
		password:=r.Form.Get("password")
		//fmt.Fprintf(w, "Username:%s Password:%s", userName, password)
		if len(userName) < 3 {
			v.AddError("username must longer than 3")
		}

		if len(password) < 3 {
			v.AddError("password must longer than 6")
		}
		if !check(userName,password){
			v.AddError("username password not correct, please input again")
		}
		if len(v.Errs)>0{
			templates[tpName].Execute(w,&v)
		}else {
			http.Redirect(w,r,"/",http.StatusSeeOther)
		}
	}

}

//账号，密码验证
func check(userName string, password string) bool {

	if userName == "wade" && password == "bosh" {
		return true
	}
	return false
}
