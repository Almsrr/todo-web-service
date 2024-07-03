package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"almsrr/todo-web-service/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var todos = []models.Todo{
	{Id: "1", Title: "Blue Train", Description: "John Coltrane", Completed: false},
	{Id: "2", Title: "Jeru", Description: "Gerry Mulligan", Completed: false},
	{Id: "3", Title: "Sarah Vaughan and Clifford Brown", Description: "Sarah Vaughan", Completed: false},
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/api/todos", getTodos)
	router.GET("/api/todos/:id", getTodoById)
	router.POST("/api/todos", postTodo)
	router.PUT("/api/todos/:id", updateTodoById)
	router.DELETE("/api/todos/:id", deleteTodoById)

	router.Run("localhost:8080")
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func postTodo(c *gin.Context) {
	var newTodo models.Todo
	if err := c.BindJSON((&newTodo)); err != nil {
		return
	}

	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoById(c *gin.Context) {
	id := c.Param("id")

	for _, todo := range todos {
		if todo.Id == id {
			c.IndentedJSON(http.StatusOK, todo)
			return
		}
	}
	c.IndentedJSON((http.StatusNotFound), gin.H{"message": "Todo not found"})
}

func updateTodoById(c *gin.Context) {
	id := c.Param("id")
	var savedTodo models.Todo
	var itWasFound bool

	for i, todo := range todos {
		if todo.Id == id {
			savedTodo = todos[i]
			itWasFound = true
			break
		}
	}

	fmt.Println(savedTodo)
	fmt.Println(itWasFound)
	rawData, _ := c.GetRawData()
	fmt.Println(string(rawData))

	var data map[string]interface{}
	json.Unmarshal(rawData, &data)
	fmt.Println(data)

}

func deleteTodoById(c *gin.Context) {
	id := c.Param("id")

	for i, todo := range todos {
		if todo.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"id": id})
			return
		}
	}

	c.IndentedJSON((http.StatusNotFound), gin.H{"message": "Todo not found"})
}
