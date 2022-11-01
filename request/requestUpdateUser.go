package request

// import (
// 	"github.com/google/uuid"
// )
import (
	"mime/multipart"
)

type RequestUpdateUser struct {
	ID       uint64 `json:"id" form:"id"`
	UserName string `json:"userName" form:"userName" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Password string `json:"password,omitempty" form:"password,omitempty"`
	Phone    string `json:"phone" binding:"required,phone"`
	// Image    string `json:"image"`
	Image *multipart.FileHeader `form:"image" binding:"required"`
}
