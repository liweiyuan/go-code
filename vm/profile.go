package vm

import "github.com/liweiyuan/go-code/model"

/**
个人主页
 */

// ProfileViewModel struct
type ProfileViewModel struct {
	BaseViewModel
	Posts       []model.Post
	Editable    bool
	ProfileUser model.User
}

// ProfileViewModelOp struct
type ProfileViewModelOp struct {
}

func (ProfileViewModelOp) GetViewModel(sUser, pUser string) (ProfileViewModel, error) {

	v := ProfileViewModel{}
	v.SetTitle("Profile")
	user, err := model.GetUserByUserName(pUser)
	if err != nil {
		return v, err
	}
	posts, _ := model.GetPostsByUserID(user.ID)
	v.ProfileUser = *user
	v.Editable = sUser == pUser
	v.Posts = *posts
	v.SetCurrentUser(sUser)

	return v, nil
}


