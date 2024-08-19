package dto

type RequestCustomerRegister struct {
	Name        string `json:"name" validate:"required, min:5"`
	Username    string `json:"username" validate:"required, min:5"`
	Password    string `json:"password" validate:"required,alphanum,min=5"`
	Email       string `json:"email" validate:"required,email"`
	NumberPhone string `json:"numberphone" validate:"required,min=10"`
	Gender      string `json:"gender" validate:"required"`
	DateOfBirth string `json:"date_of_birth" validate:"required"`
}

type RequestCustomerLogin struct {
	Username string `json:"username" validate:"required, min:5"`
	Password string `json:"password" validate:"required,alphanum,min=5"`
}


