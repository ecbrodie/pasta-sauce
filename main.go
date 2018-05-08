package main

import (
	"fmt"

	"github.com/ecbrodie/pasta-sauce/models"
)

func main() {
	evan := models.NewUser("Evan")
	fmt.Printf("Hello %s\n", evan.Name)
}
