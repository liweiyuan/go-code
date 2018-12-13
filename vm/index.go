package vm

import "github.com/liweiyuan/go-code/model"

type IndexViewModel struct {
	BaseViewModel
	model.User
	Posts []model.Post
}

type IndexViewModelOp struct {
}

func (IndexViewModelOp) GetViewModel() IndexViewModel {
	u1, _ := model.GetUserByUserName("rene")
	posts, _ := model.GetPostsByUserID(u1.ID)
	v := IndexViewModel{BaseViewModel{Title: "Homepage"}, *u1, *posts}
	return v
}
