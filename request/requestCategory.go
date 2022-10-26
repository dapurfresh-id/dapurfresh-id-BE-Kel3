package request

type RequestCategory struct {
	Name string `json:"name" form:"name" binding:"required"`
}
