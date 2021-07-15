package main

import (
	"golangorm/config"
	"golangorm/db"
	"golangorm/rest"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Not able to load environment file")

	}

	err = config.InitConfig()

	if err != nil {
		log.Fatal("Not able to create config")

	}

	db.InitDatabase()
	db.InitData()

	router := chi.NewRouter()
	router.Route("/"+config.GetConfig().APPVersion, func(r chi.Router) {
		r.Get("/person", rest.GetPerson)
		r.Post("/person", rest.PostPerson)
		r.Delete("/person", rest.DeletePerson)
		r.Post("/testlogin", rest.Login)
		r.Get("/welcome", rest.Welcome)
		r.Get("/logout", rest.Logout)
	})
	http.ListenAndServe(":"+config.GetConfig().Port, router)
}
