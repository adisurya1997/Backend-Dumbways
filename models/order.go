package models

import "time"

type Order struct {
	ID            int                 	 `json:"id" gorm:"primary_key:auto_increment"`
	ProductID     int                	 `json:"product_id"`
	ProductOrder  ProductOrderResponse 	 `json:"product"`
	ToppingID     []int                  `json:"topping_id"`
	ToppingOrder  ToppingOrderResponse 	 `json:"toppings"`
	CreatedAt     time.Time            	 `json:"-"`
	UpdatedAt     time.Time           	 `json:"-"`
}

type OrderResponse struct {
	ID         	  int                    `json:"id"`
	ProductID     int                	 `json:"-"`
	ProductOrder  ProductOrderResponse 	 `json:"product"`
	ToppingID     []int                  `json:"-"`
	ToppingOrder  ToppingOrderResponse 	 `json:"toppings"`
}

func (OrderResponse) TableName() string {
	return "orders"
}