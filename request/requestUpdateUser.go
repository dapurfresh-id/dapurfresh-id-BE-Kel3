package request

import (
	"github.com/google/uuid"
)

type RequestUpdateUser struct {
	ID       uuid.UUID `json:"id" form:"id"`
	UserName string    `json:"userName" form:"userName" binding:"required"`
	Name     string    `json:"name" form:"name" binding:"required"`
	Password string    `json:"password,omitempty" form:"password,omitempty"`
	Phone    string    `json:"phone" binding:"required,phone"`
}
