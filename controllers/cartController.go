package controllers

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"github.com/aldisaputra17/dapur-fresh-id/helpers"
	"github.com/aldisaputra17/dapur-fresh-id/request"
	"github.com/aldisaputra17/dapur-fresh-id/services"
	"github.com/gin-gonic/gin"
)

type CartController interface {
	AddCart(ctx *gin.Context)
}

type cartController struct {
	cartService services.CartService
	jwtService  services.JWTService
}

func NewCartController(cartServ services.CartService, jwtServ services.JWTService) CartController {
	return &cartController{
		cartService: cartServ,
		jwtService:  jwtServ,
	}
}

func (c *cartController) AddCart(ctx *gin.Context) {
	var reqCart request.RequestCart
	err := ctx.ShouldBind(&reqCart)
	if err != nil {
		res := helpers.BuildErrorResponse("Failed created cart", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusUnprocessableEntity, res)
		return
	} else {
		authHeader := ctx.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		reqCart.UserID = userID
		fmt.Println("product:", reqCart)
		result, err := c.cartService.AddCart(ctx, &reqCart)
		if err != nil {
			res := helpers.BuildErrorResponse("Failed to created cart", err.Error(), helpers.EmptyObj{})
			ctx.JSON(http.StatusBadRequest, res)
			return
		}
		response := helpers.BuildResponse(true, "Created!", result)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (c *cartController) getUserIDByToken(token string) string {
	Token, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := Token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
