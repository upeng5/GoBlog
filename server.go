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
	"github.com/go-chi/cors"

	"github.com/joho/godotenv"
)

var (
	corsOptions cors.Options = cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Requested-With", "x-auth-token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}
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
	// Enable CORS
	r.Use(cors.Handler(corsOptions))

	// User routes
	r.Mount("/api/v1/users", routes.UserRoutes())
	// Protected user routes
	r.Mount("/api/v1/protected/users", routes.ProtectedUserRoutes())
	// Post routes
	r.Mount("/api/v1/posts", routes.PostRoutes())

	http.ListenAndServe(":3000", r)
}
