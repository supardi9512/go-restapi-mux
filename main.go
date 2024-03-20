package main

import (
	"go-restapi-mux/controllers/productcontroller"
	"go-restapi-mux/models"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	models.ConnectDatabase()

	r := mux.NewRouter()

	r.HandleFunc("/products", productcontroller.Index).Methods("GET")

	http.ListenAndServe(":8080", r)
}
