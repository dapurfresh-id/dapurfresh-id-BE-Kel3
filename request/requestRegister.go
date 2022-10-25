package request

type RequestRegister struct {
	Username string `json:"username" form:"username" binding:"required,min=6"`
	Name     string `json:"name" form:"name" binding:"required,min=6" `
	Phone    string `json:"phone" binding:"required,max=12" `
	Password string `json:"password" form:"password" binding:"required,min=6"`
}
