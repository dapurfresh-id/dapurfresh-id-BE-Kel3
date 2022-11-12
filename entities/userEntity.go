package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"primaryKey" json:"id"`
	Username string    `gorm:"uniqueIndex;type:varchar(255)" json:"username"`
	Name     string    `gorm:"type:varchar(255)" json:"name" validate:"required, max=100"`
	Password string    `gorm:"->;<-;not null" json:"-" validate:"required, min=6"`
	Phone    string    `gorm:"type:varchar(255)" json:"phone" validate:"required, min=11"`
	// ImageID   string    `json:"image_id"`
	Image     string    `json:"file,omitempty"`
	Token     string    `gorm:"-" json:"token,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Cart      *Cart     `json:"cart,omitempty"`
	Order     *Order    `json:"order,omitempty"`
}
