package main

//this file will create the server and tell path of routes file

import (
	"log"
	"net/http"

	"book-management-api.io/packages/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()              //instance of Router
	routes.RegisterBookStoreRoutes(r) //calling router package
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:1552", r)) //starting server
}
