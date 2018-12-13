package vm


type BaseViewModel struct {
	Title string
}

func (v *BaseViewModel) setTitle(title string){
	v.Title=title
}