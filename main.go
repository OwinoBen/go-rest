package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type todo struct{
	ID 			int 	`json:"id"`
	Item 		string 	`json:"title"`
	Completed 	bool 	`json:"completed"`
}

var todos = []todo{
	{ID: 1, Item: "Clean room", Completed: false},
	{ID: 2, Item: "Read Book", Completed: false},
	{ID: 3, Item: "Record video", Completed: false},
	{ID: 4, Item: "Write codes", Completed: false},
	
}

func getTodos(context *gin.Context){
	context.IndentedJSON(http.StatusOK, todos)
}

func main(){
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.Run("localhost:9090")
}