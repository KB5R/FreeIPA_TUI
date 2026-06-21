package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("Начало")
	parseUser()
}

// Обработка юзеров
func parseUser() {
	cmd := exec.Command("ipa", "user-find", "--all", "--sizelimit=9999999")

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(out))
}

func parserUserGroup() {
	cmd := exec.Command("ipa", "group-find", "--all", "--sizelimit=9999999")

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(out))
}

// Обработка хостов
func parseHost() {
	cmd := exec.Command("ipa", "host-find", "--all", "--sizelimit=9999999")

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(out))
}

func parseHostGroup() {
	cmd := exec.Command("ipa", "hostgroup-find", "--all", "--sizelimit=9999999")

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(out))
}

// Обработка HBAC

func parseHbacRule() {
	cmd := exec.Command("ipa", "hbacrule-find", "--all", "--sizelimit=9999999")

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(out))
}

// Обработка Sudo

func parseSudoRule() {
	cmd := exec.Command("ipa", "sudorule-find", "--all", "--sizelimit=9999999")

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(out))
}
