package controllers

import (
	"log"
	"net/http"

	"github.com/aldisaputra17/dapur-fresh-id/entities"
	"github.com/aldisaputra17/dapur-fresh-id/helpers"
	"github.com/aldisaputra17/dapur-fresh-id/request"
	"github.com/google/uuid"

	// "github.com/aldisaputra17/dapur-fresh-id/request"
	"github.com/aldisaputra17/dapur-fresh-id/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ImageController interface {
	Create(ctx *gin.Context)
}

type imgController struct {
	imgService services.ImageService
	db         *gorm.DB
}

func NewImgController(imgServ services.ImageService, db *gorm.DB) ImageController {
	return &imgController{
		imgService: imgServ,
		db:         db,
	}
}

func (c *imgController) Create(ctx *gin.Context) {
	var imgs entities.Image
	FileHead, _, err := ctx.Request.FormFile("file")
	log.Println(FileHead)
	if err != nil {
		res := helpers.BuildErrorResponse("failed img fetch", err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}
	upload, err := c.imgService.Create(request.ImageRequest{File: FileHead})
	imgs.ID = uuid.New().String()
	imgs.File = upload
	c.db.Create(imgs)
	if err != nil {
		res := helpers.BuildErrorResponse("Failed save img", err.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := helpers.BuildImageResponse(http.StatusOK, "Ok", map[string]interface{}{"img": imgs})
	ctx.JSON(http.StatusOK, res)
}
