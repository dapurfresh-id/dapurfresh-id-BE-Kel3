package services

import (
	"context"

	"time"

	"github.com/aldisaputra17/dapur-fresh-id/entities"
	"github.com/aldisaputra17/dapur-fresh-id/repositories"
)

type CategoryService interface {
	FindAllCategory(ctx context.Context) ([]*entities.Category, error)
	FindById(ctx context.Context, categoryId string) (*entities.Category, error)
}

type categoryService struct {
	categoryRepository repositories.CategoryRepository
	contextTimeout     time.Duration
}

func NewCategoryService(categoryRepo repositories.CategoryRepository, time time.Duration) CategoryService {
	return &categoryService{
		categoryRepository: categoryRepo,
		contextTimeout:     time,
	}
}

func (service *categoryService) FindAllCategory(ctx context.Context) ([]*entities.Category, error) {
	res, err := service.categoryRepository.FindAllCategory(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (service *categoryService) FindById(ctx context.Context, categoryId string) (*entities.Category, error) {
	res, err := service.categoryRepository.FindById(ctx, categoryId)
	if err != nil {
		return nil, err
	}
	return res, nil
}
