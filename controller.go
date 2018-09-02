package main

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// curl -i -X POST -H "Content-Type: application/json" -d "{ \"email\": \"sunminjie91@gmail.com\", \"password\": \"12345678\", \"name\": \"Sun Minjie\" }" http://localhost:8080/api/v1/users
func CreateUser(c *gin.Context) {
	db := InitDB()

	defer db.Close()

	var user User
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
	var user User
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

	cred := &User{}
	var user User
	c.Bind(cred)
	db.Where("Email = ?", cred.Email).First(&user)

	ok := verifyPassword(user.Password, cred.Password)
	if !ok {
		c.AbortWithStatus(401)
		return
	}

	expire := time.Now().Add(time.Hour * 720)
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

func CreateTask(c *gin.Context) {
	db := InitDB()

	defer db.Close()

	var task Task
	bindErr := c.Bind(&task)

	titleErr := validateTitle(task.Title)
	descriptionErr := validateDescription(task.Description)

	if bindErr != nil {
		c.AbortWithStatusJSON(400, BuildErrorJson(bindErr))
		return
	}

	if titleErr != nil {
		c.AbortWithStatusJSON(400, BuildErrorJson(titleErr))
		return
	}

	if descriptionErr != nil {
		c.AbortWithStatusJSON(400, BuildErrorJson(descriptionErr))
		return
	}

	task.UserID = c.MustGet("user_id").(int)

	cursor := db.Create(&task)
	dbErr := cursor.Error

	if dbErr != nil {
		c.AbortWithStatusJSON(500, BuildErrorJson(dbErr))
		return
	}

	row := cursor.Value
	c.JSON(201, row)
}

func GetTasks(c *gin.Context) {
	db := InitDB()

	defer db.Close()

	userID := c.MustGet("user_id").(int)

	var tasks []Task
	db.Where("user_id = ?", userID).Find(&tasks)

	c.JSON(200, tasks)
}

func UpdateTask(c *gin.Context) {
	db := InitDB()
	defer db.Close()

	var task Task
	taskID := c.Params.ByName("id")
	userID := c.MustGet("user_id").(int)

	if err := db.Where("id = ? AND user_id = ?", taskID, userID).First(&task).Error; err != nil {
		c.AbortWithStatusJSON(404, BuildErrorJson(ErrTaskNotFound))
		return
	}

	c.BindJSON(&task)
	db.Save(&task)
	c.JSON(200, task)
}

func DeleteTask(c *gin.Context) {
	db := InitDB()
	defer db.Close()

	var task Task
	taskID := c.Params.ByName("id")
	userID := c.MustGet("user_id").(int)

	if err := db.Where("id = ? AND user_id = ?", taskID, userID).First(&task).Error; err != nil {
		c.AbortWithStatusJSON(404, BuildErrorJson(ErrTaskNotFound))
		return
	}

	db.Delete(&task)
	c.JSON(200, gin.H{
		taskID: "ok",
	})
}
