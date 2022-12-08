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
	PaginationProduct(pagination *entities.Pagination) (helpers.PaginationResult, int)
	UpdateBuy(ctx context.Context, prod *entities.Product) (*entities.Product, error)
	PopularProduct(ctx context.Context) (*[]entities.Product, error)
	Create(ctx context.Context, product *entities.Product) (*entities.Product, error)
	Update(ctx context.Context, cartID string) (*entities.Product, error)
}

type productConnection struct {
	connection *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productConnection{
		connection: db,
	}
}

func (db *productConnection) Create(ctx context.Context, product *entities.Product) (*entities.Product, error) {
	res := db.connection.WithContext(ctx).Create(&product)
	if res.Error != nil {
		return nil, res.Error
	}
	return product, nil
}

func (db *productConnection) FindAllProduct(ctx context.Context) (*[]entities.Product, error) {
	var product *[]entities.Product
	res := db.connection.WithContext(ctx).Preload("Categories.Image").Preload("Images").Find(&product)
	if res.Error != nil {
		return nil, res.Error
	}
	return product, nil
}

func (db *productConnection) FindProductById(ctx context.Context, productId string) (*entities.Product, error) {
	var product *entities.Product
	res := db.connection.WithContext(ctx).Where("id = ?", productId).Preload("Images").Preload("Categories.Image").First(&product)

	if res.Error != nil {
		return nil, res.Error
	}
	return product, nil
}

func (db *productConnection) Update(ctx context.Context, cartID string) (*entities.Product, error) {
	var carts *entities.Cart
	var products *entities.Product
	cart := db.connection.WithContext(ctx).Where("id = ?", cartID).Preload("Products").Find(&carts)
	product := db.connection.WithContext(ctx).Where("id = ?", carts.ProductID).Preload("Categories").Preload("Images").First(&products)

	if cart.Error != nil {
		return nil, cart.Error
	}
	if product.Error != nil {
		return nil, cart.Error
	}
	query := `UPDATE products set are_buyed = ? WHERE id = ?`
	stmt, err := db.connection.ConnPool.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	res, err := stmt.ExecContext(ctx, carts.Quantity+products.AreBuyed, carts.ID)
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
	return products, nil
}

func (db *productConnection) FindProductByCategory(ctx context.Context, categoryId string) (*[]entities.Product, error) {
	var product *[]entities.Product
	res := db.connection.WithContext(ctx).Where("category_id = ?", categoryId).Preload("Categories.Image").Preload("Images").Find(&product)

	if res.Error != nil {
		return nil, res.Error
	}
	return product, nil
}

func (db *productConnection) PopularProduct(ctx context.Context) (*[]entities.Product, error) {
	var product *[]entities.Product
	res := db.connection.WithContext(ctx).Order("are_buyed desc").Preload("Images").Preload("Categories.Image").Find(&product)
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

	find := db.connection.Limit(pagination.Limit).Offset(offset).Order(pagination.SortProduct)

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

	find = find.Preload("Categories.Image").Preload("Images").Find(&prod)

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

func (db *productConnection) UpdateBuy(ctx context.Context, prod *entities.Product) (*entities.Product, error) {
	cartRp := NewCartRepository(db.connection)
	cartItem, _ := cartRp.GetCartProd(ctx)
	prod.AreBuyed = cartItem.Quantity
	res := db.connection.WithContext(ctx).Model(&prod).Updates(entities.Product{
		AreBuyed: prod.AreBuyed,
	})
	if res.Error != nil {
		return nil, res.Error
	}
	return prod, nil
}
