package services

import (
	"context"
	"fmt"
	"time"

	"github.com/aldisaputra17/dapur-fresh-id/entities"
	"github.com/aldisaputra17/dapur-fresh-id/helpers"
	"github.com/aldisaputra17/dapur-fresh-id/repositories"
	"github.com/aldisaputra17/dapur-fresh-id/request"
	"github.com/google/uuid"
)

type ProductService interface {
	Create(ctx context.Context, req *request.ReqeustCreateProduct) (*entities.Product, error)
	Image(req *request.ReqeustCreateProduct) (string, error)
}

type productService struct {
	prodRepo       repositories.ProductRepository
	contextTimeOuT time.Duration
}

func NewProductService(r repositories.ProductRepository, time time.Duration) ProductService {
	return &productService{
		prodRepo:       r,
		contextTimeOuT: time,
	}
}

func (service *productService) Create(ctx context.Context, req *request.ReqeustCreateProduct) (*entities.Product, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	img, errImg := helpers.ImageUploadHelper(req.Image)
	if err != nil {
		return nil, errImg
	}
	// img, _ := service.Image(&request.ReqeustCreateProduct{Image: req.Image})

	prodCreate := &entities.Product{
		ID:         id,
		Name:       req.Name,
		Price:      req.Price,
		Unit:       req.Unit,
		Image:      img,
		UnitType:   req.UnitType,
		CategoryID: req.CategoryID,
	}
	fmt.Println("uc:", prodCreate)
	ctx, cancel := context.WithTimeout(ctx, service.contextTimeOuT)
	defer cancel()

	prodCreate.Image = img
	res, err := service.prodRepo.Create(ctx, prodCreate)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (service *productService) Image(req *request.ReqeustCreateProduct) (string, error) {
	uploadFile, err := helpers.ImageUploadHelper(req.Image)
	if err != nil {
		return "", nil
	}
	return uploadFile, nil
}
