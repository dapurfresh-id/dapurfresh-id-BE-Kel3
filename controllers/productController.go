package controllers

import (
	"net/http"
	"strconv"

	"github.com/aldisaputra17/dapur-fresh-id/helpers"
	"github.com/aldisaputra17/dapur-fresh-id/request"
	"github.com/aldisaputra17/dapur-fresh-id/services"
	"github.com/gin-gonic/gin"
)

type ProductController interface {
	GetAllProduct(ctx *gin.Context)
	GetProductById(ctx *gin.Context)
	GetProductByCategory(ctx *gin.Context)
	GetLimitProduct(ctx *gin.Context)
	PaginationProduct(ctx *gin.Context)
}

type productController struct {
	productService services.ProductService
}

func NewProductController(productService services.ProductService) ProductController {
	return &productController{
		productService: productService,
	}
}

func (c *productController) GetAllProduct(ctx *gin.Context) {
	var reqProduct request.RequestProduct
	errObj := ctx.ShouldBind(&reqProduct)
	readedProduct, err := c.productService.FindAllProduct(ctx)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to readed", errObj.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else {
		response := helpers.BuildResponse(true, "Readed!", readedProduct)
		ctx.JSON(http.StatusOK, response)
	}
}

func (c *productController) GetProductById(ctx *gin.Context) {
	productId := ctx.Param("id")
	foundProduct, err := c.productService.FindProductById(ctx, productId)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to readed", err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else {
		response := helpers.BuildResponse(true, "Readed!", foundProduct)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (c *productController) GetProductByCategory(ctx *gin.Context) {
	categoryId := ctx.Query("category_id")
	foundProduct, err := c.productService.FindProductByCategory(ctx, categoryId)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to readed", err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else {
		response := helpers.BuildResponse(true, "Readed!", foundProduct)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (c *productController) GetLimitProduct(ctx *gin.Context) {
	limit := ctx.Query("limit")
	intLimit, err := strconv.Atoi(limit)
	foundProduct, err := c.productService.LimitProduct(ctx, intLimit)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to readed", err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else {
		response := helpers.BuildResponse(true, "Readed!", foundProduct)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (c *productController) PaginationProduct(ctx *gin.Context) {
	code := http.StatusOK
	pagination := helpers.GeneratePaginationRequest(ctx)

	response, err := c.productService.PaginantionProduct(ctx, pagination)
	if err != nil {
		res := helpers.BuildErrorResponse("Failed pagination products", err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if !response.Success {
		code = http.StatusBadRequest
	}

	res := helpers.BuildResponse(true, "Ok", response)
	ctx.AbortWithStatusJSON(code, res)
}
