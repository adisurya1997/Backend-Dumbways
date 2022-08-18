package models

import "time"

type Cart struct {
	ID            int                   `json:"id" gorm:"primary_key:auto_increment"`
	SubAmount 	  int                   `json:"subamount" form:"subamount" gorm:"type: int"`
	Qty        	  int                   `json:"qty" form:"qty"`
	TransactionID int 					`json:"transaction_id" form:"transaction_id"`
	Transaction   TransactionResponse 	`json:"transaction" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductID     int                   `json:"product_id"`
	Product   	  ProductResponse       `json:"product" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Topping    	  []Topping             `json:"topping" gorm:"many2many:cart_topping"`
	ToppingID     []int                 `json:"-" form:"topping_id" gorm:"-"`
	Status        string           	    `json:"status" gorm:"type:varchar(25)"`
	CreatedAt     time.Time             `json:"-"`
	UpdatedAt     time.Time             `json:"-"`
	// UserID     int                   `json:"user_id" form:"user_id"`
	// User       UsersProfileResponse  `json:"user"`
}

type CartResponse struct {
	ID            int                   `json:"id"` 
	SubAmount     int                   `json:"subamount"`
	Qty        	  int                   `json:"qty"`
	TransactionID int 					`"-"`
	Transaction   TransactionResponse   `json:"transaction"  gorm:"foreignKey:TransactionID"`
	ProductID     int                   `json:"product_id"`
	Product       ProductResponse       `json:"product"`
	Topping    	  []Topping             `json:"topping" gorm:"many2many:cart_topping"`
	ToppingID     []int                 `json:"-" form:"topping_id" gorm:"-"`
	Status        string           	    `json:"status"`
}

func (CartResponse) TableName() string {
	return "carts"
}