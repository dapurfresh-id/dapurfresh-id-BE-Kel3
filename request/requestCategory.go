package request

import (
	"github.com/google/uuid"
)

type RequestCreateCategory struct {
	ID   uuid.UUID `json:"-" form:"id,omitempty" `
	Name string    `json:"name" form:"name" binding:"required"`
}
