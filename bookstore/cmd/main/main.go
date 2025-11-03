package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PremModhaOfficial/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("localhost:8081")
	r := mux.NewRouter()

	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8081", r))
}
