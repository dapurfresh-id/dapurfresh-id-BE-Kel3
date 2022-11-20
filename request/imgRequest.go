package request

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type ImageRequest struct {
	ID uuid.UUID `json:"-" form:"id,omitempty"`
	File multipart.File `json:"file,omitempty" validate:"required" form:"file"`
}
