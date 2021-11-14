package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/icrowley/fake"
)

//TODO make api response scheme same in all apis
/*
data:
message:
status: true or false
errors:
*/

//TODO change all name to users
func UserS(g *gin.Context) {
	//TODO if len is retrun 200
	if len(Users) == 0 {
		//TODO return empty array
		g.JSON(http.StatusOK, gin.H{
			"data":Users,
			"message":"No found users",
			"errors":"No found error",
		})
	} else {
		//TODO make transformer to return array of users not map
		/*var outputUsers []User
		for _, value := range Users {
			outputUsers = append(outputUsers, value)
		}*/
		g.JSON(http.StatusOK, gin.H{
			"data":Users,
			"message":"All users",
			"errors":"No found error",
		})
	}
}

func Store(g *gin.Context) {
	//TODO Search in gin validation and body request not query param
	var user User
	err :=g.BindJSON(&user)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"data":"",
			"message":"User not created",
			"errors":"please enter valid user data",
		})
	} else {
		//TODO rename this variabel and make function to get new id
		id := generateID()
		Users[id] = user
		g.JSON(http.StatusOK, gin.H{
			"data":user,
			"message":"user added successfully",
			"errors":"No found error",
		})
	}

}

func Update(g *gin.Context) {
	id,_:= strconv.Atoi(g.Param("id"))
	//TODO check errors form bind
	var user User
	err := g.BindJSON(&user)
	if _,found :=Users[id];found && err == nil{
		Users[id]=User{
			Name: user.Name,
			Email: user.Email,
			Password: user.Password,
			PhoneNumber: user.PhoneNumber,
		}
		g.JSON(http.StatusOK,gin.H{
			"data":Users[id],
			"message":"user updated successfully",
			"errors":"No found error",
		})
	}else{
		g.JSON(http.StatusNotFound,gin.H{
			"data":"{}",
			"message":"No found users",
			"errors":"No found users",
		})
	}	
}

func Delete(g *gin.Context) {
	id,_:= strconv.Atoi(g.Param("id"))
	if value,found :=Users[id];found {
		delete(Users, id)
		g.JSON(http.StatusOK,gin.H{
			"data":value,
			"message":"User deleted successfully",
			"errors":"No found error",
		})
	}else{
		g.JSON(http.StatusNotFound,gin.H{
			"data":"{}",
			"message":"No found users",
			"errors":"No found users",
		})
	}
}

func Show(g *gin.Context) {
	id,_:= strconv.Atoi(g.Param("id"))
	if value,found :=Users[id];found {
		g.JSON(http.StatusOK,gin.H{
			"data":value,
			"message":"User is found",
			"errors":"No found error",
		})
	}else{
		g.JSON(http.StatusNotFound,gin.H{
			"data":"{}",
			"message":"No found users",
			"errors":"No found users",
		})
	}
}

// additional function

func Seeder(){
	for i:= 1; i <= 10 ; i++ {
		Users[i] = User{
			Name:        fake.FullName(),	
			Email:       fake.EmailAddress(),
			Password:    fake.SimplePassword(),
			PhoneNumber: fake.Phone(),
		}
	}
}

func generateID()int{
	return len(Users)+1
}
