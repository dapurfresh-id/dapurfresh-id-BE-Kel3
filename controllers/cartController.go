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
	GetCount(ctx *gin.Context)
	GetCarts(ctx *gin.Context)
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
		fmt.Println("cart:", reqCart)
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

func (c *cartController) GetCount(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	counts := 0
	getCart, err := c.cartService.GetCount(ctx, userID, int64(counts))
	if err != nil {
		res := helpers.BuildErrorResponse("Data not found", "No data with given cart_id", helpers.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	} else {
		response := helpers.BuildResponse(true, "Ok", getCart)
		ctx.JSON(http.StatusOK, response)
	}
}

func (c *cartController) GetCarts(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	list, err := c.cartService.GetCarts(ctx, userID)
	total := c.cartService.GetTotalCartValue(list)
	var counts int64
	count, _ := c.cartService.GetCount(ctx, userID, counts)
	if err != nil {
		res := helpers.BuildErrorResponse("Fail fetch list carts", err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helpers.BuildSuccessAddCart(true, "Ok", list, total, count)
	ctx.JSON(http.StatusOK, res)
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
