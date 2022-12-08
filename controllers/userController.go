package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aldisaputra17/dapur-fresh-id/entities"
	"github.com/aldisaputra17/dapur-fresh-id/helpers"
	"github.com/aldisaputra17/dapur-fresh-id/repositories"
	"github.com/aldisaputra17/dapur-fresh-id/request"
	"github.com/aldisaputra17/dapur-fresh-id/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController interface {
	Update(ctx *gin.Context)
	GetUser(context *gin.Context)
}

type userController struct {
	userService services.UserService
	jwtService  services.JWTService
	db          *gorm.DB
}

func NewUserController(userServ services.UserService, jwtServ services.JWTService, db *gorm.DB) UserController {
	return &userController{
		userService: userServ,
		jwtService:  jwtServ,
		db:          db,
	}
}

func (c *userController) Update(ctx *gin.Context) {
	var user entities.User
	uRp := repositories.NewUserRepository(c.db)
	fileHead, _, err := ctx.Request.FormFile("image")
	log.Println(fileHead)
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	result, err := c.userService.UploadImage(ctx, &request.RequestImgUpdate{Image: fileHead})
	user.Image = result
	update, err := uRp.Update(ctx, user, id)
	if err != nil {
		res := helpers.BuildErrorResponse("Failed to update user", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	response := helpers.BuildResponse(true, "Updated!", update)
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
