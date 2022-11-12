package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/aldisaputra17/dapur-fresh-id/entities"
	"gorm.io/gorm"
)

type CartRepository interface {
	AddCart(ctx context.Context, cart *entities.Cart) (*entities.Cart, error)
	GetProdIdAndUserId(ctx context.Context, userID string, productID string) ([]*entities.Cart, error)
	GetCount(ctx context.Context, userID string, count int64) (int64, error)
	GetCart(ctx context.Context, id string) ([]*entities.Cart, error)
	CovertDetailCartMap(ctx context.Context, userID string) (map[string](*entities.Cart), error)
	Update(ctx context.Context, cart *entities.Cart) (*entities.Cart, error)
}

type cartConnection struct {
	connection *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartConnection{
		connection: db,
	}
}

func (db *cartConnection) GetProdIdAndUserId(ctx context.Context, userID string, productID string) ([]*entities.Cart, error) {
	var carts []*entities.Cart
	res := db.connection.WithContext(ctx).Where("user_id = ? and product_id = ?", userID, productID).Find(&carts)
	if res.Error != nil {
		return nil, res.Error
	}
	return carts, nil
}

func (db *cartConnection) AddCart(ctx context.Context, cart *entities.Cart) (*entities.Cart, error) {
	prR := NewProductRepository(db.connection)

	prodItems, err := prR.GetProducts(ctx, cart.ProductID)
	if err != nil {
		return nil, err
	}

	if len(prodItems) <= 0 {
		lenErr := fmt.Errorf("Product doesnt exists")
		return nil, lenErr
	}

	cart.SubTotal = prodItems[0].Price * cart.Quantity
	cart.Name = prodItems[0].Name
	cart.Price = prodItems[0].Price
	fmt.Println("model :", cart.UserID, cart.ProductID)

	cartItems, err := db.GetProdIdAndUserId(ctx, cart.UserID, cart.ProductID)
	if err != nil {
		return nil, err
	}

	if len(cartItems) > 0 {
		fmt.Println("Item already exists in cart. Please delete and add again ")
		return nil, errors.New("Item already exists in cart. Please delete and add again ")
	}

	res := db.connection.WithContext(ctx).Create(&cart)
	// db.connection.Preload("Products").Find(&cart)
	if res.Error != nil {
		return nil, err
	}
	return cart, nil
}

func (db *cartConnection) GetCount(ctx context.Context, userID string, count int64) (int64, error) {
	res := db.connection.WithContext(ctx).Model(&entities.Cart{}).Where("user_id = ?", userID).Count(&count)
	if res.Error != nil {
		return 0, res.Error
	}
	return count, nil
}

func (db *cartConnection) GetCart(ctx context.Context, id string) ([]*entities.Cart, error) {
	var carts []*entities.Cart
	var count int64
	res := db.connection.WithContext(ctx).Where("id = ?", id).Find(&carts).Count(&count)
	if res.Error != nil {
		return nil, res.Error
	}
	return carts, nil
}

func (db *cartConnection) CovertDetailCartMap(ctx context.Context, userID string) (map[string](*entities.Cart), error) {
	cartList, err := db.GetCart(ctx, userID)
	if err != nil {
		return nil, err
	}
	fmt.Println("cart:", cartList)
	var mapCart = map[string]*entities.Cart{}
	for i := 0; i < len(cartList); i++ {
		mapCart[cartList[i].ProductID] = cartList[i]
	}
	return mapCart, nil
}

func (db *cartConnection) Update(ctx context.Context, cart *entities.Cart) (*entities.Cart, error) {
	res := db.connection.WithContext(ctx).Where("sub_total = ? and user_id = ?", cart.SubTotal, cart.UserID).Updates(&cart)
	if res.Error != nil {
		return nil, res.Error
	}
	return cart, nil
}
