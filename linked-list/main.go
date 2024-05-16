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
	fmt.Print("Insert your value -> ")

	var value int
	if _, err := fmt.Scanf("%d", &value); err != nil {
		fmt.Println("Invalid input, try again")
		list.Add()
	}

	newNode := &Node{value, nil}
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	}

	list.tail.next = newNode
	list.tail = newNode
	list.size++
}

func (list *LinkedList) Remove() {
	fmt.Print("Insert your value that should be removed -> ")

	var value int
	if _, err := fmt.Scanf("%d", &value); err != nil {
		fmt.Println("Invalid input, try again")
		list.Add()
	}

	if list.head == nil {
		fmt.Println("list is empty")
	}

	for {
		// TODO: change to removing just a tail value to removing in O(1)
		if list.head.value == value {
			list.head = list.head.next
			list.size--

			if list.size == 0 {
				list.tail = nil
			}

			break
		}

		if list.head.next == nil {
			break
		}

		list.head = list.head.next
	}
}

func main() {
	list := LinkedList{}

	for {
		fmt.Println("1 - Insert a value")
		fmt.Println("2 - Remove a value")
		fmt.Println("3 - Display list")

		var option int
		if _, err := fmt.Scanf("%d", &option); err != nil {
			fmt.Println("Invalid option")
			continue
		}

		switch option {
		case 1:
			list.Add()
		case 2:
			list.Remove()
		}
	}
}
