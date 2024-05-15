package main

import "fmt"

func main() {
	// This list should be ordering, because this algorithm only work with this condition
	contactList := [10]string{"Ana", "Bia", "Frank", "Jackson", "Ohanna", "Olivia", "Lidiane", "Tania", "Xuxa", "Zoe"}

	// This code represents a logarithm search in O (log n)
	smallest := 0
	greatest := len(contactList) - 1

	for {
		middle := (smallest + greatest) / 2 // Every execution is change to the middle of list
		contact := contactList[middle]

		// In this algorithm we're leaving 2 executions, instead 7 of sequential search
		if contact == "Tania" {
			fmt.Println("We're found your contact:", contact)
			break
		}

		if contact > contactList[middle] {
			greatest = middle - 1 // Move the pointer to the middle lower
		} else {
			smallest = middle + 1 // Move the pointer to the middle upper
		}
	}
}
