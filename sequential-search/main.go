package main

import "fmt"

func main() {
	// Considering ordering this list of contacts
	contactList := [10]string{"Ana", "Bia", "Frank", "Jackson", "Ohanna", "Olivia", "Lidiane", "Tania", "Xuxa", "Zoe"}

	// This code represents a linear search in O (n)
	// In this case we're leaving 5 steps to found Lidiane
	for _, contact := range contactList {
		if contact == "Tania" {
			fmt.Println("We're found your contact:", contact)
		}
	}
}
