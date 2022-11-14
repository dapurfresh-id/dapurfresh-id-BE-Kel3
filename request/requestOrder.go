package request

import "github.com/google/uuid"

type RequestOrderCreate struct {
	ID      uuid.UUID `json:"-" form:"id,omitempty" `
	Noted   string    `json:"noted" validate:"required" form:"noted" bind:"required"`
	Address string    `json:"address" validate:"required" form:"address" bind:"required"`
	Cost    int       `json:"cost" form:"cost"`
	UserID  string    `json:"-" form:"user_id,omitempty"`
	CartID  string    `json:"cart_id" form:"cart_id,omitempty" binding:"required"`
}

type RequestPatchOrder struct {
	ID     uuid.UUID `json:"id" form:"id" binding:"required"`
	Status string    `json:"status" form:"status"`
	UserID string    `json:"-" form:"user_id,omitempty"`
}
