package entities

import "time"

type Cart struct {
	ID         string    `gorm:"primary_key:auto_increment" json:"id"`
	Name       string    `gorm:"type:varchar(255)" json:"name"`
	TotalPrice int       `json:"total_price"`
	Price      int       `gorm:"not null" json:"price"`
	Quantity   int       `gorm:"not null" json:"quantity"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	UserID     string    `gorm:"not null" json:"-"`
	User       User      `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	ProductID  string    `gorm:"not null" json:"-"`
	Products   Product   `gorm:"foreignkey:ProductID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"products"`
}
