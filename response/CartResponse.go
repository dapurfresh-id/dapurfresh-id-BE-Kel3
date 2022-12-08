package response

import (
	"time"

	"github.com/aldisaputra17/dapur-fresh-id/entities"
	"github.com/google/uuid"
)

type CartResponse struct {
	ID        uuid.UUID       `json:"id"`
	Name      string          `json:"name"`
	SubTotal  int             `json:"sub_total"`
	Price     int             `json:"price"`
	Quantity  int             `json:"quantity"`
	Unit      int             `json:"unit"`
	UnitType  string          `json:"unit_type"`
	ProductID string          `json:"product_id"`
	CreatedAt time.Time       `json:"created_at"`
	Products  ProductResponse `json:"products"`
}

func NewCartResponse(cart *entities.Cart) CartResponse {
	return CartResponse{
		ID:        cart.ID,
		Name:      cart.Name,
		SubTotal:  cart.SubTotal,
		Price:     cart.Price,
		Unit:      cart.Unit,
		UnitType:  cart.UnitType,
		Quantity:  cart.Quantity,
		CreatedAt: cart.CreatedAt,
		ProductID: cart.ProductID,
		Products:  NewProductResponse(&cart.Products),
	}
}
