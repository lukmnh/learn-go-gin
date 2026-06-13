package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Entry point of gin framework")
	/*
		gin.Default() membuat object router untuk menerima HTTP Request.
	*/
	var router *gin.Engine = gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "TODOS RestAPI",
			"status":  "200",
		})
	})

	router.Run(":9000")
}
