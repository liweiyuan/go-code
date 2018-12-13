package model

import "time"

type Post struct {
	ID        int        `gorm:"primary_key"`
	UserID    int
	User      User
	Body      string     `gorm:"varchar(180)"`
	Timestamp *time.Time `sql:"DEFAULT:current_timestamp"`
}
