package controllers

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	// "strconv"

	"github.com/aldisaputra17/dapur-fresh-id/helpers"
	"github.com/aldisaputra17/dapur-fresh-id/request"
	"github.com/aldisaputra17/dapur-fresh-id/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController interface {
	UpdateUser(context *gin.Context)
	saveFileHandler(context *gin.Context)
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
	id := fmt.Sprintf("%v", claims["user_id"])
	log.Println("tes up-control", id)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// var err1 = context.SaveUploadedFile(userReqUpdate.Image, "assets/"+userReqUpdate.Image.Filename)
	// if err1 != nil {
	// 	context.String(http.StatusInternalServerError, "unknown error")
	// 	return
	// }

	userReqUpdate.ID = uuid.Must(uuid.Parse(id))

	// coba upload
	// file, err := context.FormFile("file")

	// // The file cannot be received.
	// if err != nil {
	// 	context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	// 		"message": "No file is received",
	// 	})
	// 	return
	// }

	// // Retrieve file information
	// extension := filepath.Ext(file.Filename)
	// // Generate random file name for the new uploaded file so it doesn't override the old file with same name
	// image := uuid.New().String() + extension

	// // The file is received, so let's save it
	// if err := context.SaveUploadedFile(file, "assets/"+image); err != nil {
	// 	context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
	// 		"message": "Unable to save the file",
	// 	})
	// 	return
	// }

	// userReqUpdate.Image = image
	u := c.userService.UpdateUser(userReqUpdate)
	res := helpers.BuildResponse(true, "OK!", u)
	context.JSON(http.StatusOK, res)
}

// func (c *userController) UpdateUser(context *gin.Context) {
// 	authHeader := context.GetHeader("Authorization")
// 	token, errToken := c.jwtService.ValidateToken(authHeader)
// 	if errToken != nil {
// 		panic(errToken.Error())
// 	}
// 	claims := token.Claims.(jwt.MapClaims)
// 	id := fmt.Sprintf("%v", claims["user_id"])

// 	cari := c.userService.GetUser(id)
// 	if cari != nil {
// 		var input request.RequestUpdateUser

// 		errObj := context.ShouldBind(&input)
// 		if errObj != nil {
// 			res := helpers.BuildErrorResponse("Failed to process request", errObj.Error(), helpers.EmptyObj{})
// 			context.AbortWithStatusJSON(http.StatusBadRequest, res)
// 			return
// 		}
// 		u := c.userService.UpdateUser(input)
// 		res := helpers.BuildResponse(true, "OK!", u)
// 		context.JSON(http.StatusOK, res)
// 	}
// }

func (c *userController) saveFileHandler(context *gin.Context) {
	file, err := context.FormFile("file")

	// The file cannot be received.
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	// Retrieve file information
	extension := filepath.Ext(file.Filename)
	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	newFileName := uuid.New().String() + extension

	// The file is received, so let's save it
	if err := context.SaveUploadedFile(file, "/assets"+newFileName); err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}

	// File saved successfully. Return proper result
	context.JSON(http.StatusOK, gin.H{
		"message": "Your file has been successfully uploaded.",
	})
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
