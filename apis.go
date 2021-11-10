package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)



func Posts(g *gin.Context){
	if len(Users) == 0{
		g.JSON(404,"Not found users"+string(len(Users)))
	}else {
		UsersJson,_ := json.Marshal(Users)
		g.JSON(http.StatusOK,UsersJson)
	}

}

func Store(g *gin.Context){

}

func Update(g *gin.Context){

}

func Delete(g *gin.Context){

}

func Show(g *gin.Context){

}