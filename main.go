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

	client, err := NewIPAClient(config)
	if err != nil {
		fmt.Println("Error connection API", err)
		return
	}
	fmt.Println("Successful API connection")

	users, err := client.FindUsers("")
	if err != nil {
		fmt.Println("Receiving error finduser ")
		return
	}
	fmt.Println("Найдено пользователей:", len(users))
	for _, user := range users {
		fmt.Println("Login", user.Username, "Firstname", user.FirstName, "Lastname", user.LastName)
	}

	groups, err := client.FindGroups("")
	if err != nil {
		fmt.Println("Receiving error FindGroups")
	}
	fmt.Println("Найдено групп:", len(groups))
	for _, group := range groups {
		fmt.Println("Group", group.Name)
	}

	hosts, err := client.FindHosts("")
	if err != nil {
		fmt.Println("Receiving error FindHosts")
	}
	fmt.Println("Найдено хостов", len(hosts))
	for _, host := range hosts {
		fmt.Println("Host", host.FQDN)
	}

}
