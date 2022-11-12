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
	GetCount(ctx context.Context, userID string, count int64) (int64, error)
	GetCarts(ctx context.Context, userID string) ([]*entities.Cart, error)
	Refresh(ctx context.Context, userID string) error
	GetTotalCartValue(cart []*entities.Cart) int
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

func (service *cartService) GetCount(ctx context.Context, userID string, count int64) (int64, error) {
	res, err := service.repository.GetCount(ctx, userID, count)
	if err != nil {
		return 0, err
	}

	return res, nil
}

func (service *cartService) GetCarts(ctx context.Context, userID string) ([]*entities.Cart, error) {
	ctx, cancel := context.WithTimeout(ctx, service.contextTimeOut)
	defer cancel()
	updateErr := service.Refresh(ctx, userID)
	if updateErr != nil {
		return nil, updateErr
	}
	res, err := service.repository.GetCart(ctx, userID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (service *cartService) Refresh(ctx context.Context, userID string) error {
	ctx, cancel := context.WithTimeout(ctx, service.contextTimeOut)
	defer cancel()

	mCart, err := service.repository.CovertDetailCartMap(ctx, userID)
	if err != nil {
		return err
	}

	for _, v := range mCart {
		if v.SubTotal == 0 {
			v.SubTotal = v.Price * v.Quantity
			_, err := service.repository.Update(ctx, v)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (service *cartService) GetTotalCartValue(cart []*entities.Cart) int {
	total := 0
	for i := 0; i < len(cart); i++ {
		total = total + (cart[i].Quantity)*cart[i].Price
	}
	return total
}
