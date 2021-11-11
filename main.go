package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

var Users map[int]User

func main() {
	Users = make(map[int]User, 0)

	Users[1] = User{
		name:        "Mohamed Ragab",
		email:       "mohamed@gmail.com",
		password:    "123456",
		phoneNumber: "01096023385",
	}

	Users[2] = User{
		name:        "Ahmed Ragab",
		email:       "ahmed@gmail.com",
		password:    "123456",
		phoneNumber: "01101285885",
	}
	data,err :=json.Marshal(Users[1])
	fmt.Println(string(data),err)
	r := gin.Default()
	r.GET("/posts", Posts)
	r.GET("/posts/:id", Show)
	r.POST("/posts", Store)
	r.PUT("/posts/:id", Update)
	r.DELETE("/posts/:id", Delete)
	r.Run(":9090")
}
