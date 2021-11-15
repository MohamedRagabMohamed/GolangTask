package main

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100)" json:"name" binding:"required"`
	Score       string `gorm:"type:int" json:"score"`
	Users       []User `gorm:"many2many:user_courses;"`
}
