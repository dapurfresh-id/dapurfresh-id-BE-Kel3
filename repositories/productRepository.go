package repositories

import (
	"context"

	"github.com/aldisaputra17/dapur-fresh-id/entities"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProducts(ctx context.Context, id string) ([]*entities.Product, error)
}

type productConnection struct {
	connection *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productConnection{
		connection: db,
	}
}

func (db *productConnection) GetProducts(ctx context.Context, id string) ([]*entities.Product, error) {
	var products []*entities.Product
	res := db.connection.WithContext(ctx).Where("id = ?", id).Find(&products)
	if res.Error != nil {
		return nil, res.Error
	}
	return products, nil
}