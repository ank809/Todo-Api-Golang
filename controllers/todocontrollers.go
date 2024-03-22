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
	"gopkg.in/mgo.v2/bson"
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
	// var todo models.Todo = models.Todo{
	// 	ID:          id,
	// 	Title:       "Aws",
	// 	Description: "Revise Aws services as CloudFront, EC2",
	// 	CreatedAt:   time.Now(),
	// 	UpdatedAt:   time.Now(),
	// }
	var todo models.Todo
	if err := c.BindJSON(&todo); err != nil {
		fmt.Println("Error in binding json")
		return
	}
	todo.ID = id
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()
	collname := "todos"
	coll := database.OpenCollection(database.Client, collname)
	_, err := coll.InsertOne(context.Background(), todo)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, "Todo successfully added")

}

func GetTodoById(c *gin.Context) {
	id := c.Param("id")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("Invalid Id")
		return
	}
	collection_name := "todos"
	coll := database.OpenCollection(database.Client, collection_name)
	filter := bson.M{"_id": objectId}
	var todo models.Todo

	if err := coll.FindOne(context.Background(), filter).Decode(&todo); err != nil {
		fmt.Println("Todo not found")
		return
	}
	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("Invalid Id")
		return
	}
	collectionName := "todos"
	coll := database.OpenCollection(database.Client, collectionName)
	filter := bson.M{"_id": objectID}

	count, err := coll.DeleteOne(context.Background(), filter)
	if err != nil {
		fmt.Println("Error in deleting todos")
		return
	}
	if count.DeletedCount == 0 {
		fmt.Println("Todo not found")
		return
	}
	c.JSON(http.StatusOK, "Todo successfully deleted")

}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("Invalid id")
		return
	}
	var updatedtodo models.Todo
	if err := c.BindJSON(&updatedtodo); err != nil {
		fmt.Println("Error in binding json")
		return
	}
	collectionName := "todos"
	coll := database.OpenCollection(database.Client, collectionName)
	filter := bson.M{"_id": objectID}
	updateTodo := bson.M{
		"$set": bson.M{
			"title":       updatedtodo.Title,
			"description": updatedtodo.Description,
			"iscompleted": updatedtodo.IsCompleted,
			"updatedat":   updatedtodo.UpdatedAt,
		},
	}
	res, err := coll.UpdateOne(context.Background(), filter, updateTodo)
	if err != nil {
		fmt.Println(err)
		return
	}
	if res.ModifiedCount == 0 {
		fmt.Println("Todo not found")
		return
	}
	c.JSON(http.StatusOK, "TODO UPDATED SUCCESSFULLY")
}
