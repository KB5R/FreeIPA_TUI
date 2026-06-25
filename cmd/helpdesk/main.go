package main

import (
	"log"

	helpdeskui "freeipa-tui/internal/helpdesk/ui"
)

func main() {
	if err := helpdeskui.Run(); err != nil {
		log.Fatal(err)
	}
}

/*
Старый проверочный запуск FreeIPA.
Оставляем пока тут, чтобы можно было быстро подсмотреть, как мы проверяли backend.

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	helpdesk "freeipa-tui/internal/helpdesk/backend"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Failed to load .env", err)
	}

	config := helpdesk.IPAConfig{
		Host:     "ipa.example.test",
		Username: os.Getenv("FREEIPA_USERNAME"),
		Password: os.Getenv("FREEIPA_PASSWORD"),
		Insecure: true,
	}

	client, err := helpdesk.NewIPAClient(config)
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
}
*/
