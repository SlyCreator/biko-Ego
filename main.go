package main

import (
	 "github.com/gin-gonic/gin"
	"gorm.io/gorm"

	// "github.com/SlyCreator/biko-Ego/controllers"
)




func main()  {
	 r := gin.Default()
	// authRoute := r.Group("api/auth")
	// {
	// 	authRoute.POST("/login",authController.Login)
	// 	authRoute.POST("/register",authController.Register)
	// 	authRoute.POST("/reset_password",authController.ForgetPassword)
	// }
	//userRoute := r.Group("api/user")
	//{
	//	userRoute.PATCH("/",userController.Login)
	//}
	//investorRoute := r.Group("api/investor")
	//{
	//	investorRoute.PATCH("/",userController.Login)
	//}
	//borrowerRoute := r.Group("api/borrower")
	//{
	//	borrowerRoute.PATCH("/",borrowerController.Login)
	//}

	 r.Run(":2020")
}