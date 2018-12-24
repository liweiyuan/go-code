package vm


type BaseViewModel struct {
	Title string
	CurrentUser string
}

func (v *BaseViewModel) SetTitle(title string){
	v.Title=title
}

//setCurrentUser func
func (v *BaseViewModel) SetCurrentUser(username string)  {
	v.CurrentUser=username
}