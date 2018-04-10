package main

import (
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	db := InitDB()

	defer db.close()

	var user Users
	c.bind(&User)

	if !isValidEmail(user.Email) {
		c.JSON(422, gin.H{"error": "invalid email"})
	}
}
