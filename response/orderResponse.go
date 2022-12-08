package response

import (
	"time"

	"github.com/aldisaputra17/dapur-fresh-id/entities"
	"github.com/google/uuid"
)

type OrderResponse struct {
	ID        uuid.UUID    `json:"id"`
	Name      string       `json:"name"`
	Catatan   string       `json:"catatan"`
	Address   string       `json:"address"`
	Status    string       `json:"status"`
	Total     int          `json:"total"`
	SubTotal  int          `json:"sub_total"`
	Cost      int          `json:"cost"`
	CreatedAt time.Time    `json:"created_at"`
	Carts     CartResponse `json:"cart"`
}

type GetOrderResponse struct {
	ID        uuid.UUID     `json:"id"`
	Name      string        `json:"name"`
	Catatan   string        `json:"catatan"`
	Address   string        `json:"address"`
	Status    string        `json:"status"`
	Total     int           `json:"total"`
	SubTotal  int           `json:"sub_total"`
	Cost      int           `json:"cost"`
	CreatedAt time.Time     `json:"created_at"`
	UserID    string        `gorm:"not null" json:"user_id"`
	User      entities.User `json:"user"`
	CartID    string        `gorm:"default:null" json:"cart_id"`
	Carts     entities.Cart `json:"cart"`
}

func NewGetOrderResponse(order *entities.Order) GetOrderResponse {
	return GetOrderResponse{
		ID:        order.ID,
		Name:      order.Name,
		Catatan:   order.Catatan,
		Address:   order.Address,
		Status:    order.Status,
		Total:     order.Total,
		SubTotal:  order.SubTotal,
		Cost:      order.Cost,
		CreatedAt: order.CreatedAt,
		User: entities.User{
			Name:  order.User.Name,
			Phone: order.User.Phone,
		},
		Carts: entities.Cart{
			ID:       order.Carts.ID,
			Unit:     order.Carts.Unit,
			UnitType: order.Carts.UnitType,
			Products: order.Carts.Products,
		},
	}
}

func CreateOrderResponse(order *entities.Order) OrderResponse {
	return OrderResponse{
		ID:        order.ID,
		Name:      order.Name,
		Catatan:   order.Catatan,
		Address:   order.Address,
		Status:    order.Status,
		Total:     order.Total,
		SubTotal:  order.SubTotal,
		Cost:      order.Cost,
		CreatedAt: order.CreatedAt,
		Carts:     NewCartResponse(&order.Carts),
	}
}
