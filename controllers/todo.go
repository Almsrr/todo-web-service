package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

	c.IndentedJSON(http.StatusCreated, gin.H{"id": id})
}

func GetTodoById(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	row := handlers.DB.QueryRow("SELECT * FROM Todo WHERE id = ?", id)
	if err := row.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Completed); err != nil {
		if err == sql.ErrNoRows {

			c.IndentedJSON((http.StatusNotFound), gin.H{"message": "Todo not found"})
			return
		}
		panic(err.Error())
	}

	c.IndentedJSON(http.StatusOK, todo)
}

func UpdateTodoById(c *gin.Context) {
	id := c.Param("id")
	body, _ := c.GetRawData()
	rawData := make(map[string]string)

	for i, todo := range data.Todos {
		if todo.Id == 1 {
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

	result, err := handlers.DB.Exec("DELETE FROM Todo WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	rows, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	if rows == 0 {
		c.IndentedJSON((http.StatusNotFound), gin.H{"message": "Todo not found"})
		return
	}

	i, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("Cannot parse id")
	}

	c.IndentedJSON(http.StatusOK, gin.H{"id": i})

}
