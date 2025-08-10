package user

import (
	"context"
	"crud_api_example_1/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteUser(c *gin.Context) {
	userIdFromParam := c.Param("id")
	userId, err := strconv.ParseInt(userIdFromParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ERROR": "Invalid user ID"})
		return
	}

	collection := database.DB.Collection("user")
	ctx := context.Background()
	filter := bson.M{"id": userId}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"ERROR": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	
}