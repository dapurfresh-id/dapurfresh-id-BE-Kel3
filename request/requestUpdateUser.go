package request

import (
	"github.com/google/uuid"
	// "mime/multipart"
)

type RequestUpdateUser struct {
	ID       uuid.UUID `json:"id" form:"id"`
	Username string    `json:"username" form:"username" binding:"required,min=6"`
	Name     string    `json:"name" form:"name" binding:"required,min=6"`
	Password string    `json:"password" form:"password" binding:"required,min=6"`
	Phone    string    `json:"phone" binding:"required,max=12"`
	// Image    string    `json:"image" form:"image"`
	// Image *multipart.FileHeader `form:"image"`
}
