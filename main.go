package main

import (
	 "github.com/gin-gonic/gin"
	// "github.com/SlyCreator/biko-Ego/controllers"
)

func main()  {
	 r := gin.Default()

	 r.Run(":2020")
}