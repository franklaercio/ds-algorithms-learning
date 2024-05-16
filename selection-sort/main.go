package main

import "fmt"

func main() {
	// Unsorted numbers lists
	numbers := [10]int{5, 7, 8, 9, 1, 3, 2, 4, 6, 0}

	// We're using the hard way to sort one list
	// Using select one minor and change your position
	// This algorithm runs at O (n²)
	for i := 0; i < len(numbers); i++ {
		selectedNumber := numbers[i]
		pivotIndex := i

		for j := i; j < len(numbers); j++ {
			if numbers[j] < selectedNumber {
				selectedNumber = numbers[j]
				pivotIndex = j
			}
		}

		numbers[pivotIndex] = numbers[i]
		numbers[i] = selectedNumber
	}

	fmt.Println(numbers)
}
