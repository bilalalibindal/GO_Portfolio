package main

import (
	"crud_api_example_1/controllers/user"
	"crud_api_example_1/database"

	"github.com/gin-gonic/gin"
)


func main() {
	database.Connect()
	defer database.Disconnect()


	router := gin.Default()
	api := router.Group("/api")
	users := api.Group("/users")

	users.GET("/:size", user.GetUsers)
	users.POST("/add", user.InsertUser)
	users.DELETE("/:id/delete", user.DeleteUser)
	users.PUT("/:id/update", user.UpdateUser)

	router.Run()
}