package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllUsers(g *gin.Context) {
	var users []User
	db.Find(&users)
	if len(users) == 0{
		g.JSON(http.StatusOK,gin.H{
		"data":"",
		"message":"No users found",
		"error":nil,
		})
	}else{
		g.JSON(http.StatusOK,gin.H{
			"data":users,
			"message":"all users found",
			"error":nil,
			})
	}
}

func StoreUser(g *gin.Context) {
	var user User
	if err:=g.ShouldBindJSON(&user);err != nil{
		g.JSON(http.StatusBadRequest,gin.H{
			"error":err,
		})
		return
	}
	password, _ := HashPassword(user.Password)
	user.Password=string(password)
	db.Create(&user)
	g.JSON(http.StatusCreated,gin.H{
		"data":user,
		"message":"User added successfully",
		"error":nil,
	})
}

func UpdateUser(g *gin.Context) {
		oldUser:=getUserByID(g)
		if oldUser.ID == 0{
			return
		}
		var updateUser User
		var course Course
		courseId:=g.Query("courseId")
		if err:=g.ShouldBindJSON(&updateUser);err != nil{
			g.JSON(http.StatusBadRequest,gin.H{
				"data":"",
				"message":"insert valid user data",
				"error":err,
			})
			return
	    }
		fmt.Println("course id :",oldUser)
		if courseId != "" {
			db.First(&course,courseId)
			usercourse:=UserCourse{
				CourseID: courseId,
				UserID: string(oldUser.ID),
				User: oldUser,
				Course: course,
				Score: 0,
			}
			db.Create(&usercourse)
		}
		oldUser.Name=updateUser.Name
		oldUser.Email=updateUser.Email
		oldUser.Password,_= HashPassword(updateUser.Password)
		oldUser.PhoneNumber=updateUser.PhoneNumber

		db.Save(&oldUser)

		g.JSON(http.StatusOK,gin.H{
			"data":oldUser,
			"message":"User updated successfully",
			"error":nil,
		})
}

func DeleteUser(g *gin.Context) {
	user:=getUserByID(g)
	if user.ID == 0{
		return
	} 
	db.Unscoped().Delete(&user)
	g.JSON(http.StatusOK,gin.H{
		"data":user,
		"message":"User deleted seccessfully",
		"error":nil,
	})
}

func ShowUser(g *gin.Context) {
	user:=getUserByID(g)
	if user.ID == 0{
		return
	}
	var usercours []UserCourse
	var courseIDs []string
	var courses []Course
	db.Where("user_id =?",user.ID).Find(&usercours)
	for _,data :=range  usercours{
		courseIDs = append(courseIDs, data.CourseID)
	}
	fmt.Println(courseIDs)
	db.Where(courseIDs).Find(&courses)
	g.JSON(http.StatusOK,gin.H{
		"data":user,
		"courses":courses,
		"message":"User found seccessfully",
		"error":nil,
	})
}



