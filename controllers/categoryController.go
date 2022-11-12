package controllers

import (
	"net/http"

	"github.com/aldisaputra17/dapur-fresh-id/helpers"
	"github.com/aldisaputra17/dapur-fresh-id/request"
	"github.com/aldisaputra17/dapur-fresh-id/services"
	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	GetAllCategory(ctx *gin.Context)
	GetCategoryById(ctx *gin.Context)
	CreateCategory(ctx *gin.Context)
}

type categoryController struct {
	categoryService services.CategoryService
}

func NewCategoryController(categoryService services.CategoryService) CategoryController {
	return &categoryController{
		categoryService: categoryService,
	}
}

func (c *categoryController) GetAllCategory(ctx *gin.Context) {
	readedCategory, err := c.categoryService.FindAllCategory(ctx)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to readed", err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else {
		response := helpers.BuildResponse(true, "Readed!", readedCategory)
		ctx.JSON(http.StatusOK, response)
	}
}

func (c *categoryController) GetCategoryById(ctx *gin.Context) {
	categoryId := ctx.Param("id")
	foundCategory, err := c.categoryService.FindById(ctx, categoryId)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to readed", err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else {
		response := helpers.BuildResponse(true, "Readed!", foundCategory)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (c *categoryController) CreateCategory(ctx *gin.Context) {
	var req *request.RequestCreateCategory
	err := ctx.ShouldBind(&req)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to proccess", err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	result, err := c.categoryService.CreateCategory(ctx, req)
	if err != nil {
		res := helpers.BuildErrorResponse("failed create category", err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helpers.BuildResponse(true, "Created", result)
	ctx.JSON(http.StatusCreated, res)
}
