package entities

type Image struct {
	ID   string `gorm:"primaryKey" json:"id"`
	File string `json:"file,omitempty" form:"file" validate:"required"`
}

// func (img *Image) BeforeCreate() error {
// 	img.ID = uuid.New().String()
// 	return nil
// }
