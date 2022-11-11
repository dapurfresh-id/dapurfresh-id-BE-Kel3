package request

type RequestProduct struct {
	Name       string `json:"name" form:"name" binding:"required"`
	Price      int    `json:"price" form:"price" binding:"required"`
	Unit       int    `json:"unit" form:"unit" binding:"required"`
	UnitType   string `json:"unittype" form:"unittype" binding:"required"`
	Image      string `json:"image" form:"image" binding:"required"`
	CategoryID string `json:"categoryid" form:"categoryid" binding:"required"`
	Arebuyed   int    `json:"arebuyed" form:"arebuyed"`
}
