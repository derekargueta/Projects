package main

import "fmt"

type Node struct {
	value string
	next  *Node
}

type LinkedList struct {
	length uint
	head   *Node
}

func (list *LinkedList) AddNode(node *Node) {

	list.length++
	if list.head == nil {
		list.head = node
		return
	}

	iterNode := list.head
	for iterNode.next != nil {
		iterNode = iterNode.next
	}
	iterNode.next = node
}

func (list *LinkedList) RemoveNode(value string) {
	if list.head.value == value {
		list.head = list.head.next
		list.length--
		return
	}

	iterNode := list.head
	for iterNode.next.value != value {
		iterNode = iterNode.next
	}

	iterNode = iterNode.next
	list.length--
}

func (list LinkedList) Print() {
	node := list.head
	for node.next != nil {
		fmt.Println(node.value)
		node = node.next
	}

	// and print the last node
	fmt.Println(node.value)
}
