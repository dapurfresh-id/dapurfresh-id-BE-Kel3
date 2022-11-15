package entities

import (
	"github.com/google/uuid"
)

type Product struct {
	ID       uuid.UUID `gorm:"primaryKey" json:"id"`
	Name     string    `gorm:"type:varchar(255)" json:"name"`
	Price    int       `json:"price"`
	Unit     int       `json:"unit"`
	UnitType string    `gorm:"not null;default:gr" json:"unit_type"`
	// ImageID    string    `json:"image_id"`
	Image      string   `json:"image,omitempty" form:"image" validate:"required"`
	CategoryID string   `json:"category_id,omitempty"`
	Categories Category `gorm:"foreignkey:CategoryID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"categories"`
	AreBuyed   int      `gorm:"not null;default:0" json:"are_buyed"`
}
