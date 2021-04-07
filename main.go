package main

import (
	 "github.com/gin-gonic/gin"
	// "github.com/SlyCreator/biko-Ego/controllers"
)

func main()  {
	 r := gin.Default()
	 authRoute := r.Group()
	 {
	 	authRoute.POST("/login",authController.Login)
	 	authRoute.POST("/register",authController.Register)
	 }
	 r.Run(":2020")
}