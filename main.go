package main

import (
	 "github.com/gin-gonic/gin"
	// "github.com/SlyCreator/biko-Ego/controllers"
)

func main()  {
	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.Run(":2020")
}