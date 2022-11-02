package entities

import (
	"time"
)

type User struct {
	ID        string    `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Username  string    `gorm:"uniqueIndex;type:varchar(255)" json:"username"`
	Name      string    `gorm:"type:varchar(255)" json:"name" validate:"required, max=100"`
	Password  string    `gorm:"->;<-;not null" json:"-" validate:"required, min=6"`
	Phone     string    `gorm:"type:varchar(255)" json:"phone" validate:"required, min=11"`
	Image     string    `json:"image"`
	Token     string    `gorm:"-" json:"token,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Cart      *Cart     `json:"cart,omitempty"`
	Order     *Order    `json:"order,omitempty"`
}
