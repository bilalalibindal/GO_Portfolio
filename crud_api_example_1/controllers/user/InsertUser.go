package user

import (
	"context"
	"crud_api_example_1/database"
	"crud_api_example_1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsertUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ERROR": err.Error()})
		return
	}
	collection := database.DB.Collection("user")
	ctx := context.Background()

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}