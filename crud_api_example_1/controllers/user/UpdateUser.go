package user

import (
	"context"
	"crud_api_example_1/database"
	"crud_api_example_1/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func UpdateUser(c *gin.Context) {
	// We are using map[string]interface{} to store the partial user data
	var partialUser map[string]interface{}

	// We are getting the user id from the url
	userIdFromParam := c.Param("id")
	// turn userId to INT
	userId, err := strconv.ParseInt(userIdFromParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ERROR": "Invalid user ID"})
		return
	}
	// We are binding the partial user data to the partialUser map
	if err := c.ShouldBindJSON(&partialUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ERROR": err.Error()})
		return
	}

	// As default we are not allowed to update id and _id fields
	delete(partialUser, "id")
	delete(partialUser, "_id")

	// We are creating a context
	ctx := context.Background()

	// Creating a updated user variable from the models package
	var updatedUser models.User
	// Define collection
	collection := database.DB.Collection("user")
	// Create a $set operation to update user only the fields that are in the partialUser map
	update := bson.M{"$set": partialUser}
    
	// Create a options to return the updated user
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	filter := bson.M{"id": userId}

	err = collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// If user not found, return 404 Not Found
			c.JSON(http.StatusNotFound, gin.H{"ERROR": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, updatedUser)
}
