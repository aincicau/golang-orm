package main

import (
	"golangorm/db"
	"golangorm/rest"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	db.InitDatabase("postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable", 1)

	router := chi.NewRouter()
	router.Route("/v1", func(r chi.Router) {
		r.Get("/person", rest.GetPerson)
		r.Post("/person", rest.PostPerson)
		r.Delete("/person", rest.DeletePerson)
	})
	http.ListenAndServe(":8080", router)
}
