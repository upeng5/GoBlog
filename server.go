package main

import (
	"fmt"
	"goblog/database"
	"goblog/routes"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Starting server...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongodbURI := os.Getenv("MONGODB_URI")
	// fmt.Println("MongoDB URI:", mongodbURI)

	database.ConnectDB(mongodbURI)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// User routes
	r.Mount("/api/v1/users", routes.UserRoutes())
	// Protected user routes
	r.Mount("/api/v1/protected/users", routes.ProtectedUserRoutes())
	// Post routes
	r.Mount("/api/v1/posts", routes.PostRoutes())

	http.ListenAndServe(":3000", r)
}
