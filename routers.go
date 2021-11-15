package main

import "github.com/gin-gonic/gin"

func UserRouter(r *gin.Engine){
	r.GET("/users", GetAllUsers)
	r.GET("/users/:id", ShowUser)
	r.POST("/users", StoreUser)
	r.PUT("/users/:id", UpdateUser)
	r.DELETE("/users/:id", DeleteUser)
}


func CourseRouter(r *gin.Engine){
	r.GET("/courses",GetAllCourses)
	r.GET("/courses/:id",ShowCourse)
	r.POST("/courses",StoreCourse)
	r.PUT("/courses/:id",UpdateCourse)
	r.DELETE("/courses/:id",DeleteCourse)
}