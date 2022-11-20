package repositories

import (
	"context"
	"fmt"

	"github.com/aldisaputra17/dapur-fresh-id/entities"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(ctx context.Context, order *entities.Order) (*entities.Order, error)
	GetCartID(ctx context.Context, CartID string, UserID string) ([]*entities.Order, error)
	GetOrder(ctx context.Context, userID string) ([]*entities.Order, error)
	GetDetail(ctx context.Context, id string) (*entities.Order, error)
	PatchStatus(ctx context.Context, order *entities.Order) (*entities.Order, error)
}

type orderConnection struct {
	connection *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderConnection{
		connection: db,
	}
}

func (db *orderConnection) GetCartID(ctx context.Context, CartID string, UserID string) ([]*entities.Order, error) {
	var order []*entities.Order
	res := db.connection.WithContext(ctx).Where("cart_id = ? and user_id = ?", CartID, UserID).Find(&order)
	if res.Error != nil {
		return nil, res.Error
	}
	return order, nil
}

func (db *orderConnection) Create(ctx context.Context, order *entities.Order) (*entities.Order, error) {
	CartRp := NewCartRepository(db.connection)

	cartItem, err := CartRp.GetCartByid(ctx, order.CartID)
	if err != nil {
		return nil, err
	}

	if len(cartItem) <= 0 {
		lenErr := fmt.Errorf("cart dosent exists")
		return nil, lenErr
	}
	order.UserID = cartItem[0].UserID
	order.SubTotal = cartItem[0].SubTotal
	order.Total = cartItem[0].SubTotal + order.Cost
	fmt.Println("model:", order.CartID)

	orderItem, err := db.GetCartID(ctx, order.CartID, order.UserID)
	if err != nil {
		return nil, err
	}

	if len(orderItem) < 0 {
		fmt.Println("order dosent exists")
		return nil, fmt.Errorf("order not foud, please add order again")
	}
	res := db.connection.WithContext(ctx).Save(&order)
	db.connection.Preload("User").Preload("Carts.Products").Find(&order)
	if res.Error != nil {
		return nil, res.Error
	}
	CartRp.Trancate(ctx, order.UserID, order.CartID)
	return order, nil
}

func (db *orderConnection) GetOrder(ctx context.Context, userID string) ([]*entities.Order, error) {
	var orders []*entities.Order
	res := db.connection.WithContext(ctx).Where("user_id = ?", userID).Preload("User").Preload("Carts.Products.Image").Find(&orders)
	if res.Error != nil {
		return nil, res.Error
	}
	return orders, nil
}

func (db *orderConnection) GetDetail(ctx context.Context, id string) (*entities.Order, error) {
	var order *entities.Order
	res := db.connection.WithContext(ctx).Where("id = ?", id).Preload("User").Preload("Carts.Products.Image").Find(&order)
	if res.Error != nil {
		return nil, res.Error
	}
	return order, nil
}

func (db *orderConnection) PatchStatus(ctx context.Context, order *entities.Order) (*entities.Order, error) {
	orderID := fmt.Sprintf("%v", order.ID)
	orderItem, err := db.GetDetail(ctx, orderID)
	if err != nil {
		return nil, err
	}
	orderItem.Status = "cancel"
	res := db.connection.WithContext(ctx).Model(&order).Where("id = ?", order.ID).Updates(entities.Order{
		Status: order.Status}).Preload("User").Preload("Carts.Products").Find(&order)
	if res.Error != nil {
		return nil, res.Error
	}
	return order, nil
}
