package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetALlUser(g *gin.Context) {
	users := UserTransformer(Users)
	g.JSON(http.StatusOK, gin.H{
		"data":    users,
		"message": "Done Get All users",
		"errors":  nil,
	})
}

func Store(g *gin.Context) {
	var user User
	err := g.BindJSON(&user)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"data":    "",
			"message": "User not created",
			"errors":  err,
		})
		return
	}
	id := generateID()
	Users[id] = user
	g.JSON(http.StatusOK, gin.H{
		"data":    user,
		"message": "user added successfully",
		"errors":  nil,
	})
}

func Update(g *gin.Context) {
	//TODO refactor this api
	id, _ := strconv.Atoi(g.Param("id"))
	var user User
	err := g.BindJSON(&user)
	if _, found := Users[id]; found && err == nil {
		Users[id] = User{
			Name:        user.Name,
			Email:       user.Email,
			Password:    user.Password,
			PhoneNumber: user.PhoneNumber,
		}
		g.JSON(http.StatusOK, gin.H{
			"data":    Users[id],
			"message": "user updated successfully",
			"errors":  "No found error",
		})
	} else {
		g.JSON(http.StatusNotFound, gin.H{
			"data":    "{}",
			"message": "No found users",
			"errors":  "No found users",
		})
	}
}

func Delete(g *gin.Context) {
	id, _ := strconv.Atoi(g.Param("id"))
	if value, found := Users[id]; found {
		delete(Users, id)
		g.JSON(http.StatusOK, gin.H{
			"data":    value,
			"message": "User deleted successfully",
			"errors":  nil,
		})
	} else {
		g.JSON(http.StatusNotFound, gin.H{
			"data":    "{}",
			"message": "No found users",
			"errors":  nil,
		})
	}
}

func Show(g *gin.Context) {
	id, _ := strconv.Atoi(g.Param("id"))
	if value, found := Users[id]; found {
		g.JSON(http.StatusOK, gin.H{
			"data":    value,
			"message": "User is found",
			"errors":  nil,
		})
	} else {
		g.JSON(http.StatusNotFound, gin.H{
			"data":    "",
			"message": "No found users",
			"errors":  nil,
		})
	}
}
