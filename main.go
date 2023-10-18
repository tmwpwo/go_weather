package main

import (
	"fmt"
	"net/http"
	"weather/render"
	"weather/routes"

	"github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()

	r.Get("/",
		routes.StaticHandler(
			render.RenderError(
				render.Parse("templates/home.html"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)

}
