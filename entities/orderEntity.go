package entities

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID        uuid.UUID `gorm:"primaryKey" json:"id"`
	Catatan   string    `gorm:"type:varchar(255)" json:"catatan" validate:"required, max=255"`
	Address   string    `gorm:"type:varchar(255)" json:"address" validate:"required, max=100"`
	Status    string    `gorm:"not null;default:proses" json:"status"`
	Total     int       `json:"total"`
	SubTotal  int       `json:"sub_total"`
	Cost      int       `json:"cost"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
	UserID    string    `gorm:"not null" json:"user_id"`
	User      User      `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	CartID    string    `gorm:"default:null" json:"cart_id"`
	Carts     *Cart     `gorm:"foreignkey:CartID" json:"carts"`
}
