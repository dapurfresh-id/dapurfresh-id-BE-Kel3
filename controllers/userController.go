package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/aldisaputra17/dapur-fresh-id/helpers"
	"github.com/aldisaputra17/dapur-fresh-id/request"
	"github.com/aldisaputra17/dapur-fresh-id/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	UpdateUser(context *gin.Context)
	GetUser(context *gin.Context)
}

type userController struct {
	userService services.UserService
	jwtService  services.JWTService
}

func NewUserController(userService services.UserService, jwtService services.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *userController) UpdateUser(context *gin.Context) {
	var userReqUpdate request.RequestUpdateUser
	errObj := context.ShouldBind(&userReqUpdate)
	if errObj != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errObj.Error(), helpers.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}

	var err1 = context.SaveUploadedFile(userReqUpdate.Image, "assets/"+userReqUpdate.Image.Filename)
	if err1 != nil {
		context.String(http.StatusInternalServerError, "unknown error")
		return
	}

	userReqUpdate.ID = id
	u := c.userService.UpdateUser(userReqUpdate)
	res := helpers.BuildResponse(true, "OK!", u)
	context.JSON(http.StatusOK, res)
}

func (c *userController) GetUser(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user := c.userService.GetUser(id)
	res := helpers.BuildResponse(true, "OK", user)
	context.JSON(http.StatusOK, res)

}
