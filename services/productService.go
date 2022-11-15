package services

import (
	"context"
	"fmt"
	"time"

	"github.com/aldisaputra17/dapur-fresh-id/entities"
	"github.com/aldisaputra17/dapur-fresh-id/helpers"
	"github.com/aldisaputra17/dapur-fresh-id/repositories"
	"github.com/gin-gonic/gin"
)

type ProductService interface {
	FindAllProduct(ctx context.Context) (*[]entities.Product, error)
	FindProductById(ctx context.Context, productId string) (*entities.Product, error)
	FindProductByCategory(ctx context.Context, categoryId string) (*[]entities.Product, error)
	FindProductByNameEqual(ctx context.Context, name string) (*entities.Product, error)
	FindProductByNameContains(ctx context.Context, name string) (*entities.Product, error)
	FindProductByNameLike(ctx context.Context, name string) (*entities.Product, error)
	LimitProduct(ctx context.Context, limit int) (*[]entities.Product, error)
	PaginantionProduct(ctx *gin.Context, paginat *entities.Pagination) (helpers.Response, error)
	PopularProduct(ctx context.Context) (*[]entities.Product, error)
	// Create(ctx context.Context, req *request.RequestProduct) (*entities.Product, error)
	// Image(req *request.RequestProduct) (string, error)
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

// func (service *productService) Create(ctx context.Context, req *request.RequestProduct) (*entities.Product, error) {
// 	id, err := uuid.NewRandom()
// 	if err != nil {
// 		return nil, err
// 	}
// 	img, errImg := helpers.ImageUploadHelper(req.Image)
// 	if err != nil {
// 		return nil, errImg
// 	}
// 	// img, _ := service.Image(&request.RequestProduct{Image: req.Image})

// 	prodCreate := &entities.Product{
// 		ID:         id,
// 		Name:       req.Name,
// 		Price:      req.Price,
// 		Unit:       req.Unit,
// 		Image:      img,
// 		UnitType:   req.UnitType,
// 		CategoryID: req.CategoryID,
// 	}
// 	fmt.Println("uc:", prodCreate)
// 	ctx, cancel := context.WithTimeout(ctx, service.contextTimeOuT)
// 	defer cancel()

// 	prodCreate.Image = img
// 	res, err := service.productRepository.Create(ctx, prodCreate)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }

// func (service *productService) Image(req *request.RequestProduct) (string, error) {
// 	uploadFile, err := helpers.ImageUploadHelper(req.Image)
// 	if err != nil {
// 		return "", nil
// 	}
// 	return uploadFile, nil
// }

func (service *productService) FindAllProduct(ctx context.Context) (*[]entities.Product, error) {
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

func (service *productService) FindProductByCategory(ctx context.Context, categoryId string) (*[]entities.Product, error) {
	res, err := service.productRepository.FindProductByCategory(ctx, categoryId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (service *productService) FindProductByNameEqual(ctx context.Context, name string) (*entities.Product, error) {
	res, err := service.productRepository.FindProductByNameEqual(ctx, name)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (service *productService) FindProductByNameContains(ctx context.Context, name string) (*entities.Product, error) {
	res, err := service.productRepository.FindProductByNameContains(ctx, name)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (service *productService) FindProductByNameLike(ctx context.Context, name string) (*entities.Product, error) {
	res, err := service.productRepository.FindProductByNameLike(ctx, name)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (service *productService) PopularProduct(ctx context.Context) (*[]entities.Product, error) {
	res, err := service.productRepository.PopularProduct(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (service *productService) LimitProduct(ctx context.Context, limit int) (*[]entities.Product, error) {
	res, err := service.productRepository.LimitProduct(ctx, limit)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (service *productService) PaginantionProduct(ctx *gin.Context, paginat *entities.Pagination) (helpers.Response, error) {
	operationResult, totalPages := service.productRepository.PaginationProduct(paginat)

	if operationResult.Error != nil {
		return helpers.Response{Success: true, Message: operationResult.Error.Error()}, nil
	}

	var data = operationResult.Result.(*entities.Pagination)

	urlPath := ctx.Request.URL.Path

	searchQueryParams := ""

	for _, search := range paginat.Searchs {
		searchQueryParams += fmt.Sprintf("&%s.%s=%s", search.Column, search.Action, search.Query)
	}

	data.FirstPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, paginat.Limit, 0, paginat.Sort) + searchQueryParams
	data.LastPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, paginat.Limit, totalPages, paginat.Sort) + searchQueryParams

	if data.Page > 0 {
		// set previous page pagination response
		data.PreviousPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, paginat.Limit, data.Page-1, paginat.Sort) + searchQueryParams
	}

	if data.Page < totalPages {
		// set next page pagination response
		data.NextPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, paginat.Limit, data.Page+1, paginat.Sort) + searchQueryParams
	}

	if data.Page > totalPages {
		// reset previous page
		data.PreviousPage = ""
	}
	return helpers.BuildResponse(true, "Ok", data), nil
}
