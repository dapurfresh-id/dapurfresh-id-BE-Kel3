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

func (c *orderController) getUserIDByToken(token string) string {
	Token, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := Token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
