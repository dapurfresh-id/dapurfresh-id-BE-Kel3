package entities

import (
	"github.com/google/uuid"
)

type Image struct {
	ID   uuid.UUID `gorm:"primaryKey" json:"id"`
	File string `json:"file,omitempty" form:"file" validate:"required"`
}
