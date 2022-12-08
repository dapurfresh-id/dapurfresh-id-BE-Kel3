package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/aldisaputra17/dapur-fresh-id/entities"
	"github.com/aldisaputra17/dapur-fresh-id/response"
	"gorm.io/gorm"
)

type CartRepository interface {
	AddCart(ctx context.Context, cart *entities.Cart) (*response.CartResponse, error)
	GetProdIdAndUserId(ctx context.Context, userID string, productID string) ([]*entities.Cart, error)
	GetCarts(ctx context.Context, userID string) ([]*entities.Cart, error)
	GetCartByid(ctx context.Context, id string) ([]*entities.Cart, error)
	GetCount(ctx context.Context, userID string, count int64) (int64, error)
	CovertDetailCartMap(ctx context.Context, userID string) (map[string](*entities.Cart), error)
	Update(ctx context.Context, cart *entities.Cart) (*entities.Cart, error)
	UpdateDetailCart(ctx context.Context, cart *entities.Cart) (*entities.Cart, error)
	Delete(ctx context.Context, cart entities.Cart) error
	GetCartProd(ctx context.Context) (*entities.Cart, error)
	Trancate(ctx context.Context, userID string, id string) error
	FindById(ctx context.Context, id string, userID string) (*entities.Cart, error)
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

func (db *cartConnection) AddCart(ctx context.Context, cart *entities.Cart) (*response.CartResponse, error) {
	prR := NewProductRepository(db.connection)

	prodItems, err := prR.FindProductById(ctx, cart.ProductID)
	if err != nil {
		return nil, err
	}

	cart.SubTotal = prodItems.Price * cart.Quantity
	cart.Name = prodItems.Name
	cart.Price = prodItems.Price
	cart.Unit = prodItems.Unit
	cart.UnitType = prodItems.UnitType
	prodItems.AreBuyed = prodItems.AreBuyed + cart.Quantity
	fmt.Println("model :", cart.UserID, cart.ProductID)

	cartItems, err := db.GetProdIdAndUserId(ctx, cart.UserID, cart.ProductID)
	if err != nil {
		return nil, err
	}

	if len(cartItems) > 0 {
		fmt.Println("Item already exists in cart. Please delete and add again ")
		return nil, errors.New("Item already exists in cart. Please delete and add again ")
	}

	res := db.connection.WithContext(ctx).Preload("Products.Images").Create(&cart)
	if res.Error != nil {
		return nil, err
	}
	_, errUpdate := prR.UpdateBuy(ctx, prodItems)
	if errUpdate != nil {
		return nil, errUpdate
	}
	rsp := &response.CartResponse{
		ID:        cart.ID,
		Name:      cart.Name,
		SubTotal:  cart.SubTotal,
		Price:     cart.Price,
		Quantity:  cart.Quantity,
		Unit:      cart.Unit,
		UnitType:  cart.UnitType,
		CreatedAt: cart.CreatedAt,
		ProductID: cart.ProductID,
		Products:  response.NewProductResponse(prodItems),
	}
	return rsp, nil
}

func (db *cartConnection) GetCarts(ctx context.Context, userID string) ([]*entities.Cart, error) {
	var carts []*entities.Cart
	res := db.connection.WithContext(ctx).Where("user_id = ?", userID).Preload("Products.Images").Preload("Products.Categories").Preload("User").Find(&carts)
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

func (db *cartConnection) Delete(ctx context.Context, cart entities.Cart) error {
	res := db.connection.WithContext(ctx).Delete(cart)
	if res.Error != nil {
		return nil
	}
	return nil
}

func (db *cartConnection) UpdateDetailCart(ctx context.Context, cart *entities.Cart) (*entities.Cart, error) {
	cartID := fmt.Sprintf("%v", cart.ID)
	userID := fmt.Sprintf("%v", cart.UserID)
	cartItems, err := db.FindById(ctx, cartID, userID)
	if err != nil {
		return nil, err
	}
	cart.SubTotal = cartItems.Price * cart.Quantity

	res := db.connection.WithContext(ctx).Model(&cart).Updates(entities.Cart{
		Quantity: cart.Quantity, SubTotal: cart.SubTotal})
	db.connection.Preload("Products.Images").Preload("Products.Categories").Preload("User").Find(&cart)
	if res.Error != nil {
		return nil, res.Error
	}
	cart.SubTotal = cart.Price * cart.Quantity
	return cart, nil
}

func (db *cartConnection) FindById(ctx context.Context, id string, userID string) (*entities.Cart, error) {
	var cart *entities.Cart
	res := db.connection.WithContext(ctx).Where("id = ? and user_id = ?", id, userID).Preload("Products.Images").Preload("Products.Categories").Preload("User").Find(&cart)
	if res.Error != nil {
		return nil, res.Error
	}
	return cart, nil
}

func (db *cartConnection) Trancate(ctx context.Context, userID string, id string) error {
	var cart *entities.Cart
	res := db.connection.WithContext(ctx).Where("user_id = ? and id = ?", userID, id).Delete(&cart)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (db *cartConnection) GetCartByid(ctx context.Context, id string) ([]*entities.Cart, error) {
	var carts []*entities.Cart
	res := db.connection.WithContext(ctx).Where("id = ?", id).Find(&carts)
	if res.Error != nil {
		return nil, res.Error
	}
	return carts, nil
}

func (db *cartConnection) GetCount(ctx context.Context, userID string, count int64) (int64, error) {
	res := db.connection.WithContext(ctx).Model(&entities.Cart{}).Where("user_id = ?", userID).Count(&count)
	if res.Error != nil {
		return 0, res.Error
	}
	return count, nil
}

func (db *cartConnection) GetCartProd(ctx context.Context) (*entities.Cart, error) {
	var cart *entities.Cart
	res := db.connection.WithContext(ctx).First(&cart)
	if res.Error != nil {
		return nil, res.Error
	}
	return cart, nil
}
