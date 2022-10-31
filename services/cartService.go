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

type CartService interface {
	AddCart(ctx context.Context, cartReq *request.RequestCart) (*entities.Cart, error)
}

type cartService struct {
	repository     repositories.CartRepository
	contextTimeOut time.Duration
}

func NewCartService(repo repositories.CartRepository, time time.Duration) CartService {
	return &cartService{
		repository:     repo,
		contextTimeOut: time,
	}
}

func (service *cartService) AddCart(ctx context.Context, cartReq *request.RequestCart) (*entities.Cart, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	cartCreate := &entities.Cart{
		ID:        id,
		Quantity:  cartReq.Quantity,
		ProductID: cartReq.ProductID,
		UserID:    cartReq.UserID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	fmt.Println("uc:", cartCreate)
	ctx, cancel := context.WithTimeout(ctx, service.contextTimeOut)
	defer cancel()

	res, err := service.repository.AddCart(ctx, cartCreate)
	if err != nil {
		return nil, err
	}
	fmt.Println("id", res)
	return res, nil
}
