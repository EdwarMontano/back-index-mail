package main

import (
	"chi-example/httpd/handler"
	"chi-example/platform/newsfeed"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	// r := chi.NewRouter()
	// r.Use(middleware.Logger)
	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello Aliens World!"))
	// })
	// http.ListenAndServe(":3000", r)
	port := ":3000"
	feed := newsfeed.New()

	r := chi.NewRouter()

	r.Get("/newsfeed", handler.NewsfeedGet(feed))
	// r.Post("/newsfeed", handler.NewsfeedPost(feed))

	fmt.Println("Serving on " + port)
	http.ListenAndServe(port, r)
}
