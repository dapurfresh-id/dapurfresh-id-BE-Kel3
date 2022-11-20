package entities

import "github.com/google/uuid"

type Category struct {
	ID      uuid.UUID `gorm:"primaryKey" json:"id"`
	Name    string    `gorm:"type:varchar(255)" json:"name"`
	ImageID string    `json:"image_id"`
	Image   Image     `gorm:"foreignkey:ImageID;constraint:onUpdate:CASCADE,onDelete:SET NULL" json:"image"`
}
