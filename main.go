package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ecbrodie/pasta-sauce/models"
	"github.com/gorilla/mux"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	evan := models.NewUser("Evan")
	fmt.Fprintf(w, "Hello %s\n", evan.Name)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
