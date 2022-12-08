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
	Unit      int       `json:"unit"`
	UnitType  string    `json:"unit_type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ProductID string    `gorm:"not null" json:"product_id"`
	Products  Product   `gorm:"foreignkey:ProductID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"products"`
	UserID    string    `gorm:"not null" json:"user_id"`
	User      User      `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
