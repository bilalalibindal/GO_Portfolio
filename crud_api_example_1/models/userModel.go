package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type User struct {
	ID            int                `bson:"id" json:"id"`
	IsActive      bool               `bson:"isActive" json:"isActive"`
	Balance       string             `bson:"balance" json:"balance"`
	Picture       string             `bson:"picture" json:"picture"`
	Age           int                `bson:"age" json:"age"`
	Name          string             `bson:"name" json:"name"`
	Gender        string             `bson:"gender" json:"gender"`
	Company       string             `bson:"company" json:"company"`
	Email         string             `bson:"email" json:"email"`
	Phone         string             `bson:"phone" json:"phone"`
	Address       string             `bson:"address" json:"address"`
	About         string             `bson:"about" json:"about"`
	Registered    primitive.DateTime `bson:"registered" json:"registered"`
	Latitude      float64            `bson:"latitude" json:"latitude"`
	Longitude     float64            `bson:"longitude" json:"longitude"`
	FavoriteFruit string             `bson:"favoriteFruit" json:"favoriteFruit"`
}