package entities

import "github.com/google/uuid"

type Category struct {
	ID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name string    `gorm:"type:varchar(255)" json:"name"`
}
