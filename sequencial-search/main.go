package main

import "fmt"

func main() {
	contactList := [6]string{"Ohanna", "Bia", "Jackson", "Elizabete", "Olivia", "Lidiane"}

	// This code represents a linear search in O (n)
	// In this case we're leaving 5 steps to found Lidiane
	for _, contact := range contactList {
		if contact == "Lidiane" {
			fmt.Println("We're found your contact:", contact)
		}
	}
}
