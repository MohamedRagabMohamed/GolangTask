package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB =nil
var err error

func main() {
	
	dsn := "root:@tcp(127.0.0.1:3306)/crud?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting database")
	}
	db.AutoMigrate(&User{})
	r := gin.Default()
	r.GET("/users", Users)
	r.GET("/users/:id", Show)
	r.POST("/users", Store)
	r.PUT("/users/:id", Update)
	r.DELETE("/users/:id", Delete)
	r.Run(":9090")
}
