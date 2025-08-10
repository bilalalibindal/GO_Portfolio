package user

import (
	"context"
	"crud_api_example_1/database"
	"crud_api_example_1/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUsers(c *gin.Context) {
	sizeParam := c.Param("size")
	size, err := strconv.ParseInt(sizeParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	collection := database.DB.Collection("user")
	ctx := context.Background()
	var users []models.User
	
	findOptions := options.Find().SetLimit(size)
	result, err := collection.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
		return
	}

	defer result.Close(ctx)

	if err = result.All(ctx, &users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}