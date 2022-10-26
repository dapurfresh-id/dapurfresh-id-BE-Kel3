package services

import (
	"context"

	"time"

	"github.com/aldisaputra17/dapur-fresh-id/entities"
	"github.com/aldisaputra17/dapur-fresh-id/repositories"
	"github.com/aldisaputra17/dapur-fresh-id/request"
	"github.com/google/uuid"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, categoryReq request.RequestCategory) (*entities.Category, error)
	ReadCategory(ctx context.Context, category *entities.Category) (*entities.Category, error)
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

func (service *categoryService) ReadCategory(ctx context.Context, category *entities.Category) (*entities.Category, error) {
	res, err := service.categoryRepository.Read(ctx, category)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (service *categoryService) CreateCategory(ctx context.Context, categoryReq request.RequestCategory) (*entities.Category, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	categoryCreate := &entities.Category{
		ID:   id,
		Name: categoryReq.Name,
	}
	ctx, cancel := context.WithTimeout(ctx, service.contextTimeout)
	defer cancel()

	res, err := service.categoryRepository.Create(ctx, categoryCreate)
	if err != nil {
		return nil, err
	}
	return res, nil
}
