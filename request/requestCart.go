package request

type RequestCart struct {
	Quantity  int    `json:"quantity" form:"quantity" binding:"required"`
	UserID    string `json:"-" form:"user_id,omitempty"`
	ProductID string `json:"product_id" form:"product_id" binding:"required"`
}
