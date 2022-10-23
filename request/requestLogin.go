package request

type RequestLogin struct {
	Username string `json:"username" form:"username" binding:"required,username"`
	Password string `json:"password" form:"password" binding:"required,password"`
}
