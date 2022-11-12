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
	GetCarts(ctx context.Context, userID string) ([]*entities.Cart, error)
	CovertDetailCartMap(ctx context.Context, userID string) (map[string](*entities.Cart), error)
	Update(ctx context.Context, cart *entities.Cart) (*entities.Cart, error)
	UpdateDetailCart(ctx context.Context, cart *entities.Cart) (*entities.Cart, error)
	Delete(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (*entities.Cart, error)
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
	// var carts []*entities.Cart
	res := db.connection.WithContext(ctx).Model(&entities.Cart{}).Where("user_id = ?", userID).Count(&count)
	if res.Error != nil {
		return 0, res.Error
	}
	return count, nil
}

func (db *cartConnection) GetCarts(ctx context.Context, userID string) ([]*entities.Cart, error) {
	var carts []*entities.Cart
	res := db.connection.WithContext(ctx).Where("user_id = ?", userID).Preload("Products").Preload("User").Find(&carts)
	if res.Error != nil {
		return nil, res.Error
	}
	return carts, nil
}

func (db *cartConnection) CovertDetailCartMap(ctx context.Context, userID string) (map[string](*entities.Cart), error) {
	cartList, err := db.GetCarts(ctx, userID)
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
	query := `UPDATE carts set sub_total = ? WHERE id = ?`
	stmt, err := db.connection.ConnPool.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	res, err := stmt.ExecContext(ctx, cart.SubTotal, cart.ID)
	if err != nil {
		return nil, err
	}

	rowAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowAffected != 1 {
		err = fmt.Errorf("Weird  Behaviour. Total Affected: %d", rowAffected)
		return nil, err
	}
	return cart, nil
}

func (db *cartConnection) Delete(ctx context.Context, id string) error {
	var cart *entities.Cart
	res := db.connection.WithContext(ctx).Where("id = ?", id).Delete(&cart)
	if res.Error != nil {
		return nil
	}
	return nil
}

func (db *cartConnection) UpdateDetailCart(ctx context.Context, cart *entities.Cart) (*entities.Cart, error) {
	cartID := fmt.Sprintf("%v", cart.ID)
	cartItems, err := db.FindById(ctx, cartID)
	if err != nil {
		return nil, err
	}

	cart.SubTotal = cartItems.Price * cart.Quantity

	res := db.connection.WithContext(ctx).Model(&cart).Updates(entities.Cart{
		Quantity: cart.Quantity, SubTotal: cart.SubTotal})
	db.connection.Preload("Products").Find(&cart)
	if res.Error != nil {
		return nil, res.Error
	}
	cart.SubTotal = cart.Price * cart.Quantity
	return cart, nil
}

func (db *cartConnection) FindById(ctx context.Context, id string) (*entities.Cart, error) {
	var cart *entities.Cart
	res := db.connection.WithContext(ctx).Where("id = ?", id).Find(&cart)
	if res.Error != nil {
		return nil, res.Error
	}
	return cart, nil
}
