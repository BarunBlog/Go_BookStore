package main

import (
	"log"
	"net/http"

	"github.com/BarunBlog/Go_BookStore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	log.Fatal(http.ListenAndServe("127.0.0.1:4000", r))
}
