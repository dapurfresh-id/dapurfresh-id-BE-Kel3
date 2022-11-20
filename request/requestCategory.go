package request

import (
	"github.com/google/uuid"
)

type RequestCreateCategory struct {
	ID      uuid.UUID `json:"-" form:"id,omitempty"`
	ImageID string    `json:"image_id" form:"image_id" binding:"required"`
	Name    string    `json:"name" form:"name" binding:"required"`
}
