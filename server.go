package main

import (
	"github.com/gin-gonic/gin"
)

var mySwitch bool = false

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/switch", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": mySwitch,
		})
	})
	router.POST("/switch", func(c *gin.Context) {
		mySwitch = !mySwitch
		c.JSON(200, gin.H{
			"message": mySwitch,
		})
	})
	router.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
