package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

// curl -i -X POST -H "Content-Type: application/json" -d "{ \"email\": \"sunminjie91@gmail.com\", \"password\": \"12345678\", \"name\": \"Sun Minjie\" }" http://localhost:8080/api/v1/users
func CreateUser(c *gin.Context) {
	db := InitDB()

	defer db.Close()

	var user Users
	c.Bind(&user)

	emailErr := validateEmail(user.Email)
	passwordErr := validatePassword(user.Password)
	nameErr := validateName(user.Name)

	if emailErr != nil {
		c.AbortWithStatusJSON(400, BuildErrorJson(emailErr))
		return
	}

	if passwordErr != nil {
		c.AbortWithStatusJSON(400, BuildErrorJson(passwordErr))
		return
	}

	if nameErr != nil {
		c.AbortWithStatusJSON(400, BuildErrorJson(nameErr))
		return
	}

	user.Password = hashPassword(user.Password)
	db.Create(&user)
	c.JSON(201, BuildUserJson(user))
}

func GetUser(c *gin.Context) {
	db := InitDB()

	defer db.Close()

	id := c.Params.ByName("id")
	var user Users
	db.Where("ID = ?", id).First(&user)

	if user.ID != 0 {
		c.JSON(200, BuildUserJson(user))
	} else {
		c.AbortWithStatusJSON(404, BuildErrorJson(ErrUserNotFound))
	}

}

// curl -i -X POST -H "Content-Type: application/json" -d "{ \"email\": \"sunminjie91+1@gmail.com\", \"password\": \"12345678\" }" http://localhost:8080/api/v1/login
func LoginUser(c *gin.Context) {
	db := InitDB()

	defer db.Close()

	cred := &Users{}
	var user Users
	c.Bind(cred)
	db.Where("Email = ?", cred.Email).First(&user)

	ok := verifyPassword(user.Password, cred.Password)
	if !ok {
		c.AbortWithStatus(401)
		return
	}

	expire := time.Now().Add(time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": expire.Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := token.SignedString(appSecrect)

	if err != nil {
		c.AbortWithError(401, ErrCreateToken)
		return
	}

	c.JSON(200, gin.H{
		"token":  tokenString,
		"expire": expire.Unix(),
	})
}
