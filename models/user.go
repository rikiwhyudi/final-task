package models

import "time"

type User struct {
	ID int `json:"id"`
	// TransactionID *int                `json:"transaction_id"`
	// AccountNumber string          `json:"account_number" gorm:"type: varchar(255)"`
	// Image         string          `json:"image" gorm:"type: varchar(255)"`
	// Profile       ProfileResponse `json:"profile"`
	// Remaining     string          `json:"remaining" gorm:"type: varchar(255)"`
	Transaction  []Transaction `json:"transaction"`
	Name         string        `json:"name" gorm:"type: varchar(255)"`
	Email        string        `json:"email" gorm:"type: varchar(255)"`
	Password     string        `json:"-" gorm:"type: varchar(255)"`
	Status       string        `json:"status" gorm:"type: varchar(255)"`
	Subscription string        `json:"subscription" gorm:"type: varchar(255)"`
	CreatedAt    time.Time     `json:"-"`
	UpdatedAt    time.Time     `json:"-"`
}

type UserProfileResponse struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email" gorm:"type: varchar(255)"`
	Status       string `json:"status"`
	Subscription string `json:"subscription"`
}

func (UserProfileResponse) TableName() string {
	return "users"
}
