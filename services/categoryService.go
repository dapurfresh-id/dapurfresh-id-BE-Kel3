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

type CategoryService interface {
	FindAllCategory(ctx context.Context) ([]*entities.Category, error)
	FindById(ctx context.Context, categoryId string) (*entities.Category, error)
	CreateCategory(ctx context.Context, req *request.RequestCreateCategory) (*entities.Category, error)
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

func (service *categoryService) CreateCategory(ctx context.Context, req *request.RequestCreateCategory) (*entities.Category, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	catCreate := &entities.Category{
		ID:      id,
		ImageID: req.ImageID,
		Name:    req.Name,
	}
	fmt.Println("uc:", catCreate)
	ctx, cancel := context.WithTimeout(ctx, service.contextTimeout)
	defer cancel()

	res, err := service.categoryRepository.CreateCategory(ctx, catCreate)
	if err != nil {
		return nil, err
	}
	return res, nil
}
