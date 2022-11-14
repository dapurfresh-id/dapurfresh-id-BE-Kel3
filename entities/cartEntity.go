package entities

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	ID        uuid.UUID `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255)" json:"name"`
	SubTotal  int       `json:"sub_total"`
	Price     int       `gorm:"not null" json:"price"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Products  Product   `gorm:"foreignkey:ProductID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"products"`
	User      User      `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	UserID    string    `gorm:"not null" json:"-"`
	ProductID string    `gorm:"not null" json:"-"`
}
