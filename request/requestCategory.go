package request

type RequestCategory struct {
	Name  string `json:"name" form:"name" binding:"required"`
	Image string `json:"image" form:"name" binding:"required"`
}
