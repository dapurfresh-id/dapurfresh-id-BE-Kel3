package request

type RequestRegister struct {
	Username string `json:"username" form:"username" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required" `
	Phone    string `json:"phone" binding:"required" `
	Password string `json:"password" form:"password" binding:"required"`
}
