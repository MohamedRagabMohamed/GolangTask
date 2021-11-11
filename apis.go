package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

//TODO change all name to Users
func Posts(g *gin.Context) {
	//TODO if len is retrun 200
	if len(Users) == 0 {
		g.JSON(404, "Not found users")
	} else {
		/*var outUsers []interface{}
		for _, user := range Users {
			outUsers = append(outUsers, user)
		}*/
		g.JSON(http.StatusOK, gin.H{
			"users": Users[1],
		})
	}
}

func Store(g *gin.Context) {
	//TODO Search in gin validation and body request not query param
	enteredName := g.Query("name")
	enteredEmail := g.Query("email")
	enteredPassword := g.Query("password")
	enteredPhoneNumber := g.Query("phoneNumber")

	user := User{
		Name:        enteredName,
		Email:       enteredEmail,
		Password:    enteredPassword,
		PhoneNumber: enteredPhoneNumber,
	}

	validate := validator.New()
	err := validate.Struct(user)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		len := len(Users) + 1
		Users[len] = user
		g.JSON(http.StatusOK, gin.H{
			"New User": user,
		})
	}

}

func Update(g *gin.Context) {

}

func Delete(g *gin.Context) {

}

func Show(g *gin.Context) {

}
