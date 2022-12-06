package models

import "time"

type Profile struct {
	ID        int                 `json:"id" gorm:"primary_key:auto_increment"`
	Image     string              `json:"image" gorm:"type: varchar(255)"`
	Phone     int                 `json:"phone" gorm:"type: varchar(255)"`
	Address   string              `json:"address" gorm:"type: text"`
	UserID    int                 `json:"user_id"`
	User      UserProfileResponse `json:"user"`
	CreatedAt time.Time           `json:"-"`
	UpdatedAt time.Time           `json:"-"`
}

//for association relation with another table (user)
type ProfileResponse struct {
	Image   string `json:"image"`
	Phone   int    `json:"phone"`
	Address string `json:"address"`
	UserID  int    `json:"-"`
}

func (ProfileResponse) TableName() string {
	return "profiles"
}
