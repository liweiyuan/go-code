package model

type User struct {
	ID           int     `gorm:"primary_key"`
	Username     string  `gorm:"varchar(64)"`
	Email        string  `gorm:"varchar(12)"`
	PasswordHash string  `gorm:"varchar(128)"`
	Posts        []Post
	Followers    []*User `gorm:"many2many:follower;association_jointable_foreignkey:follower_id"`
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

//AddUser func
func AddUser(username, password, email string) error {
	user := User{Username: username, Email: email}
	user.SetPassword(password)
	return db.Create(&user).Error
}
