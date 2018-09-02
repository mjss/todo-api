package main

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Name      string
	Email     string `gorm:"type:varchar(100);unique_index"`
	Password  string `gorm:"type:char(64)"`
}

type Task struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	Title       string     `gorm:"size:255;index" json:"title" binding:"required"`
	Description string     `gorm:"size:2048" json:"description" binding:"required"`
	Completed   bool       `gorm:"default:false" json:"completed"`
	UserID      int        `json:"-"`
}
