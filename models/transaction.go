package models

import "time"

type Transaction struct {
	ID         int                   `json:"id" gorm:"primary_key:auto_increment"`
	Status     string                `json:"status"  gorm:"type:varchar(25)"`
	UserID     int                   `json:"user_id" form:"user_id"`
	Userorder  UsersProfileResponse  `json:"user_order" gorm:"foreignKey:UserID;references:ID`
	OrderID    []int                 `json:"order_id" form:"order_id" `
	Order      OrderResponse 		 `json:"order"`
	TotalPrice int             	 	 `json:"price"`
	CreatedAt  time.Time             `json:"-"`
	UpdatedAt  time.Time             `json:"-"`
}