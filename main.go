package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/subosito/gotenv"
)

var Users map[int]User

func main() {
	Users = make(map[int]User, 0)
	Seeder()
	r := gin.Default()
	UserRoute(r)
	err := gotenv.Load()
	if err != nil {
		return
	}
	err = r.Run(os.Getenv("PORT"))
	if err != nil {
		return
	}
}
