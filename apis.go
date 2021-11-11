package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//TODO make api response scheme same in all apis
/*
data:
message:
status: true or false
errors:
*/

//TODO change all name to users
func Posts(g *gin.Context) {
	//TODO if len is retrun 200
	if len(Users) == 0 {
		//TODO return empty array
		g.JSON(http.StatusOK, "Not found users")
	} else {
		//TODO make transformer to return array of users not map
		g.JSON(http.StatusOK, gin.H{
			"users": Users,
		})
	}
}

func Store(g *gin.Context) {
	//TODO Search in gin validation and body request not query param
	var user User
	err :=g.BindJSON(&user)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		//TODO rename this variabel and make function to get new id
		len := len(Users) + 1
		Users[len] = user
		g.JSON(http.StatusOK, gin.H{
			"New User": user,
		})
	}

}

func Update(g *gin.Context) {
	id,_:= strconv.Atoi(g.Param("id"))
	//TODO check errors form bind
	var user User
	g.BindJSON(&user)
	if _,found :=Users[id];found {
		Users[id]=User{
			Name: user.Name,
			Email: user.Email,
			Password: user.Password,
			PhoneNumber: user.PhoneNumber,
		}
		g.JSON(http.StatusOK,gin.H{
			"id":Users[id],
		})
	}else{
		g.JSON(http.StatusNotFound,gin.H{
			"message":"No found user",
		})
	}	
}

func Delete(g *gin.Context) {
	id,_:= strconv.Atoi(g.Param("id"))
	if value,found :=Users[id];found {
		delete(Users, id)
		g.JSON(http.StatusOK,gin.H{
			"Deleted user":value,
		})
	}else{
		g.JSON(http.StatusNotFound,gin.H{
			"message":"No found user",
		})
	}
}

func Show(g *gin.Context) {
	id,_:= strconv.Atoi(g.Param("id"))
	if value,found :=Users[id];found {
		g.JSON(http.StatusOK,gin.H{
			"id":value,
		})
	}else{
		g.JSON(http.StatusNotFound,gin.H{
			"message":"No found user",
		})
	}
}
