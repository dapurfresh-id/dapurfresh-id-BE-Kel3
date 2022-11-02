package entities

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name       string    `gorm:"type:varchar(255)" json:"name"`
	TotalPrice int       `json:"total_price"`
	Price      int       `gorm:"not null" json:"price"`
	Quantity   int       `gorm:"not null" json:"quantity"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	UserID     string    `gorm:"not null" json:"user_id"`
	User       User      `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	ProductID  string    `gorm:"not null" json:"-"`
	Products   Product   `gorm:"foreignkey:ProductID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
}
