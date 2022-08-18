package cartdto

type CartRequest struct {
	SubAmount	 	int	`json:"subamount" form:"subamount"`
	ProductID    	int    `json:"product_id" form:"product_id"`
	ToppingID   	[]int  `json:"category_id" form:"category_id" gorm:"type: int"`
	TransactionID   int    `json:"transaction_id" form:"transaction_id"`
}