package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCourses(g *gin.Context) {
	var courses []Course
	db.Find(&courses)
	if len(courses) == 0 {
		g.JSON(http.StatusOK, gin.H{
			"data":    "{}",
			"message": "No courses found",
			"error":   nil,
		})
		return
	}
		g.JSON(http.StatusOK,gin.H{
			"data":courses,
			"message":"all Courses found",
			"error":nil,
		})
}

func StoreCourse(g *gin.Context) {
	var course Course
	if err:=g.ShouldBindJSON(&course);err != nil{
		g.JSON(http.StatusBadRequest,gin.H{
			"error":err,
		})
		return
	}
	db.Create(&course)
	g.JSON(http.StatusCreated,gin.H{
		"data":course,
		"message":"Course added successfully",
		"error":nil,
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
				"error":err,
			})
			return
	    }
		oldCourse.Name=updateCourse.Name
		db.Save(&oldCourse)

		g.JSON(http.StatusOK,gin.H{
			"data":oldCourse,
			"message":"Course updated successfully",
			"error":nil,
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
		"error":nil,
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
		"error":nil,
	})
}