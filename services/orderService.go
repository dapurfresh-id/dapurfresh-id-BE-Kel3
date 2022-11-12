package services

import (
	"context"
	"time"

	"github.com/aldisaputra17/dapur-fresh-id/entities"
	"github.com/aldisaputra17/dapur-fresh-id/helpers"
	"github.com/aldisaputra17/dapur-fresh-id/repositories"
	"github.com/aldisaputra17/dapur-fresh-id/request"
	"github.com/google/uuid"
)

type OrderService interface {
	Create(ctx context.Context, req *request.RequestOrderCreate) (*entities.Order, error)
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
