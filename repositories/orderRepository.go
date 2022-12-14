package repositories

import (
	"context"
	"fmt"
	"math"
	"strings"

	"github.com/aldisaputra17/dapur-fresh-id/entities"
	"github.com/aldisaputra17/dapur-fresh-id/helpers"
	"github.com/aldisaputra17/dapur-fresh-id/response"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(ctx context.Context, order *entities.Order) (*response.OrderResponse, error)
	GetOrderID(ctx context.Context, CartID string, UserID string) ([]*entities.Order, error)
	GetOrder(ctx context.Context, paginate *entities.Pagination, userID string) (helpers.PaginationResult, int)
	GetDetail(ctx context.Context, userID string, id string) (*entities.Order, error)
	PatchStatus(ctx context.Context, order *entities.Order) (*entities.Order, error)
}

type orderConnection struct {
	connection *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderConnection{
		connection: db,
	}
}

func (db *orderConnection) GetOrderID(ctx context.Context, CartID string, UserID string) ([]*entities.Order, error) {
	var order []*entities.Order
	res := db.connection.WithContext(ctx).Where("cart_id = ? and user_id = ?", CartID, UserID).Find(&order)
	if res.Error != nil {
		return nil, res.Error
	}
	return order, nil
}

func (db *orderConnection) Create(ctx context.Context, order *entities.Order) (*response.OrderResponse, error) {
	CartRp := NewCartRepository(db.connection)

	cartItem, err := CartRp.GetCartByid(ctx, order.CartID)
	if err != nil {
		return nil, err
	}

	if len(cartItem) <= 0 {
		lenErr := fmt.Errorf("cart dosent exists")
		return nil, lenErr
	}
	order.Name = cartItem[0].Name
	order.UserID = cartItem[0].UserID
	order.SubTotal = cartItem[0].SubTotal
	order.Total = cartItem[0].SubTotal + order.Cost
	order.Unit = cartItem[0].Unit
	order.UnitType = cartItem[0].UnitType
	fmt.Println("model:", order.CartID)

	orderItem, err := db.GetOrderID(ctx, order.CartID, order.UserID)
	if err != nil {
		return nil, err
	}

	if len(orderItem) < 0 {
		fmt.Println("order dosent exists")
		return nil, fmt.Errorf("order not foud, please add order again")
	}
	res := db.connection.WithContext(ctx).Save(&order)
	if res.Error != nil {
		return nil, res.Error
	}

	if cartItem[0].UserID == order.UserID {
		CartRp.Trancate(ctx, order.UserID, order.CartID)
	}
	rsp := &response.OrderResponse{
		ID:        order.ID,
		Name:      order.Name,
		Catatan:   order.Catatan,
		Address:   order.Address,
		Status:    order.Status,
		Total:     order.Total,
		SubTotal:  order.SubTotal,
		Cost:      order.Cost,
		CreatedAt: order.CreatedAt,
		Carts:     response.NewCartResponse(cartItem[0]),
	}
	return rsp, nil
}

func (db *orderConnection) GetOrder(ctx context.Context, paginate *entities.Pagination, userID string) (helpers.PaginationResult, int) {
	var (
		order      []*entities.Order
		totalRows  int64
		totalPages int
		fromRow    int
		toRow      int
	)

	offset := paginate.Page * paginate.Limit

	find := db.connection.Limit(paginate.Limit).Offset(offset).Order(paginate.SortOrder)
	searchs := paginate.Searchs

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
	find = find.Where("user_id", userID).Preload("User").Find(&order)
	errFind := find.Error

	if errFind != nil {
		return helpers.PaginationResult{Error: errFind}, totalPages
	}

	paginate.Rows = order

	errCount := db.connection.Model(&entities.Order{}).Count(&totalRows).Error

	if errCount != nil {
		return helpers.PaginationResult{Error: errCount}, totalPages
	}

	paginate.TotalRows = int(totalRows)

	totalPages = int(math.Ceil(float64(totalRows)/float64(paginate.Limit))) - 1

	if paginate.Page == 0 {
		fromRow = 1
		toRow = paginate.Limit
	} else {
		if paginate.Page <= totalPages {
			fromRow = paginate.Page*paginate.Limit + 1
			toRow = (paginate.Page + 1) * paginate.Limit
		}
	}

	if toRow > int(totalRows) {
		toRow = int(totalRows)
	}

	paginate.FromRow = fromRow
	paginate.ToRow = toRow

	return helpers.PaginationResult{Result: paginate}, totalPages
}

func (db *orderConnection) GetDetail(ctx context.Context, userID string, id string) (*entities.Order, error) {
	var order *entities.Order
	res := db.connection.WithContext(ctx).Where("user_id = ? and id = ?", userID, id).Preload("User").Preload("Carts.Products.Image").Find(&order)
	if res.Error != nil {
		return nil, res.Error
	}
	return order, nil
}

func (db *orderConnection) PatchStatus(ctx context.Context, order *entities.Order) (*entities.Order, error) {
	orderID := fmt.Sprintf("%v", order.ID)
	orderItem, err := db.GetDetail(ctx, order.UserID, orderID)
	if err != nil {
		return nil, err
	}
	orderItem.Status = "cancel"
	res := db.connection.WithContext(ctx).Model(&order).Where("id = ?", order.ID).Updates(entities.Order{
		Status: order.Status}).Preload("User").Preload("Carts.Products").Find(&order)
	if res.Error != nil {
		return nil, res.Error
	}
	return order, nil
}
