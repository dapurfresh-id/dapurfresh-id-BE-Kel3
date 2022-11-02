package entities

import "github.com/google/uuid"

type OrderDetail struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	OrderID     string    `json:"order_id,omitempty"`
	Orders      Order     `gorm:"foreignkey:OrderID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"order"`
	ProductID   string    `json:"-"`
	Product     Product   `gorm:"foreignkey:ProductID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"products"`
	ProductName string    `json:"product_name"`
	SubTotal    int       `json:"sub_total"`
	Cost        int       `json:"cost"`
	Total       int       `json:"total"`
	Cash        int       `json:"cash"`
}
