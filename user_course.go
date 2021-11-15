package main

import "gorm.io/gorm"

type UserCourse struct {
	gorm.Model
	UserID        string `gorm:"type:int" json:"userID" binding:"required"`
	CourseID       string `gorm:"type:int" json:"courseID"`
}
