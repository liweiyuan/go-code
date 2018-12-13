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
	user1 := model.User{Username: "wade"}
	user2 := model.User{Username: "bosh"}

	posts := []model.Post{
		model.Post{User: user1, Body: "Beautiful day in Portland!"},
		model.Post{User: user2, Body: "The Avengers movie was so cool!"},
	}
	v := IndexViewModel{BaseViewModel{Title: "Homepage"}, user1, posts}
	return v
}
