package repositories

import (
	"context"
	"fmt"
	"math"
	"strings"

	"github.com/aldisaputra17/dapur-fresh-id/entities"
	"github.com/aldisaputra17/dapur-fresh-id/helpers"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAllProduct(ctx context.Context) (*[]entities.Product, error)
	FindProductById(ctx context.Context, productId string) (*entities.Product, error)
	FindProductByCategory(ctx context.Context, categoryId string) (*[]entities.Product, error)
	FindProductByNameEqual(ctx context.Context, name string) (*entities.Product, error)
	FindProductByNameContains(ctx context.Context, name string) (*entities.Product, error)
	FindProductByNameLike(ctx context.Context, name string) (*entities.Product, error)
	LimitProduct(ctx context.Context, limit int) (*[]entities.Product, error)
	PaginationProduct(pagination *entities.Pagination) (helpers.PaginationResult, int)
	PopularProduct(ctx context.Context) (*[]entities.Product, error)
}

type productConnection struct {
	connection *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productConnection{
		connection: db,
	}
}

func (db *productConnection) FindAllProduct(ctx context.Context) (*[]entities.Product, error) {
	var product *[]entities.Product
	res := db.connection.WithContext(ctx).Preload("Categories").Find(&product)
	if res.Error != nil {
		return nil, res.Error
	}
	return product, nil
}

func (db *productConnection) FindProductById(ctx context.Context, productId string) (*entities.Product, error) {
	var product *entities.Product
	res := db.connection.WithContext(ctx).Where("id = ?", productId).Preload("Categories").First(&product)

	if res.Error != nil {
		return nil, res.Error
	}
	return product, nil
}
func (db *productConnection) FindProductByCategory(ctx context.Context, categoryId string) (*[]entities.Product, error) {
	var product *[]entities.Product
	res := db.connection.WithContext(ctx).Where("category_id = ?", categoryId).Preload("Categories").Find(&product)

	if res.Error != nil {
		return nil, res.Error
	}
	return product, nil
}

func (db *productConnection) FindProductByNameEqual(ctx context.Context, name string) (*entities.Product, error) {
	var product *entities.Product
	res := db.connection.WithContext(ctx).Where("name = ?", name).Preload("Categories").First(&product)
	if res.Error != nil {
		return nil, res.Error
	}
	return product, nil
}

func (db *productConnection) FindProductByNameLike(ctx context.Context, name string) (*entities.Product, error) {
	var product *entities.Product
	stringLike := "%" + name + "%"
	res := db.connection.WithContext(ctx).Where("name LIKE ?", stringLike).Preload("Categories").First(&product)
	if res.Error != nil {
		return nil, res.Error
	}
	return product, nil
}

func (db *productConnection) FindProductByNameContains(ctx context.Context, name string) (*entities.Product, error) {
	var product *entities.Product
	res := db.connection.WithContext(ctx).Where("name IN ?", []string{name}).Preload("Categories").First(&product)
	if res.Error != nil {
		return nil, res.Error
	}
	return product, nil
}

func (db *productConnection) PopularProduct(ctx context.Context) (*[]entities.Product, error) {
	var product *[]entities.Product
	res := db.connection.WithContext(ctx).Preload("Categories").Order("are_buyed desc").Order("name asc").Find(&product)
	if res.Error != nil {
		return nil, res.Error
	}
	return product, nil
}

func (db *productConnection) LimitProduct(ctx context.Context, limit int) (*[]entities.Product, error) {
	var product *[]entities.Product
	res := db.connection.WithContext(ctx).Limit(limit).Preload("Categories").Find(&product)
	if res.Error != nil {
		return nil, res.Error
	}
	return product, nil
}

func (db *productConnection) PaginationProduct(pagination *entities.Pagination) (helpers.PaginationResult, int) {
	var prod []entities.Product

	var (
		totalRows  int64
		totalPages int
		fromRow    int
		toRow      int
	)

	offset := pagination.Page * pagination.Limit

	find := db.connection.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)

	searchs := pagination.Searchs

	for _, value := range searchs {
		column := value.Column
		action := value.Action
		query := value.Query

		switch action {
		case "equals":
			whereQuery := fmt.Sprintf("%s = ?", column)
			find = find.Where(whereQuery, query)
		case "contains":
			whereQuery := fmt.Sprintf("%s LIKE ?", column)
			find = find.Where(whereQuery, "%"+query+"%")
		case "in":
			whereQuery := fmt.Sprintf("%s IN (?)", column)
			queryArray := strings.Split(query, ",")
			find = find.Where(whereQuery, queryArray)
		}
	}

	find = find.Preload("Categories").Find(&prod)

	errFind := find.Error

	if errFind != nil {
		return helpers.PaginationResult{Error: errFind}, totalPages
	}

	pagination.Rows = prod

	errCount := db.connection.Model(&entities.Product{}).Count(&totalRows).Error

	if errCount != nil {
		return helpers.PaginationResult{Error: errCount}, totalPages
	}

	pagination.TotalRows = int(totalRows)

	totalPages = int(math.Ceil(float64(totalRows)/float64(pagination.Limit))) - 1

	if pagination.Page == 0 {
		fromRow = 1
		toRow = pagination.Limit
	} else {
		if pagination.Page <= totalPages {
			fromRow = pagination.Page*pagination.Limit + 1
			toRow = (pagination.Page + 1) * pagination.Limit
		}
	}

	if toRow > int(totalRows) {
		toRow = int(totalRows)
	}

	pagination.FromRow = fromRow
	pagination.ToRow = toRow

	return helpers.PaginationResult{Result: pagination}, totalPages
}
