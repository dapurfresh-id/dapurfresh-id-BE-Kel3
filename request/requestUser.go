package request

import (
	"github.com/google/uuid"
)

type RequestUserUpdate struct {
	ID       uuid.UUID `json:"-" form:"id,omitempty"`
	Username string    `json:"username" form:"username" bind:"required"`
	Name     string    `json:"name" form:"name" bind:"required"`
	Phone    string    `json:"phone" form:"phone" bind:"required"`
	Password string    `json:"password,omitempty" form:"password,omitempty" bind:"required" `
	ImageID  string    `json:"image_id" form:"image_id" binding:"required"`
}
