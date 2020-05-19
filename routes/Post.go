package routes

import (
	"goblog/controllers"
	"goblog/mw"
	"net/http"

	"github.com/go-chi/chi"
)

// PostRoutes : routes for posts
func PostRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(mw.Auth)
	r.Get("/", controllers.GetPosts)
	r.Get("/user", controllers.GetUserPosts)
	r.Post("/", controllers.CreatePost)
	r.Patch("/{postId}", controllers.UpdatePost)
	r.Delete("/{postId}", controllers.DeletePost)

	return r
}
