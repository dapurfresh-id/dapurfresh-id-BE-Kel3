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
	GetAllProduct(ctx *gin.Context)
	GetProductById(ctx *gin.Context)
	GetProductByCategory(ctx *gin.Context)
	PaginationProduct(ctx *gin.Context)
	GetPopularProduct(ctx *gin.Context)
}

type productController struct {
	productService services.ProductService
}

func NewProductController(productService services.ProductService) ProductController {
	return &productController{
		productService: productService,
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
	result, err := c.productService.Create(ctx, req)
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

func (c *productController) GetAllProduct(ctx *gin.Context) {
	readedProduct, err := c.productService.FindAllProduct(ctx)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to readed", err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else {
		response := helpers.BuildResponse(true, "Readed!", readedProduct)
		ctx.JSON(http.StatusOK, response)
	}
	fmt.Println("product:", readedProduct)
	if err != nil {
		res := helpers.BuildErrorResponse("failed create product", err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helpers.BuildResponse(true, "Created", readedProduct)
	ctx.JSON(http.StatusCreated, res)
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
	categoryId := ctx.Param("category_id")
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

func (c *productController) GetPopularProduct(ctx *gin.Context) {
	readedProduct, err := c.productService.PopularProduct(ctx)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to readed", err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else {
		response := helpers.BuildResponse(true, "Readed!", readedProduct)
		ctx.JSON(http.StatusOK, response)
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
