package main

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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
			c.AbortWithStatusJSON(403, BuildErrorJson(ErrNoAuthHeader))
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)

		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(403, BuildErrorJson(ErrNoBearer))
			return
		}

		jwtToken := parts[1]
		token, tokenParseErr := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			return appSecrect, nil
		})

		if tokenParseErr != nil {
			c.AbortWithStatusJSON(403, BuildErrorJson(tokenParseErr))
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		id := int(claims["id"].(float64))

		c.Set("user_id", id)

		c.Next()
	}
}
