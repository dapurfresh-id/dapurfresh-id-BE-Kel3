package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aldisaputra17/dapur-fresh-id/helpers"
	"github.com/aldisaputra17/dapur-fresh-id/request"
	"github.com/aldisaputra17/dapur-fresh-id/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController interface {
	Update(ctx *gin.Context)
	GetUser(context *gin.Context)
}

type userController struct {
	userService services.UserService
	jwtService  services.JWTService
}

func NewUserController(userServ services.UserService, jwtServ services.JWTService) UserController {
	return &userController{
		userService: userServ,
		jwtService:  jwtServ,
	}
}

func (c *userController) Update(ctx *gin.Context) {
	var reqUser *request.RequestUserUpdate
	errObj := ctx.ShouldBindJSON(&reqUser)
	if errObj != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errObj.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	fileHeader, err := ctx.FormFile("image")
	log.Println(fileHeader)
	if err != nil {
		res := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	reqUser.ID = uuid.Must(uuid.Parse(id))
	reqUser.Image = fileHeader
	result, img, err := c.userService.Update(ctx, reqUser)
	if err != nil {
		res := helpers.BuildErrorResponse("Failed to update user", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	response := helpers.BuildSuccessUpdate(true, "Updated!", result, map[string]interface{}{"img": img})
	ctx.JSON(http.StatusCreated, response)
}

func (c *userController) GetUser(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	log.Println("tes", id)
	user := c.userService.GetUser(id)
	res := helpers.BuildResponse(true, "OK", user)
	context.JSON(http.StatusOK, res)

}
