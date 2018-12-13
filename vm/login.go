package vm

type LoginViewModel struct {
	BaseViewModel
	Errs []string
}

type LoginViewModelOp struct {
}

func (LoginViewModelOp) GetViewModel() LoginViewModel {
	v := LoginViewModel{}
	v.setTitle("Login")
	return v
}

//AddError func
func (v *LoginViewModel) AddError(errs ...string) {
	v.Errs = append(v.Errs, errs...)
}
