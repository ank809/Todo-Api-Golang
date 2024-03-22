package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ank809/Todo-Api-Golang/database"
	"github.com/ank809/Todo-Api-Golang/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllTodos(c *gin.Context) {
	collection_name := "todos"
	coll := database.OpenCollection(database.Client, collection_name)
	cursor, err := coll.Find(context.Background(), coll)
	if err != nil {
		fmt.Println(err)
		return
	}
	var todos []models.Todo
	if err = cursor.All(context.Background(), &todos); err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, todos)

}

func CreateTodo(c *gin.Context) {
	id := primitive.NewObjectID()
	var todo models.Todo = models.Todo{
		ID:          id,
		Title:       "Aws",
		Description: "Revise Aws services as CloudFront, EC2",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	collname := "todos"
	coll := database.OpenCollection(database.Client, collname)
	_, err := coll.InsertOne(context.Background(), todo)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, "Document successfully added")

}
