package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Courses(g *gin.Context) {
	var courses []Course
	db.Find(&courses)
	if len(courses) == 0{
		g.JSON(http.StatusOK,gin.H{
		"data":"{}",
		"message":"No courses found",
		"error":"",
		})
	}else{
		g.JSON(http.StatusOK,gin.H{
			"data":courses,
			"message":"all Courses found",
			"error":"",
			})
	}
}

func StoreCourse(g *gin.Context) {
	var course Course
	if err:=g.ShouldBindJSON(&course);err != nil{
		g.JSON(http.StatusBadRequest,gin.H{
			"error":"insert valid course data",
		})
		return
	}
	db.Create(&course)
	g.JSON(http.StatusCreated,gin.H{
		"data":course,
		"message":"Course added successfully",
		"error":"",
	})
}

func UpdateCourse(g *gin.Context) {
		oldCourse:=getCourseByID(g)
		if oldCourse.ID == 0{
			return
		}
		var updateCourse Course
		if err:=g.ShouldBindJSON(&updateCourse);err != nil{
			g.JSON(http.StatusBadRequest,gin.H{
				"data":"{}",
				"message":"insert valid course data",
				"error":"",
			})
			return
	    }

		oldCourse.Name=updateCourse.Name
		oldCourse.Score=updateCourse.Score

		db.Save(&oldCourse)

		g.JSON(http.StatusOK,gin.H{
			"data":oldCourse,
			"message":"Course updated successfully",
			"error":"",
		})
}

func DeleteCourse(g *gin.Context) {
	course:=getCourseByID(g)
	if course.ID == 0{
		return
	} 
	db.Unscoped().Delete(&course)
	g.JSON(http.StatusOK,gin.H{
		"data":course,
		"message":"Course deleted seccessfully",
		"error":"",
	})
}

func ShowCourse(g *gin.Context) {
	course:=getCourseByID(g)
	if course.ID == 0{
		return
	} 
	g.JSON(http.StatusOK,gin.H{
		"data":course,
		"message":"Course found seccessfully",
		"error":"",
	})
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