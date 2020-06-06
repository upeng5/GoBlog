package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"goblog/database"
	"goblog/models"
	"goblog/utils"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type postRequestBody struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// GetPosts : Gets all the posts in the database
func GetPosts(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if req.Header.Get("user-name") == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Unauthorized",
		})
		return
	}

	client := database.GetDB()
	var posts []models.Post

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	collection := client.Database(dbName).Collection("posts")

	// TODO find by logged in user id
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	for cursor.Next(ctx) {
		var post models.Post
		cursor.Decode(&post)
		posts = append(posts, post)
	}

	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(posts)

}

// GetUserPosts : Gets all the posts by the authenticated user in the database
func GetUserPosts(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client := database.GetDB()
	var posts []models.Post

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	collection := client.Database(dbName).Collection("posts")

	filter := bson.M{"username": req.Header.Get("user-name")}

	// TODO find by logged in user id
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Server Error Occurred",
			"error":   err.Error(),
		})
		return
	}

	for cursor.Next(ctx) {
		var post models.Post
		cursor.Decode(&post)
		posts = append(posts, post)
	}

	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Server Error Occurred",
			"error":   err.Error(),
		})
		return
	}

	if len(posts) == 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "User has not posted anything",
		})
		return
	}

	json.NewEncoder(w).Encode(posts)

}

// CreatePost : Creates a post using a POST request
func CreatePost(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client := database.GetDB()
	var requestBody postRequestBody
	var post models.Post
	var user models.User

	json.NewDecoder(req.Body).Decode(&requestBody)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	userCollection := client.Database(dbName).Collection("users")
	postCollection := client.Database(dbName).Collection("posts")

	filter := bson.M{"username": req.Header.Get("user-name")}

	userRes := userCollection.FindOne(ctx, filter)
	if userRes.Err() != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Server Error Occurred",
			"error":   userRes.Err().Error(),
		})
		return
	}

	userRes.Decode(&user)

	post = models.Post{ID: primitive.NewObjectID(), Title: requestBody.Title, Content: requestBody.Content, UserName: user.Username}
	user.Posts = append(user.Posts, post.ID)

	update := bson.M{
		"$set": user,
	}

	_, err := userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Server Error Occurred",
			"error":   err.Error(),
		})
		return
	}

	res, err := postCollection.InsertOne(ctx, post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Server Error Occurred",
			"error":   err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(res)

}

// UpdatePost : Updates a post with username
func UpdatePost(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client := database.GetDB()
	var requestBody postRequestBody
	var post models.Post
	postID := chi.URLParam(req, "postId")
	fmt.Println("UpdatePost DEBUG:", postID)

	json.NewDecoder(req.Body).Decode(&requestBody)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := client.Database(dbName).Collection("posts")

	postObjID, _ := primitive.ObjectIDFromHex(postID)

	filter := bson.M{"_id": postObjID}

	postRes := collection.FindOne(ctx, filter)
	if postRes.Err() != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Server Error Occurred",
			"error":   postRes.Err().Error(),
		})
		return
	}

	postRes.Decode(&post)

	post.Title = requestBody.Title
	post.Content = requestBody.Content

	update := bson.M{
		"$set": post,
	}

	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Server Error Occurred",
			"error":   err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(res)
}

// DeletePost : Deletes a post with a user id
func DeletePost(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client := database.GetDB()
	var user models.User
	postID := chi.URLParam(req, "postId")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	userCollection := client.Database(dbName).Collection("users")
	postCollection := client.Database(dbName).Collection("posts")

	postObjID, _ := primitive.ObjectIDFromHex(postID)

	userFilter := bson.M{"username": req.Header.Get("user-name")}
	postFilter := bson.M{"_id": postObjID}

	userRes := userCollection.FindOne(ctx, userFilter)
	if userRes.Err() != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Server Error Occurred",
			"error":   userRes.Err().Error(),
		})
		return
	}

	userRes.Decode(&user)

	res, err := postCollection.DeleteOne(ctx, postFilter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Server Error Occurred",
			"error":   err.Error(),
		})
		return
	}

	var i int
	for i = 0; i < len(user.Posts); i++ {
		if user.Posts[i] == postObjID {
			user.Posts = utils.RemoveIndex(user.Posts, i)
			break
		}
	}

	update := bson.M{
		"$set": user,
	}

	_, err = userCollection.UpdateOne(ctx, userFilter, update)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Server Error Occurred",
			"error":   err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(res)
}
