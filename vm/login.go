package vm

import (
	"github.com/liweiyuan/go-code/model"
	"log"
)

type LoginViewModel struct {
	BaseViewModel
	Errs []string
}

type LoginViewModelOp struct {
}

func (LoginViewModelOp) GetViewModel() LoginViewModel {
	v := LoginViewModel{}
	v.SetTitle("Login")
	return v
}

//AddError func
func (v *LoginViewModel) AddError(errs ...string) {
	v.Errs = append(v.Errs, errs...)
}

// CheckLogin func

func CheckLogin(username, password string) bool {
	user, err := model.GetUserByUserName(username)
	if err != nil {
		log.Println("Can not find username: ", username)
		log.Println("Error:", err)
		return false
	}
	return user.CheckPassword(password)
}
