package main

import (
	"agro/database"
	"agro/server/http/route"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	DB, err := database.InitDBConnection("postgres", "postgres://postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()
	apiRouterProvider := route.APIProvider{}
	apiRouterProvider.InitComponentsUsed(DB)
	router := apiRouterProvider.GetRouter()

	log.Fatal(http.ListenAndServe(":8082", router))
}
