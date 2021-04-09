package controllers

import (
	"github.com/SlyCreator/biko-Ego/dto"
	"github.com/SlyCreator/biko-Ego/helper"
	"github.com/SlyCreator/biko-Ego/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AuthControlller interface {
	Register(ctx *gin.Context)
}

type authController struct {
	authService services.AuthService
	jwtService  services.JWTService
}

func NewAuthController(authService services.AuthService,jwtService services.JWTService)  AuthControlller{
	return &authController{
		authService: authService,
		jwtService: jwtService,
	}
}

func (c *authController) Register(ctx *gin.Context)  {
	var registerDTO dto.RegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(),helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest , response)
		return
	}
	if !c.authService.IsDuplicateEmail(registerDTO.Email){
		response := helper.BuildErrorResponse("Failed to process request","Duplicate email",helper.EmptyObj{})
		ctx.JSON(http.StatusConflict,response)
		return
	}else {
		createUser := c.authService.CreateUser(registerDTO)
		token := c.jwtService.GenerateToken(strconv.FormatUint(createUser.ID,10))
		createUser.Token = token
		response := helper.BuildResponse(true,"Ok",createUser)
		ctx.JSON(http.StatusCreated,response)
	}
}