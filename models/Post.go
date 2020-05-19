package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Post : Mongodb schema for posts
type Post struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserName string             `json:"username" bson:"username"`
	Title    string             `json:"title" bson:"title"`
	Content  string             `json:"content" bson:"content"`
}
