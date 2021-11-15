package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)



func getUserByID(g *gin.Context)User {
	var user User
	id:=g.Param("id")
	db.Find(&user,id)
	if user.Name==""{
		g.JSON(http.StatusNotFound,gin.H{
			"data":"",
			"message":"No user found",
			"error":"",
		})
	}
	return user
}

func getCourseByID(g *gin.Context)Course {
	var course Course
	id:=g.Param("id")
	db.Find(&course,id)
	if course.Name==""{
		g.JSON(http.StatusNotFound,gin.H{
			"data":"",
			"message":"No course found",
			"error":"",
		})
	}
	return course
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
