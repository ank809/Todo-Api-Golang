package main

import (
	"github.com/ank809/Todo-Api-Golang/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/todos", controllers.GetAllTodos)
	router.GET("/create", controllers.CreateTodo)
	router.GET("/getbyid/:id", controllers.GetTodoById)
	router.GET("/delete/:id", controllers.DeleteTodo)
	router.GET("/update/:id", controllers.UpdateTodo)
	router.Run(":8081")

}
