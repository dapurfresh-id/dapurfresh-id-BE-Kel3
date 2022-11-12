package request

import (
	"mime/multipart"
)

type ImageRequest struct {
	File multipart.File `json:"file,omitempty" validate:"required" form:"file"`
}
