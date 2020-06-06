package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"goblog/database"
	"goblog/models"
	"goblog/utils"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type registerUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registerUserResponse struct {
	Token  string                 `json:"token"`
	Result *mongo.InsertOneResult `json:"result"`
}

type loginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type customClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

var (
	defaultURL                    string = "uploads/default/default_profile.png"
	defaultAvatarBytes, readError        = ioutil.ReadFile(defaultURL)
	dbName                        string = "GoBlog"
)

// RegisterUser : Creates user and signs a json web token
func RegisterUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client := database.GetDB()
	if readError != nil {
		defaultAvatarBytes = []byte{}
	}
	var requestBody registerUserRequest

	json.NewDecoder(req.Body).Decode(&requestBody)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := client.Database(dbName).Collection("users")

	// check if user already exists in the database
	filter := bson.M{"email": requestBody.Email}

	found := collection.FindOne(ctx, filter)

	if found.Err() == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Email is taken",
		})
		return
	}

	hashedPassword, err := utils.HashPassword(requestBody.Password, 10)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	userID := primitive.NewObjectID()
	defaultURL := utils.DefaultAvatar()

	user := models.User{
		ID:       userID,
		Username: requestBody.Username,
		Email:    requestBody.Email,
		Password: hashedPassword,
		Avatar:   defaultURL,
	}

	// TODO fix _id issue
	res, err := collection.InsertOne(ctx, user)

	token, err := utils.GenerateToken(user.Username)
	fmt.Println("DEBUG:", user.ID.String())
	// Sign and get the complete encoded token as a string using the secret
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	response := registerUserResponse{
		Token:  token,
		Result: res,
	}

	req.Header.Set("x-auth-token", token)
	json.NewEncoder(w).Encode(response)

}

// LoginUser : Logs in the user and signs a json web token
func LoginUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client := database.GetDB()
	var requestBody loginUserRequest
	var user models.User

	json.NewDecoder(req.Body).Decode(&requestBody)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := client.Database(dbName).Collection("users")

	filter := bson.M{"email": requestBody.Email}

	res := collection.FindOne(ctx, filter)

	if res.Err() != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid credentials",
		})
		return
	}

	res.Decode(&user)

	isMatch := utils.CheckPasswordHash(requestBody.Password, user.Password)

	if isMatch == false {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid credentials",
		})
		return
	}

	token, err := utils.GenerateToken(user.Username)
	// Sign and get the complete encoded token as a string using the secret
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	req.Header.Set("x-auth-token", token)
	json.NewEncoder(w).Encode(token)

}

// GetUser : Gets the user by username
func GetUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client := database.GetDB()
	username := chi.URLParam(req, "username")
	var user models.User

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := client.Database(dbName).Collection("users")

	filter := bson.M{"username": username}

	res := collection.FindOne(ctx, filter)

	if res.Err() != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid credentials",
		})
		return
	}

	res.Decode(&user)

	json.NewEncoder(w).Encode(user)
}

// DeleteUser : Deletes the user by the id
func DeleteUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client := database.GetDB()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := client.Database(dbName).Collection("users")

	filter := bson.M{"username": req.Header.Get("user-name")}

	res := collection.FindOneAndDelete(ctx, filter)

	if res.Err() != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{
			"message": res.Err().Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(res)

}
