package repositories

import (
	"context"

	"github.com/aldisaputra17/dapur-fresh-id/entities"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAllCategory(ctx context.Context) ([]*entities.Category, error)
	FindById(ctx context.Context, categoryId string) (*entities.Category, error)
	CreateCategory(ctx context.Context, category *entities.Category) (*entities.Category, error)
}

type categoryConnection struct {
	connection *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryConnection{
		connection: db,
	}
}

func (db *categoryConnection) FindAllCategory(ctx context.Context) ([]*entities.Category, error) {
	var category []*entities.Category
	res := db.connection.WithContext(ctx).Preload("Image").Find(&category)
	if res.Error != nil {
		return nil, res.Error
	}
	return category, nil
}

func (db *categoryConnection) FindById(ctx context.Context, categoryId string) (*entities.Category, error) {
	var category *entities.Category
	res := db.connection.WithContext(ctx).Where("id = ?", categoryId).Find(&category)

	if res.Error != nil {
		return nil, res.Error
	}
	return category, nil
}

func (db *categoryConnection) CreateCategory(ctx context.Context, category *entities.Category) (*entities.Category, error) {
	res := db.connection.WithContext(ctx).Create(&category)
	if res.Error != nil {
		return nil, res.Error
	}
	return category, nil
}
