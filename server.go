package main

import (
	"github.com/gin-gonic/gin"
)

var mySwitch bool = false

func main() {
	router := gin.Default()
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
