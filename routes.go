package main

import "github.com/gin-gonic/gin"

func UserRoute(r *gin.Engine) {
	r.GET("/users", GetALlUser)
	r.GET("/users/:id", Show)
	r.POST("/users", Store)
	r.PUT("/users/:id", Update)
	r.DELETE("/users/:id", Delete)
}
