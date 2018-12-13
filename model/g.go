package model

import (
	"github.com/jinzhu/gorm"
	"log"
	"github.com/liweiyuan/go-code/config"
)

var db *gorm.DB


func SetDB(database *gorm.DB) {
	db = database
}

// ConnectToDB func

func ConnectToDB() *gorm.DB {
	connectingStr := config.GetMysqlConnectingString()
	log.Println("Connet to db...")
	db, err := gorm.Open("mysql", connectingStr)
	if err != nil {
		panic("Failed to connect database")
	}
	db.SingularTable(true)
	return db
}
