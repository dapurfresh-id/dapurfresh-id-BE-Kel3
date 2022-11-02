package request

import "mime/multipart"

type RequestUserUpdate struct {
	ID       string                `json:"-" form:"id,omitempty" `
	Username string                `json:"username" form:"username"`
	Name     string                `json:"name" form:"name" `
	Phone    string                `json:"phone" form:"phone" `
	Password string                `json:"password,omitempty" form:"password,omitempty" `
	Image    *multipart.FileHeader `json:"image" form:"image"`
}
