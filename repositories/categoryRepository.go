package repositories

import (
	"context"

	"github.com/aldisaputra17/dapur-fresh-id/entities"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Read(ctx context.Context, category *entities.Category) (*entities.Category, error)
	Create(ctx context.Context, category *entities.Category) (*entities.Category, error)
}

type categoryConnection struct {
	connection *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryConnection{
		connection: db,
	}
}

func (db *categoryConnection) Read(ctx context.Context, category *entities.Category) (*entities.Category, error) {
	res := db.connection.WithContext(ctx).Find(&category)
	if res.Error != nil {
		return nil, res.Error
	}
	return category, nil
}

func (db *categoryConnection) Create(ctx context.Context, category *entities.Category) (*entities.Category, error) {
	res := db.connection.WithContext(ctx).Create(&category)
	if res.Error != nil {
		return nil, res.Error
	}
	return category, nil
}
