package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	Id          primitive.ObjectID
	Title       string
	Description string
	ReleaseDate string
	Author      string
	Value       float32
}
