package entity

import "time"

type Customer struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name" gorm:"type:varchar(255)" validate:"required, min:5"`
	Username    string    `json:"username" gorm:"type:varchar(255)" validate:"required, min:5"`
	Password    string    `json:"password" gorm:"type:varchar(255)" validate:"required,alphanum,min=5"`
	Email       string    `json:"email" gorm:"type:varchar(255)" validate:"required,email"`
	NumberPhone string    `json:"numberphone" gorm:"type:varchar(255)" validate:"required,min=10"`
	Gender      string    `json:"gender" gorm:"type:varchar(50)" validate:"required"`
	DateOfBirth string    `json:"date_of_birth" gorm:"type:varchar(100)" validate:"required"`
	Image       string    `json:"image" gorm:"type:varchar(100)" validate:"required"`
	Address     string    `json:"address" gorm:"type:varchar(255)" validate:"required"`
	CreatedAt   time.Time `json:"createdat" gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedat" gorm:"autoCreateTime" db:"updated_at"`
}
