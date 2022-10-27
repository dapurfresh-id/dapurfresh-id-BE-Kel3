package controllers

import (
	"net/http"

	"github.com/aldisaputra17/dapur-fresh-id/entities"
	"github.com/aldisaputra17/dapur-fresh-id/helpers"
	"github.com/aldisaputra17/dapur-fresh-id/request"
	"github.com/aldisaputra17/dapur-fresh-id/services"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type authController struct {
	authService services.AuthService
	jwtService  services.JWTService
}

func NewAuthController(authService services.AuthService, jwtService services.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var reqLogin request.RequestLogin
	err := ctx.ShouldBind(&reqLogin)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(reqLogin.Username, reqLogin.Password)
	if v, ok := authResult.(entities.User); ok {
		generatedToken := c.jwtService.GenerateToken(v.ID)
		v.Token = generatedToken
		response := helpers.BuildResponse(true, "Ok!", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helpers.BuildErrorResponse("Please check again your credential", "Invalid Credential", helpers.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Register(ctx *gin.Context) {
	var reqRegister request.RequestRegister
	errObj := ctx.ShouldBind(&reqRegister)
	if errObj != nil {
		response := helpers.BuildErrorResponse("Failed to process request", errObj.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateUsername(reqRegister.Username) {
		response := helpers.BuildErrorResponse("Failed to process request", "Duplicate username", helpers.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
		return
	}
	createdUser, err := c.authService.CreateUser(ctx, reqRegister)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to created", errObj.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else {
		token := c.jwtService.GenerateToken(createdUser.ID)
		createdUser.Token = token
		response := helpers.BuildResponse(true, "Created!", createdUser)
		ctx.JSON(http.StatusCreated, response)
	}
}
