package model

import (
	"time"
	"fmt"
)

type User struct {
	ID           int     `gorm:"primary_key"`
	Username     string  `gorm:"varchar(64)"`
	Email        string  `gorm:"varchar(12)"`
	PasswordHash string  `gorm:"varchar(128)"`
	LastSeen     *time.Time
	AboutMe      string  `gorm:"type:varchar(140)"`
	Avatar       string  `gorm:"type:varchar(200)"`
	Posts        []Post
	Followers    []*User `gorm:"many2many:follower;association_jointable_foreignkey:follower_id"`
}

//SetAvatar func
func (u *User) SetAvatar(email string) {
	u.Avatar = fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon", Md5(email))
}

// SetPasswordHash func 
func (u *User) SetPassword(password string) {
	u.PasswordHash = GeneratePasswordHash(password)
}

// CheckPassword func
func (u *User) CheckPassword(password string) bool {
	return GeneratePasswordHash(password) == u.PasswordHash
}

// GetUserByUserName func
func GetUserByUserName(userName string) (*User, error) {
	var user User
	if err := db.Where("username=?", userName).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

//UpdateUserLastSeen
func UpdateLastSeen(userName string) error {
	contents := map[string]interface{}{"last_seen": time.Now()}
	return UpdateUserByUsername(userName, contents)
}

// UpdateUserByUsername func
func UpdateUserByUsername(username string, contents map[string]interface{}) error {
	item, err := GetUserByUserName(username)
	if err != nil {
		return err
	}
	return db.Model(item).Updates(contents).Error
}

//UpdateAboutMe
func UpdateAboutMe(userName, text string) error {
	content := map[string]interface{}{"about_me": text}
	return UpdateUserByUsername(userName, content)
}

//AddUser func
func AddUser(username, password, email string) error {
	user := User{Username: username, Email: email}
	user.SetPassword(password)
	user.SetAvatar(email)
	return db.Create(&user).Error
}
