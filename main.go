package main

import (
	"fmt"
	"log"
	"net/http"

	. "star-wars-api/config"
	planetController "star-wars-api/controllers"
	. "star-wars-api/repositories"

	"github.com/gorilla/mux"
)

var repository = PlanetRepository{}
var config = Config{}

func init() {
	config.Read()

	repository.Server = config.Server
	repository.Database = config.Database
	repository.Connect()
}

func routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/planets", planetController.GetAll).Methods("GET")
	r.HandleFunc("/planets/{id}", planetController.GetByID).Methods("GET")
	// r.HandleFunc("/planets/{name}", planetController.GetByName).Methods("GET")
	r.HandleFunc("/planets", planetController.Create).Methods("POST")
	r.HandleFunc("/planets/{id}", planetController.Update).Methods("PUT")
	r.HandleFunc("/planets/{id}", planetController.Delete).Methods("DELETE")

	return r
}

func main() {
	r := routes()

	var port = ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}
