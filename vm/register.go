package vm

import (
	"github.com/liweiyuan/go-code/model"
	"log"
)

type RegisterViewModel struct {
	LoginViewModel
}

//RegisterViewModelOp struct
type RegisterViewModelOp struct {
}

//GetViewModel func
func (RegisterViewModelOp) GetViewModel() RegisterViewModel {
	v := RegisterViewModel{}
	v.SetTitle("Register")
	return v
}

//AddUser func
func AddUser(username, password, email string) error {
	return model.AddUser(username, password, email)
}

//CheckUserExit func
func CheckUserExist(username string) bool{
	_,err:=model.GetUserByUserName(username)
	if err!=nil{
		log.Println("Can not find username: ", username)
		return true
	}
	return false
}
