package main

import "gorm.io/gorm"

type UserCourse struct {
	gorm.Model
	UserID   string `json:"userID" binding:"required"`
	User     User   `json:"user"`
	CourseID string `json:"courseID"`
	Course   Course `json:"course"`
	Score    int    `json:"score"`
}
