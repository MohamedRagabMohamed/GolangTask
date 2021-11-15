package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type Course struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100)" json:"name" binding:"required"`
	Code        string `json:"code"`
	TeacherName string `json:"teacher_name"`
}

type UserCourse struct {
	gorm.Model
	UserID   string `json:"userID" binding:"required"`
	User     User   `json:"user"`
	CourseID string `json:"courseID"`
	Course   Course `json:"course"`
	Score    int    `json:"score"`
}
