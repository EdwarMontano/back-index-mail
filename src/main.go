package main

import (
	"fmt"
	"net/http"

	"github.com/EdwarMontano/back-index-mail/src/httpd/handler"
	"github.com/EdwarMontano/back-index-mail/src/platform/enronmail"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func main() {
	port := ":3001"
	feed := enronmail.New()
	feed.Add(enronmail.Item{
		IdMsg:      "Hello Aliens",
		DateMsg:    "How are you",
		FromMsg:    "from  test",
		ToMsg:      "to test",
		SubjectMsg: "test",
		CcMsg:      "test",
		BccMsg:     "test",
		XFromMsg:   "test",
		XToMsg:     "test",
		XccMsg:     "test",
		XbccMsg:    "test",
	})

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("ruta no existe"))
	})
	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		w.Write([]byte("Método no valido"))
	})
	r.Get("/mailstest", handler.MockMailGet(feed))
	r.Post("/mailstest", handler.MockMailPost(feed))

	fmt.Println("Servidor en el " + port)
	http.ListenAndServe(port, r)
}

func AllowOriginFunc(r *http.Request, origin string) bool {
	if origin == "http://127.0.0.1:5173/" {
		return true
	}
	return false
}
