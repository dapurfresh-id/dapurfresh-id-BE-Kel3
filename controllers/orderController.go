package controllers

import (
	"fmt"
	"net/http"

	"github.com/aldisaputra17/dapur-fresh-id/helpers"
	"github.com/aldisaputra17/dapur-fresh-id/request"
	"github.com/aldisaputra17/dapur-fresh-id/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type OrderController interface {
	Create(ctx *gin.Context)
	GetOrder(ctx *gin.Context)
	GetDetail(ctx *gin.Context)
	PatchStatus(ctx *gin.Context)
}

type orderController struct {
	orderService services.OrderService
	jwtService   services.JWTService
}

func NewOrderController(odr services.OrderService, jwtServ services.JWTService) OrderController {
	return &orderController{
		orderService: odr,
		jwtService:   jwtServ,
	}
}

func (c *orderController) Create(ctx *gin.Context) {
	var req *request.RequestOrderCreate

	errObj := ctx.ShouldBind(&req)
	if errObj != nil {
		res := helpers.BuildErrorResponse("failed process order", errObj.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusUnprocessableEntity, res)
		return
	} else {
		authHeader := ctx.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		req.UserID = userID
		fmt.Println("order", req)
		res, err := c.orderService.Create(ctx, req)
		if err != nil {
			res := helpers.BuildErrorResponse("Failed to created order", err.Error(), helpers.EmptyObj{})
			ctx.JSON(http.StatusBadRequest, res)
			return
		}
		response := helpers.BuildResponse(true, "Crated!", res)
		ctx.JSON(http.StatusCreated, response)
	}

}

func (c *orderController) GetOrder(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, errTkn := c.jwtService.ValidateToken(authHeader)
	if errTkn != nil {
		panic(errTkn.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	list, err := c.orderService.GetOrder(ctx, userID)
	if err != nil {
		res := helpers.BuildErrorResponse("Failed to list order", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helpers.BuildResponse(true, "Ok", list)
	ctx.JSON(http.StatusOK, res)
}

func (c *orderController) GetDetail(ctx *gin.Context) {
	id := ctx.Param("id")
	authHeader := ctx.GetHeader("Authorization")
	_, errTkn := c.jwtService.ValidateToken(authHeader)
	if errTkn != nil {
		panic(errTkn.Error())
	}
	list, err := c.orderService.GetDetail(ctx, id)
	if err != nil {
		res := helpers.BuildErrorResponse("Failed to list detail order", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helpers.BuildResponse(true, "Ok", list)
	ctx.JSON(http.StatusOK, res)
}

func (c *orderController) PatchStatus(ctx *gin.Context) {
	var req *request.RequestPatchOrder
	errObj := ctx.BindJSON(&req)
	if errObj != nil {
		res := helpers.BuildErrorResponse("failed process cart", errObj.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	authHeader := ctx.GetHeader("Authorization")
	token, errTkn := c.jwtService.ValidateToken(authHeader)
	if errTkn != nil {
		panic(errTkn.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	req.UserID = userID
	result, err := c.orderService.PatchStatus(ctx, req)
	if err != nil {
		res := helpers.BuildErrorResponse("Failed to patch order", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := helpers.BuildResponse(true, "Ok", result)
	ctx.JSON(http.StatusOK, res)
}

func (c *orderController) getUserIDByToken(token string) string {
	Token, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := Token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
