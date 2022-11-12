package request

import "github.com/google/uuid"

type RequestCart struct {
	Quantity  int    `json:"quantity" form:"quantity" binding:"required"`
	UserID    string `json:"-" form:"user_id,omitempty"`
	ProductID string `json:"product_id" form:"product_id" binding:"required"`
}

type RequestCartUpdate struct {
	ID       uuid.UUID `json:"id" form:"id" binding:"required"`
	Quantity int       `json:"quantity" form:"quantity" binding:"required"`
	UserID   string    `json:"-" form:"user_id,omitempty"`
}
