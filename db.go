package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", DB_SQLITE)
	db.LogMode(true)

	if err != nil {
		panic(err)
	}

	if !db.HasTable(&User{}) {
		db.CreateTable(&User{})
	}

	if !db.HasTable(&Task{}) {
		db.CreateTable(&Task{})
	}

	return db
}
