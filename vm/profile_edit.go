package vm

import "github.com/liweiyuan/go-code/model"


// ProfileEditViewModel struct
type ProfileEditViewModel struct {
	LoginViewModel
	ProfileUser model.User
}

// ProfileEditViewModelOp struct
type ProfileEditViewModelOp struct{}

// GetVM func
func (ProfileEditViewModelOp) GetViewModel(username string) ProfileEditViewModel {

	v := ProfileEditViewModel{}
	u, _ := model.GetUserByUserName(username)
	v.SetTitle("Profile Edit")
	v.SetCurrentUser(username)
	v.ProfileUser = *u
	return v
}


//UpdateAboutMe
func UpdateAboutMe(userName, text string) error {
	return model.UpdateAboutMe(userName, text)
}