package controllers

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"

	"github.com/aldisaputra17/dapur-fresh-id/entities"
	"github.com/aldisaputra17/dapur-fresh-id/helpers"
	"github.com/aldisaputra17/dapur-fresh-id/request"
	"github.com/aldisaputra17/dapur-fresh-id/services"
	"github.com/gin-gonic/gin"
)

type CartController interface {
	AddCart(ctx *gin.Context)
	GetCarts(ctx *gin.Context)
	GetCart(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
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
	fmt.Println("carts", list)
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

func (c *cartController) Delete(ctx *gin.Context) {
	var cart entities.Cart
	id := ctx.Param("id")
	cart.ID = uuid.Must(uuid.Parse(id))
	authHeader := ctx.GetHeader("Authorization")
	token, errTkn := c.jwtService.ValidateToken(authHeader)
	if errTkn != nil {
		panic(errTkn)
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.cartService.IsAllowedToEdit(ctx, id, userID) {
		err := c.cartService.Delete(ctx, cart)
		if err != nil {
			res := helpers.BuildErrorResponse("Fail deleted carts", err.Error(), helpers.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		}
		res := helpers.BuildResponse(true, "Ok", helpers.EmptyObj{})
		ctx.JSON(http.StatusOK, res)
	} else {
		response := helpers.BuildErrorResponse("You dont have permission", "You are not the owner", helpers.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)

	}
}

func (c *cartController) Update(ctx *gin.Context) {
	var req *request.RequestCartUpdate
	errObj := ctx.BindJSON(&req)
	if errObj != nil {
		res := helpers.BuildErrorResponse("failed process cart", errObj.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	authHeader := ctx.GetHeader("Authorization")
	token, errTkn := c.jwtService.ValidateToken(authHeader)
	if errTkn != nil {
		panic(errTkn)
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	cartID := fmt.Sprintf("%v", req.ID)
	if c.cartService.IsAllowedToEdit(ctx, cartID, userID) {
		req.UserID = userID
	}
	result, err := c.cartService.Update(ctx, req)
	if err != nil {
		res := helpers.BuildErrorResponse("Failed updated cart check your permissons", err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	response := helpers.BuildResponse(true, "Updated", result)
	ctx.JSON(http.StatusOK, response)
}

func (c *cartController) GetCart(ctx *gin.Context) {
	id := ctx.Param("id")
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	result, err := c.cartService.GetCart(ctx, id, userID)
	if err != nil {
		res := helpers.BuildErrorResponse("Failed updated cart check your permissons", err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helpers.BuildResponse(true, "Ok", result)
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
