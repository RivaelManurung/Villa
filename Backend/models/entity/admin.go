package entity

import "time"

type Admin struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username" gorm:"type:Varchar(255)" validate:"required,min=5`
	Password  string    `json:"password"  gorm:"type:Varchar(255)" validate:"required,min=5"`
	CreatedAt time.Time `json:"createdat gorm:autoCreateTime" db:"create_at"`
	UpdatedAt time.Time `json:"createdat gorm:autoCreateTime" db:"update_At"`
}
