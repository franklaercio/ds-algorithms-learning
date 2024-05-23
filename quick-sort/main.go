package main

import "fmt"

// QuickSort is a function that implements the QuickSort algorithm in Go.
// The QuickSort algorithm is a divide and conquer algorithm that has an average time complexity of O(n log n).
// However, in the worst case (when the input is already sorted), it can run at O(n²).
// This implementation of QuickSort uses the middle element of the input slice as the pivot.
//
// Parameters:
//   - numbers: A slice of integers that we want to sort.
//
// Returns:
//   - A new slice of integers that contains the elements of the input slice in ascending order.
func QuickSort(numbers []int) []int {
	// Base case: if the slice is empty or contains only one element, it is already sorted.
	if len(numbers) <= 1 {
		return numbers
	}

	// Choose the pivot. We use the middle element of the slice.
	pivot := numbers[len(numbers)%2]

	// Initialize two slices to hold the elements that are smaller and bigger than the pivot.
	var smaller []int
	var bigger []int

	// Divide the input slice into two parts: one with elements smaller than the pivot and one with elements bigger.
	for _, value := range numbers[1:] {
		if value < pivot {
			smaller = append(smaller, value)
		} else {
			bigger = append(bigger, value)
		}
	}

	// Recursively sort the smaller and bigger slices, and combine them with the pivot to get the sorted slice.
	return append(append(QuickSort(smaller), pivot), QuickSort(bigger)...)
}

// The main function of the program.
// It initializes a slice of integers, sorts it using the QuickSort function, and prints the sorted slice.
func main() {
	numbers := [8]int{3, 2, 1, 4, 6, 5, 6, 7}
	sorted := QuickSort(numbers[:])
	fmt.Println(sorted)
}
