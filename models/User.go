package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User : Mongodb schema for users
type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username"`
	Email    string             `json:"email" bson:"email"`
	Avatar   string             `json:"avatar" bson:"avatar"`
	// Avatar   primitive.Binary   `json:"avatar" bson:"avatar"`
	Password string               `json:"password" bson:"password"`
	Posts    []primitive.ObjectID `json:"posts" bson:"posts"`
}
