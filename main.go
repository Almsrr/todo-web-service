package main

import (
	"net/http"
	"reflect"
	"test/my-todo-api/models"

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

	for i, a := range todos {
		if a.Id != id {
			c.IndentedJSON((http.StatusNotFound), gin.H{"message": "Todo not found"})
		} else {
			savedTodo = todos[i]
			itWasFound = true
			break
		}
	}

	if itWasFound {
		if body := c.Request.Body; body != nil {

			v := reflect.ValueOf(body)
			for i := 0; i < v.NumField(); i++ {
				propName := v.Type().Field(i).Name
				propValue := v.Type().Field(i)

				_, exists := savedTodo.propName

				if exists {
					savedTodo[propName] = propValue
				}
				// c.IndentedJSON(http.StatusOK, a)
			}

		} else {
			c.IndentedJSON((http.StatusOK), gin.H{"message": "Nothing to update"})
		}
	}

	// if body := c.Request.Body; body != nil {

	// 	v := reflect.ValueOf(body)
	// 	for i := 0; i < v.NumField(); i++ {
	// 		propName := v.Type().Field(i).Name
	// 		propValue := v.Type().Field(i)

	// 		_, exists := savedTodo.propName

	// 		if exists {
	// 			savedTodo[propName] = propValue
	// 			c.IndentedJSON(http.StatusOK, a)
	// 		}
	// 	}

	// }

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
