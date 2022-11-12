package request

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type RequestUserUpdate struct {
	ID       uuid.UUID             `json:"-" form:"id,omitempty"`
	Username string                `json:"username" form:"username" bind:"required"`
	Name     string                `json:"name" form:"name" bind:"required"`
	Phone    string                `json:"phone" form:"phone" bind:"required"`
	Password string                `json:"password,omitempty" form:"password,omitempty" bind:"required" `
	Image    *multipart.FileHeader `json:"image,omitempty" validate:"required" form:"image" bind:"required"`
}
