package routes

import (
	"goblog/controllers"
	"goblog/mw"
	"net/http"

	"github.com/go-chi/chi"
)

// UserRoutes : routes for users
func UserRoutes() http.Handler {
	r := chi.NewRouter()
	r.Post("/register", controllers.RegisterUser)
	r.Post("/", controllers.LoginUser)

	return r
}

// ProtectedUserRoutes : JWT protected user routes
func ProtectedUserRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(mw.Auth)
	r.Get("/{username}", controllers.GetUser)
	r.Delete("/{username}", controllers.DeleteUser)

	return r
}

// UploadRoutes : Upload routes for images
func UploadRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(mw.Auth)
	r.Post("/upload", controllers.GetUser)

	return r
}
