package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)


func Posts(g *gin.Context){
	if len(Users) == 0{
		g.JSON(404,"Not found users")
	}else {
		/*var outUsers []interface{}
		for _, user := range Users {
			outUsers = append(outUsers, user)
		}*/
	 	data,err :=json.Marshal(Users)
		fmt.Println(string(data),err)
		g.JSON(http.StatusOK,gin.H{
			"Users":data,
		})
	}

}

func Store(g *gin.Context){
	enteredName:=g.Query("name")
	enteredEmail:=g.Query("email")
	enteredPassword:=g.Query("password")
	enteredPhoneNumber:=g.Query("phoneNumber")

	user:=User{
		name: enteredName,
		email: enteredEmail,
		password: enteredPassword,
		phoneNumber: enteredPhoneNumber,
	}

	validate:= validator.New()
	err:=validate.Struct(user)

	if err!=nil{
		g.JSON(http.StatusBadRequest,gin.H{
			"message":err.Error(),
		})
	}else{
		len := len(Users)+1
		Users[len]=user
		g.JSON(http.StatusOK,gin.H{
			"New User":user,
		})
	}

}

func Update(g *gin.Context){
	
}

func Delete(g *gin.Context){
	
}

func Show(g *gin.Context){
	
}