package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var Users map[int]User

func main() {
	Users = make(map[int]User, 0)

	Users[1] = User{
		Name:        "Mohamed Ragab",
		Email:       "mohamed@gmail.com",
		Password:    "123456",
		PhoneNumber: "01096023385",
	}

	Users[2] = User{
		Name:        "Ahmed Ragab",
		Email:       "ahmed@gmail.com",
		Password:    "123456",
		PhoneNumber: "01101285885",
	}
	fmt.Println(Users)
	r := gin.Default()
	r.GET("/posts", Posts)
	r.GET("/posts/:id", Show)
	r.POST("/posts", Store)
	r.PUT("/posts/:id", Update)
	r.DELETE("/posts/:id", Delete)
	r.Run(":9090")
}
