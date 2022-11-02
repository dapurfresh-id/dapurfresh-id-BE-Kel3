package entities

import (
	"time"

	"github.com/google/uuid"
)

const ()

type Order struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Catatan   string    `gorm:"type:varchar(255)" json:"catatan" validate:"required, max=255"`
	Address   string    `gorm:"type:varchar(255)" json:"address" validate:"required, max=100"`
	Status    string    `gorm:"not null;default:proses" json:"status"`
	Cost      int       `json:"cost"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
	UserID    string    `gorm:"not null" json:"user_id"`
	User      User      `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	CartID    string    `gorm:"not null" json:"-"`
	Carts     Cart      `gorm:"foreignkey:CartID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"cart"`
}
