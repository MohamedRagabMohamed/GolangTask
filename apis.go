package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Users(g *gin.Context) {
	var users []User
	db.Find(&users)
	if len(users) == 0{
		g.JSON(http.StatusOK,gin.H{
		"data":"{}}",
		"message":"No users found",
		"error":"",
		})
	}else{
		g.JSON(http.StatusOK,gin.H{
			"data":users,
			"message":"all found users",
			"error":"",
			})
	}
}

func Store(g *gin.Context) {
	var user User
	if err:=g.ShouldBindJSON(&user);err != nil{
		g.JSON(http.StatusBadRequest,gin.H{
			"error":"insert valid user data",
		})
		return
	}
	password, _ := HashPassword(user.Password)
	user.Password=string(password)
	db.Create(&user)
	g.JSON(http.StatusCreated,gin.H{
		"data":user,
		"message":"User added successfully",
		"error":"",
	})
}

func Update(g *gin.Context) {
		oldUser:=getUserByID(g)
		if oldUser.ID == 0{
			return
		}
		var updateUser User
		if err:=g.ShouldBindJSON(&updateUser);err != nil{
			g.JSON(http.StatusBadRequest,gin.H{
				"data":"{}",
				"message":"insert valid user data",
				"error":"",
			})
			return
	    }

		oldUser.Name=updateUser.Name
		oldUser.Email=updateUser.Email
		oldUser.Password,_= HashPassword(updateUser.Password)
		oldUser.PhoneNumber=updateUser.PhoneNumber

		db.Save(&oldUser)

		g.JSON(http.StatusOK,gin.H{
			"data":oldUser,
			"message":"User updated successfully",
			"error":"",
		})
}

func Delete(g *gin.Context) {
	user:=getUserByID(g)
	if user.ID == 0{
		return
	} 
	db.Unscoped().Delete(&user)
	g.JSON(http.StatusOK,gin.H{
		"data":user,
		"message":"User deleted seccessfully",
		"error":"",
	})
}

func Show(g *gin.Context) {
	user:=getUserByID(g)
	if user.ID == 0{
		return
	} 
	g.JSON(http.StatusOK,gin.H{
		"data":user,
		"message":"User found seccessfully",
		"error":"",
	})
}


func getUserByID(g *gin.Context)User {
	var user User
	id:=g.Param("id")
	db.Find(&user,id)
	if user.Name==""{
		g.JSON(http.StatusNotFound,gin.H{
			"data":"",
			"message":"No user found",
			"error":"",
		})
	}
	return user
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}