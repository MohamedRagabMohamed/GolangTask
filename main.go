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

	db.AutoMigrate(&User{},&Course{},&UserCourse{})

	// user APIs
	r := gin.Default()
	r.GET("/users", Users)
	r.GET("/users/:id", ShowUser)
	r.POST("/users", StoreUser)
	r.PUT("/users/:id", UpdateUser)
	r.DELETE("/users/:id", DeleteUser)

	// course APIs
	r.GET("/courses",Courses)
	r.GET("/courses/:id",ShowCourse)
	r.POST("/courses",StoreCourse)
	r.PUT("/courses/:id",UpdateCourse)
	r.DELETE("/courses/:id",DeleteCourse)

	
	r.Run(":9090")
}
