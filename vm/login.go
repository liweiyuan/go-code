package vm

type LoginViewModel struct {
	BaseViewModel
}

type LoginViewModelOp struct {
}

func (LoginViewModelOp) GetViewModel() LoginViewModel {
	v := LoginViewModel{}
	v.setTitle("Login")
	return v
}
