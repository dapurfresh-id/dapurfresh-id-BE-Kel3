package services

import (
	"context"

	"time"

	"github.com/aldisaputra17/dapur-fresh-id/entities"
	"github.com/aldisaputra17/dapur-fresh-id/repositories"
)

type ProductService interface {
	FindAllProduct(ctx context.Context) ([]*entities.Product, error)
	FindProductById(ctx context.Context, productId string) (*entities.Product, error)
}

type productService struct {
	productRepository repositories.ProductRepository
	contextTimeout    time.Duration
}

func NewProductService(productRepo repositories.ProductRepository, time time.Duration) ProductService {
	return &productService{
		productRepository: productRepo,
		contextTimeout:    time,
	}
}

func (service *productService) FindAllProduct(ctx context.Context) ([]*entities.Product, error) {
	res, err := service.productRepository.FindAllProduct(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (service *productService) FindProductById(ctx context.Context, productId string) (*entities.Product, error) {
	res, err := service.productRepository.FindProductById(ctx, productId)
	if err != nil {
		return nil, err
	}
	return res, nil
}
