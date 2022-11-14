package request

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type ReqeustCreateProduct struct {
	ID         uuid.UUID             `json:"-" form:"id,omitempty" `
	Name       string                `json:"name" validate:"required" form:"name" bind:"required"`
	Price      int                   `json:"price" validate:"required" form:"price" bind:"required"`
	Unit       int                   `json:"unit" validate:"required" form:"unit" bind:"required"`
	UnitType   string                `json:"unit_type" validate:"required" form:"unit_type" bind:"required"`
	Image      *multipart.FileHeader `json:"image,omitempty" validate:"required" form:"image" binding:"required"`
	CategoryID string                `json:"category_id" validate:"required" form:"category_id" bind:"required"`
}
