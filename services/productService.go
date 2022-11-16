package services

import (
	"context"
	"fmt"
	"time"

	"github.com/aldisaputra17/dapur-fresh-id/entities"
	"github.com/aldisaputra17/dapur-fresh-id/repositories"
	"github.com/aldisaputra17/dapur-fresh-id/request"
	"github.com/google/uuid"
)

type ProductService interface {
	Create(ctx context.Context, req *request.ReqeustCreateProduct) (*entities.Product, error)
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
	// img, _ := service.Image(&request.ReqeustCreateProduct{Image: req.Image})

	prodCreate := &entities.Product{
		ID:         id,
		Name:       req.Name,
		Price:      req.Price,
		Unit:       req.Unit,
		ImageID:    req.ImageID,
		UnitType:   req.UnitType,
		CategoryID: req.CategoryID,
	}
	fmt.Println("uc:", prodCreate)
	ctx, cancel := context.WithTimeout(ctx, service.contextTimeOuT)
	defer cancel()

	res, err := service.prodRepo.Create(ctx, prodCreate)
	if err != nil {
		return nil, err
	}
	return res, nil
}
