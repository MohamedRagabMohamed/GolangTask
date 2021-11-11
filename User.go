package main

type User struct {
	name        string `json:"name" validate:"required"`
	password    string `json:"password" validate:"required"`
	email       string `json:"email" validate:"required&email"`
	phoneNumber string `json:"phoneNumber" validate:"required,gte=10"`
}
