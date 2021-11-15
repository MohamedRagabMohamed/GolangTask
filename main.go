package main

import (
	"log"
	"os"
	"github.com/subosito/gotenv"
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

	r := gin.Default()

	UserRouter(r)
	
	CourseRouter(r)

	err := gotenv.Load()
	if err != nil {
		return
	}

	err = r.Run(os.Getenv("PORT"))
	if err != nil {
		return 
	}
}
