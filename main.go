package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/subosito/gotenv"
)

var Users map[int]User

func main() {
	Users = make(map[int]User, 0)
	//TODO make this add users in seeder function and users fake package to add users
	Seeder()
	//TODO change posts to users
	r := gin.Default()
	r.GET("/users", UserS)
	r.GET("/users/:id", Show)
	r.POST("/users", Store)
	r.PUT("/users/:id", Update)
	r.DELETE("/users/:id", Delete)
	//TODO make value in env file for port number
	gotenv.Load()
	r.Run(os.Getenv("PORT"))
}
