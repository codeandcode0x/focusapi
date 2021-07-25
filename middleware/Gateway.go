package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

// time out mid
func Gateway(router *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("gateway handler ", c.Request.Header)
		c.Next()
	}
}
