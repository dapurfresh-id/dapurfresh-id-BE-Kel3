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

	cart.TotalPrice = prodItems[0].Price * cart.Quantity
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
