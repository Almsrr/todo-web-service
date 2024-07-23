package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"almsrr/todo-web-service/handlers"
	"almsrr/todo-web-service/models"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todos []models.Todo

	result := handlers.DB.Order("completed asc").Order("id desc").Find(&todos)

	if result.Error != nil {
		log.Fatal(result.Error.Error())
	}

	c.IndentedJSON(http.StatusOK, todos)
}

func GetTodoById(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	result := handlers.DB.First(&todo, id)

	if result.RowsAffected == 0 {
		c.IndentedJSON((http.StatusNotFound), gin.H{"message": "Todo not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, todo)
}

func PostTodo(c *gin.Context) {
	var newTodo models.Todo

	if err := c.BindJSON((&newTodo)); err != nil {
		return
	}

	handlers.DB.Create(&newTodo)

	c.IndentedJSON(http.StatusCreated, gin.H{"id": newTodo.ID})
}

func UpdateTodoById(c *gin.Context) {
	id := c.Param("id")
	rawData, _ := c.GetRawData()
	var todo models.Todo

	result := handlers.DB.First(&todo, id)

	if result.RowsAffected == 0 {
		c.IndentedJSON((http.StatusNotFound), gin.H{"message": "Todo not found"})
		return
	}

	if err := json.Unmarshal(rawData, &todo); err != nil {
		panic(err.Error())
	}

	handlers.DB.Model(models.Todo{}).Where("id = ?", id).Update("title", todo.Title).
		Update("description", todo.Description).Update("completed", todo.Completed)

	c.IndentedJSON(http.StatusOK, gin.H{"id": todo.ID})
}

func DeleteTodoById(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	result := handlers.DB.First(&todo, id)

	if result.RowsAffected == 0 {
		c.IndentedJSON((http.StatusNotFound), gin.H{"message": "Todo not found"})
		return
	}

	handlers.DB.Delete(&todo)

	c.IndentedJSON(http.StatusOK, gin.H{"id": todo.ID})

}
