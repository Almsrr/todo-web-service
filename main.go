package main

import (
	"almsrr/todo-web-service/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)



func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/api/todos", controllers.GetTodos)
	router.GET("/api/todos/:id", controllers.GetTodoById)
	router.POST("/api/todos", controllers.PostTodo)
	router.PUT("/api/todos/:id", controllers.UpdateTodoById)
	router.DELETE("/api/todos/:id", controllers.DeleteTodoById)

	router.Run("localhost:8080")
}



