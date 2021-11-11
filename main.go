package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var Users map[int]User

func main() {
	Users = make(map[int]User, 0)
	//TODO make this add users in seeder function and users fake package to add users
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
	//TODO change posts to users
	r := gin.Default()
	r.GET("/posts", Posts)
	r.GET("/posts/:id", Show)
	r.POST("/posts", Store)
	r.PUT("/posts/:id", Update)
	r.DELETE("/posts/:id", Delete)
	//TODO make value in env file for port number
	r.Run(":9090")
}
