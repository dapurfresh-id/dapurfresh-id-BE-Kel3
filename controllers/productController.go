package controllers

import (
	"fmt"
	"net/http"

	"github.com/aldisaputra17/dapur-fresh-id/helpers"
	"github.com/aldisaputra17/dapur-fresh-id/request"
	"github.com/aldisaputra17/dapur-fresh-id/services"
	"github.com/gin-gonic/gin"
)

type ProductController interface {
	Create(ctx *gin.Context)
}

type productController struct {
	prodService services.ProductService
}

func NewProductController(prodService services.ProductService) ProductController {
	return &productController{
		prodService: prodService,
	}
}

func (c *productController) Create(ctx *gin.Context) {
	var req *request.ReqeustCreateProduct
	err := ctx.ShouldBind(&req)
	if err != nil {
		res := helpers.BuildErrorResponse("failed proccess product", err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	result, err := c.prodService.Create(ctx, req)
	if err != nil {
		res := helpers.BuildErrorResponse("failed create product", err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	fmt.Println("product:", result)
	if err != nil {
		res := helpers.BuildErrorResponse("failed create product", err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helpers.BuildResponse(true, "Created", result)
	ctx.JSON(http.StatusCreated, res)
}
