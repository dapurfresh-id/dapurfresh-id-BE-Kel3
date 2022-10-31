package entities

import "github.com/google/uuid"

type OrderDetail struct {
	ID          uuid.UUID `gorm:"primaryKey" json:"id"`
	OrderID     string    `json:"category_id,omitempty"`
	Order       Order     `gorm:"foreignkey:CategoryID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"order"`
	ProductID   string    `json:"-"`
	Product     Product   `gorm:"foreignkey:ProductID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"products"`
	ProductName string    `json:"product_name"`
	SubTotal    int       `json:"sub_total"`
	Cost        int       `json:"cost"`
	Total       int       `json:"total"`
	Cash        int       `json:"cash"`
}
