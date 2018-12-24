package vm

import "github.com/liweiyuan/go-code/model"

type IndexViewModel struct {
	BaseViewModel
	Posts []model.Post
}

type IndexViewModelOp struct {
}

func (IndexViewModelOp) GetViewModel(username string) IndexViewModel {
	u1, _ := model.GetUserByUserName(username)
	posts, _ := model.GetPostsByUserID(u1.ID)
	v := IndexViewModel{BaseViewModel{Title: "Homepage"}, *posts}
	v.SetCurrentUser(username)
	return v
}
