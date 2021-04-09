package main

import (
	"github.com/SlyCreator/biko-Ego/config"
	"github.com/gin-gonic/gin"
	"github.com/SlyCreator/biko-Ego/controllers"
	"github.com/SlyCreator/biko-Ego/repository"
	"github.com/SlyCreator/biko-Ego/services"
	"gorm.io/gorm"
)

var (
	db			*gorm.DB							= config.OpenDatabaseConnection()
	userRepository  repository.UserRepository 		= repository.NewUserRepository(db)
	authService		services.AuthService		  	= services.NewAuthService(userRepository)
	jwtService		services.JWTService				= services.NewJWTService()
	authController controllers.AuthControlller 		= controllers.NewAuthController(authService, jwtService)
)



func main()  {
	 r := gin.Default()
	authRoute := r.Group("api/user")
	{
		//authRoute.POST("/login",authController.Login)
		authRoute.POST("/register",authController.Register)
		//authRoute.POST("/reset_password",authController.ForgetPassword)
	}
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



//Repository ====> Service ====> Controller ====> Route