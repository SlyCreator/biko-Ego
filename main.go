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
		authRoute.POST("/register",authController.Register)
		authRoute.POST("/login",authController.Login)
		authRoute.POST("/reset_password",authController.ForgetPassword)
	}

	userRoute := r.Group("api/user")
	{
		userRoute.GET("/",userController.FetchProfile)
		userRoute.PATCH("/",userController.UpdateProfile)
		userRoute.PATCH("/",userController.DeleteAccount)
		userRoute.POST("/next_of_kins",userController.UpdateNextOfKins)
		userRoute.POST("/send_complain",userController.SendReport)

		transactionRoute := r.Group("account")
		{
			transactionRoute.GET("/",userController.FetchAccount)
			transactionRoute.POST("/add_payment_method",userController.AddCard)
			transactionRoute.DELETE("/delete_payment_method",userController.RemoveCard)
			transactionRoute.POST("/fund_acct",userController.FundAccount)
			transactionRoute.POST("/withdraw_fund",userController.WithdrawFromAccount)

		}

		transactionRoute := r.Group("investor")
		{
			transactionRoute.GET("/",userController.FetchAccount)
			transactionRoute.POST("/add_payment_method",userController.AddCard)
			transactionRoute.DELETE("/delete_payment_method",userController.RemoveCard)
			transactionRoute.POST("/fund_acct",userController.FundAccount)
			transactionRoute.POST("/withdraw_fund",userController.WithdrawFromAccount)

		}
	}




	investorRoute := r.Group("api/investor")
	{
		investorRoute.PATCH("/",userController.Login)
	}
	borrowerRoute := r.Group("api/borrower")
	{
		borrowerRoute.PATCH("/",borrowerController.Login)
	}

	 r.Run(":2020")
}



//Route ======> Repository use (Entity) ====> Service(DTO) ====> Controller(helper) ====>