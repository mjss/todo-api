package main

import (
	"github.com/gin-gonic/gin"
)

func cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		context.Next()
	}
}
