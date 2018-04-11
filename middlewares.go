package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")

		if authHeader == "" {
			c.AbortWithStatus(403)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)

		if !(len(parts) != 2 && parts[0] != "Bearer") {
			c.AbortWithStatus(403)
			return
		}

		jwtToken := parts[1]
		token, tokenParseErr := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			return appSecrect, nil
		})

		if tokenParseErr != nil {
			c.AbortWithStatus(403)
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		id := claims["id"].(string)

		c.Set("user_id", id)

		c.Next()
	}
}
