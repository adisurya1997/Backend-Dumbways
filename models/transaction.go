package models

import "time"

type Transaction struct {
	ID             int                	     `json:"id" gorm:"primary_key:auto_increment"`
	ProductID 	   int                  	 `json:"product_id"`
	Product        ProductResponse      	 `json:"product"`
	ToppingID 	   int                  	 `json:"topping_id"`
	Topping        ToppingOrderResponse      `json:"topping"`
	BuyerID   	   int                 		 `json:"buyer_id"`
	Buyer     	   UsersProfileResponse 	 `json:"buyer"`
	TotalPrice     int               	     `json:"price"`
	Status         string           	     `json:"status"  gorm:"type:varchar(25)"`
	CreatedAt      time.Time          	     `json:"-"`
	UpdatedAt      time.Time            	 `json:"-"`
}