package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Nurmanbetov/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// main function, runs the app on localhost:9010
func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Printf("Starting server at port 9010")
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}
