package response

import (
	"github.com/aldisaputra17/dapur-fresh-id/entities"
	"github.com/google/uuid"
)

type ProductResponse struct {
	ID    uuid.UUID     `gorm:"primaryKey" json:"id"`
	Image ImageResponse `json:"image"`
}

type ImageResponse struct {
	ID   string `json:"id"`
	File string `json:"file,omitempty"`
}

func NewImgResponse(img *entities.Image) ImageResponse {
	return ImageResponse{
		ID:   img.ID,
		File: img.File,
	}
}

func NewProductResponse(prod *entities.Product) ProductResponse {
	return ProductResponse{
		ID:    prod.ID,
		Image: NewImgResponse(&prod.Images),
	}
}
