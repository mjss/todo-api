package main

import (
	"github.com/gin-gonic/gin"
)

var appSecrect = []byte("secrect")

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(Cors())

	// Ping test
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	v1 := router.Group("api/v1")
	{
		v1.POST("/users", CreateUser)
		v1.POST("/login", LoginUser)

		private := v1.Group("")
		private.Use(JwtAuth())
		private.GET("/users/:id", GetUser)
	}

	return router
}

func main() {
	router := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	router.Run(":8080")
}
