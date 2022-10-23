package entities

type Category struct {
	ID   string `gorm:"primary_key:auto_increment" json:"id"`
	Name string `gorm:"type:varchar(255)" json:"name"`
}