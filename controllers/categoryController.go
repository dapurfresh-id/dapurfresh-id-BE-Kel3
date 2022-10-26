package controllers

import (
	"net/http"

	"github.com/aldisaputra17/dapur-fresh-id/entities"
	"github.com/aldisaputra17/dapur-fresh-id/helpers"
	"github.com/aldisaputra17/dapur-fresh-id/request"
	"github.com/aldisaputra17/dapur-fresh-id/services"
	"github.com/gin-gonic/gin"
)

type CategoryController interface {
	Create(ctx *gin.Context)
	Read(ctx *gin.Context)
}

type categoryController struct {
	categoryService services.CategoryService
}

func NewCategoryController(categoryService services.CategoryService) CategoryController {
	return &categoryController{
		categoryService: categoryService,
	}
}

func (c *categoryController) Read(ctx *gin.Context) {
	var reqCategory request.RequestCategory
	var category entities.Category
	errObj := ctx.ShouldBind(&reqCategory)
	readedCategory, err := c.categoryService.ReadCategory(ctx, &category)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to readed", errObj.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else {
		response := helpers.BuildResponse(true, "Readed!", readedCategory)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (c *categoryController) Create(ctx *gin.Context) {
	var reqCategory request.RequestCategory
	errObj := ctx.ShouldBind(&reqCategory)
	if errObj != nil {
		response := helpers.BuildErrorResponse("Failed to process create", errObj.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	createdCategory, err := c.categoryService.CreateCategory(ctx, reqCategory)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to created", errObj.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else {
		response := helpers.BuildResponse(true, "Created!", createdCategory)
		ctx.JSON(http.StatusCreated, response)
	}
}
