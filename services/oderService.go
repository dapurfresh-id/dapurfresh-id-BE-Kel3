package services

import (
	"context"
	"fmt"
	"time"

	"github.com/aldisaputra17/dapur-fresh-id/entities"
	"github.com/aldisaputra17/dapur-fresh-id/helpers"
	"github.com/aldisaputra17/dapur-fresh-id/repositories"
	"github.com/aldisaputra17/dapur-fresh-id/request"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OrderService interface {
	Create(ctx context.Context, req *request.RequestOrderCreate) (*entities.Order, error)
	GetOrder(ctx *gin.Context, paginat *entities.Pagination) (helpers.Response, error)
	GetDetail(ctx context.Context, id string) (*entities.Order, error)
	PatchStatus(ctx context.Context, req *request.RequestPatchOrder) (*entities.Order, error)
}

type orderService struct {
	orderRepo     repositories.OrderRepository
	contexTimeOut time.Duration
}

func NewOrderService(orderRepo repositories.OrderRepository, time time.Duration) OrderService {
	return &orderService{
		orderRepo:     orderRepo,
		contexTimeOut: time,
	}
}

func (service *orderService) Create(ctx context.Context, req *request.RequestOrderCreate) (*entities.Order, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	orderCreate := &entities.Order{
		ID:        id,
		Catatan:   req.Noted,
		Address:   req.Address,
		Cost:      helpers.RandomPrice(),
		Status:    "proses",
		CartID:    req.CartID,
		UserID:    req.UserID,
		CreatedAt: time.Now(),
	}
	ctx, cancel := context.WithTimeout(ctx, service.contexTimeOut)
	defer cancel()

	res, err := service.orderRepo.Create(ctx, orderCreate)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (service *orderService) GetOrder(ctx *gin.Context, paginat *entities.Pagination) (helpers.Response, error) {
	operationResult, totalPages := service.orderRepo.GetOrder(ctx, paginat)

	if operationResult.Error != nil {
		return helpers.Response{Success: true, Message: operationResult.Error.Error()}, nil
	}

	var data = operationResult.Result.(*entities.Pagination)
	urlPath := ctx.Request.URL.Path

	searchQueryParams := ""

	for _, search := range paginat.Searchs {
		searchQueryParams += fmt.Sprintf("&%s.%s=%s", search.Column, search.Action, search.Query)
	}

	data.FirstPage = fmt.Sprintf("%s?limit=%d&page=%d&sort_order=%s", urlPath, paginat.Limit, 0, paginat.SortOrder) + searchQueryParams
	data.LastPage = fmt.Sprintf("%s?limit=%d&page=%d&sort_order=%s", urlPath, paginat.Limit, totalPages, paginat.SortOrder) + searchQueryParams

	if data.Page > 0 {
		// set previous page pagination response
		data.PreviousPage = fmt.Sprintf("%s?limit=%d&page=%d&sort_order=%s", urlPath, paginat.Limit, data.Page-1, paginat.SortOrder) + searchQueryParams
	}

	if data.Page < totalPages {
		// set next page pagination response
		data.NextPage = fmt.Sprintf("%s?limit=%d&page=%d&sort_order=%s", urlPath, paginat.Limit, data.Page+1, paginat.SortOrder) + searchQueryParams
	}

	if data.Page > totalPages {
		// reset previous page
		data.PreviousPage = ""
	}

	return helpers.BuildResponse(true, "Ok", data), nil
}

func (service *orderService) GetDetail(ctx context.Context, id string) (*entities.Order, error) {
	res, err := service.orderRepo.GetDetail(ctx, id)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, service.contexTimeOut)
	defer cancel()

	return res, nil
}

func (service *orderService) PatchStatus(ctx context.Context, req *request.RequestPatchOrder) (*entities.Order, error) {
	orderUpdate := &entities.Order{
		ID:     req.ID,
		Status: "cancel",
	}
	ctx, cancel := context.WithTimeout(ctx, service.contexTimeOut)
	defer cancel()
	res, err := service.orderRepo.PatchStatus(ctx, orderUpdate)
	if err != nil {
		return nil, err
	}
	return res, nil
}
