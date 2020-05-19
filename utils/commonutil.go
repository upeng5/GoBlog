package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

// RemoveIndex : Removes an element from a slice of mongodb objectids
func RemoveIndex(s []primitive.ObjectID, index int) []primitive.ObjectID {
	return append(s[:index], s[index+1:]...)
}
