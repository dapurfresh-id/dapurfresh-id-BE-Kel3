package repositories

import (
	"context"

	"github.com/aldisaputra17/dapur-fresh-id/entities"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAllProduct(ctx context.Context) ([]*entities.Product, error)
	FindProductById(ctx context.Context, productId string) (*entities.Product, error)
}

type productConnection struct {
	connection *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productConnection{
		connection: db,
	}
}

func (db *productConnection) FindAllProduct(ctx context.Context) ([]*entities.Product, error) {
	var product []*entities.Product
	res := db.connection.WithContext(ctx).Find(&product)
	if res.Error != nil {
		return nil, res.Error
	}
	return product, nil
}

func (db *productConnection) FindProductById(ctx context.Context, productId string) (*entities.Product, error) {
	var product *entities.Product
	res := db.connection.WithContext(ctx).Where("id = ?", productId).First(&product)

	if res.Error != nil {
		return nil, res.Error
	}
	return product, nil
}
