package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string `gorm:"type:char(64)"`
}

type Task struct {
	gorm.Model
	Title       string `gorm:"size:255;index" binding:"required"`
	Description string `gorm:"size:2048" binding:"required"`
	Completed   bool   `gorm:"default:false"`
	UserID      int
}
