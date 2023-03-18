package dto

type LoginDTO struct {
	Account  string `json:"account" form:"account"`
	Password string `json:"password" form:"password"`
}

type Register struct {
	Name        string `json:"name" form:"name"`
	Account     string `json:"account" form:"account"`
	Password    string `json:"password" form:"password"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
}
