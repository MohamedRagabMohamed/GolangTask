package main

type User struct {
	Name        string `json:"name" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Email       string `json:"email" validate:"required&email"`
	PhoneNumber string `json:"phoneNumber" validate:"required,gte=10"`
}
