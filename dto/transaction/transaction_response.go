package transactiondto

import "waysbuck/models"

type TransactionResponse struct {
	ID      int                         `json:"id" gorm:"primary_key:auto_increment"`
	UserID  int                         `json:"user_id"`
	User    models.UsersProfileResponse `json:"user"`
	Cart    []models.CartResponse       `json:"order"`
}
