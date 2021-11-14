package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100)" json:"name" binding:"required"`
	Email       string `gorm:"type:varchar(100)" json:"email" binding:"required"`
	Password    string `gorm:"type:varchar(100)" json:"password" `
	PhoneNumber string `gorm:"type:varchar(100)" json:"phone_number`
}
