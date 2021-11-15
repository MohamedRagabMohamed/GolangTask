package main

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100)" json:"name" binding:"required"`
	Code        string `json:"code"`
	TeacherName string `json:"teacher_name"`
}
