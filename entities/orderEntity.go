package entities

import "time"

type Order struct {
	ID          string    `gorm:"primary_key:auto_increment" json:"id"`
	Catatan     string    `gorm:"type:varchar(255)" json:"catatan" validate:"required, max=255"`
	Address     string    `gorm:"type:varchar(255)" json:"address" validate:"required, max=100"`
	Status      string    `gorm:"not null;default:proses" json:"status"`
	OngkosKirim int       `json:"ongkos_kirim"`
	CreatedAt   time.Time `json:"created_at"`
	DeletedAt   time.Time `json:"deleted_at"`
	UserID      string    `gorm:"not null" json:"-"`
	User        User      `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	CartID      string    `gorm:"not null" json:"-"`
	Carts       Cart      `gorm:"foreignkey:CartID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"cart"`
}
