package request

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type RequestImgUpdate struct {
	ID    uuid.UUID         `json:"-" form:"id,omitempty"`
	Image multipart.File `json:"image,omitempty" form:"image" validate:"required"`
}
