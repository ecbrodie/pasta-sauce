package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ecbrodie/pasta-sauce/models"
)

func handler(w http.ResponseWriter, r *http.Request) {
	evan := models.NewUser("Evan")
	fmt.Fprintf(w, "Hello %s\n", evan.Name)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
