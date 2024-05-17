package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

// Add is a function that insert a value in O(1), because inserting a some value in linked list is easy
func (list *LinkedList) Add() {
	fmt.Print("Insert your value: ")

	var value int
	if _, err := fmt.Scanf("%d", &value); err != nil {
		fmt.Println("Invalid input, try again")
		list.Add()
	}

	newNode := &Node{value, nil}
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		list.tail.next = newNode
		list.tail = newNode
	}

	list.size++
}

// Remove is a function that remove a value in O(1), because removing a some value in linked list is easy in head
func (list *LinkedList) Remove() {
	if list.head == nil {
		fmt.Println("list is empty")
		return
	}

	fmt.Println("Removing:", list.head.value)

	if list.head == list.tail {
		list.head = nil
		list.tail = nil
	} else {
		list.head = list.head.next
		list.tail.next = nil
	}
	list.size--
}

// Display is a function that display all values in linked list in O(n)
func (list *LinkedList) Display() {
	if list.head == nil {
		fmt.Println("List is empty")
		return
	}

	fmt.Printf("There are %d in list: ", list.size)
	current := list.head
	for current != nil {
		fmt.Print(current.value)
		if current.next != nil {
			fmt.Print(" -> ")
		}
		current = current.next
	}
	fmt.Println()
}

func main() {
	list := LinkedList{}

	for {
		fmt.Println("\n1 - Insert a value")
		fmt.Println("2 - Remove a value")
		fmt.Println("3 - Display list")
		fmt.Println("4 - Exit")

		var option int
		fmt.Print("\nInsert your option: ")

		if _, err := fmt.Scanf("%d", &option); err != nil {
			fmt.Println("\nInvalid option")
			panic(err)
		}

		switch option {
		case 1:
			list.Add()
		case 2:
			list.Remove()
		case 3:
			list.Display()
		case 4:
			return
		}
	}
}
