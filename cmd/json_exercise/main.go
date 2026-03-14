package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/minhvq36/go-learning/internal/models"
)

func main() {
	user := models.User{
		ID:    123,
		Name:  "Mr. A",
		Email: "A@example.com",
	}

	jsonData, err := json.Marshal(user)

	if err != nil {
		log.Fatalf("Error encoding JSON: %v\n", err)
	} else {
		fmt.Printf("Struct to JSON string: %s\n", string(jsonData))
	}

	jsonInput := `{"id": 141, "name": "DKT", "email": "B@example.com"}`

	var newUser models.User
	err = json.Unmarshal([]byte(jsonInput), &newUser)

	if err != nil {
		log.Fatalf("Error decoding JSON: %v\n", err)
	} else {
		fmt.Printf("Decoded struct: %+v\n", newUser)
	}
}
