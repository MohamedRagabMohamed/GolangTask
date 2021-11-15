package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
