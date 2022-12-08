package models

import "time"

type Transaction struct {
	ID            int                 `json:"id" gorm:"primary_key:auto_increment"`
	UserID        int                 `json:"user_id"`
	SubscribeEnd  time.Time           `json:"subscribe_end"`
	User          UserProfileResponse `json:"user"`
	AccountNumber []User              `json:"account_number"`
	Status        string              `json:"status"`
	Amount        int                 `json:"ammount"`
	UpdatedAt     time.Time           `json:"updated_at"`
}

type TransactionResponse struct {
	ID     int
	UserID int
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
