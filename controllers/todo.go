package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"almsrr/todo-web-service/data"
	"almsrr/todo-web-service/handlers"
	"almsrr/todo-web-service/models"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todos []models.Todo

	rows, err := handlers.DB.Query("SELECT * FROM Todo")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Completed); err != nil {
			panic(err.Error())
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		panic(err.Error())
	}

	c.IndentedJSON(http.StatusOK, todos)
}

func PostTodo(c *gin.Context) {
	var body models.Todo
	if err := c.BindJSON((&body)); err != nil {
		return
	}

	result, err := handlers.DB.Exec("INSERT INTO Todo (title, description, completed) VALUES (?, ?, ?)", body.Title, body.Description, body.Completed)
	if err != nil {
		panic(err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	c.IndentedJSON(http.StatusCreated, id)
}

func GetTodoById(c *gin.Context) {
	id := c.Param("id")

	for _, todo := range data.Todos {
		if todo.Id == id {
			c.IndentedJSON(http.StatusOK, todo)
			return
		}
	}
	c.IndentedJSON((http.StatusNotFound), gin.H{"message": "Todo not found"})
}

func UpdateTodoById(c *gin.Context) {
	id := c.Param("id")
	body, _ := c.GetRawData()
	rawData := make(map[string]string)

	for i, todo := range data.Todos {
		if todo.Id == id {
			/* Converting request body type []byte to map
			to delete the key id before updating */
			json.Unmarshal(body, &rawData)
			delete(rawData, "id")

			/* Turning rawData back to []bytes so
			Marshal() can match keys and update them */
			jsonBytes, _ := json.Marshal(rawData)
			json.Unmarshal(jsonBytes, &data.Todos[i])

			c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Todo %v updated", id)})
			return

		} else {
			if i == len(data.Todos)-1 {
				c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
				return
			}
		}
	}
}

func DeleteTodoById(c *gin.Context) {
	id := c.Param("id")

	for i, todo := range data.Todos {
		if todo.Id == id {
			data.Todos = append(data.Todos[:i], data.Todos[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"id": id})
			return
		}
	}

	c.IndentedJSON((http.StatusNotFound), gin.H{"message": "Todo not found"})
}
