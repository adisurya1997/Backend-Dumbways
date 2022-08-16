package models

import "time"

type Transaction struct {
	ID             int                	     `json:"id" gorm:"primary_key:auto_increment"`
	BuyerID   	   int                 		 `json:"buyer_id"`
	Buyer     	   UsersProfileResponse 	 `json:"buyer"`
	Carts          []CartResponse            `json:"carts"`
	Amount    	   int               	     `json:"Amount"`
	Status         string           	     `json:"status"  gorm:"type:varchar(25)"`
	CreatedAt      time.Time          	     `json:"-"`
	UpdatedAt      time.Time            	 `json:"-"`
}


type TransactionResponse struct {
	ID   	 int    `json:"id"`
	BuyerID  int    `json:"buyer_id"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}