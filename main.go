package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Failed to load .env", err)
	}

	config := IPAConfig{
		Host:     "ipa.example.test",
		Username: os.Getenv("FREEIPA_USERNAME"),
		Password: os.Getenv("FREEIPA_PASSWORD"),
		Insecure: true,
	}

	_, err := NewIPAClient(config)
	if err != nil {
		fmt.Println("Error connection API", err)
		return
	}
	fmt.Println("Successful API connection")
}
