package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"uniqueIndex;type:varchar(255)" json:"username"`
	Name      string    `gorm:"type:varchar(255)" json:"name" validate:"required, max=100"`
	Password  string    `gorm:"->;<-;not null" json:"-" validate:"required, min=6"`
	Phone     string    `gorm:"type:varchar(255)" json:"phone" validate:"required, min=11"`
	Image     string    `gorm:"default:null" json:"image,omitempty" form:"image" validate:"required"`
	Token     string    `gorm:"-" json:"token,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Cart      *Cart     `json:"cart,omitempty"`
}
