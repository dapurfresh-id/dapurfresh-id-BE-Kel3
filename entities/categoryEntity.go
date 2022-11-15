package entities

import "github.com/google/uuid"

type Category struct {
	ID    uuid.UUID `gorm:"primaryKey" json:"id"`
	Name  string    `gorm:"type:varchar(255)" json:"name"`
	Image string    `json:"image,omitempty"`
}
