package request

type RequestCheckoutProduct struct {
	Arebuyed int `json:"are_buyed" form:"are_buyed" binding:"required"`
}
