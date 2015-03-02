package main

import "fmt"

// import "os"

type Node struct {
	value string
	next  *Node
}

func main() {
	fmt.Println("Creating a node")
	n := Node{value: "test", next: nil}
	fmt.Println(n.value)
	otherNode := Node{value: "other", next: nil}
	n.next = &otherNode
	fmt.Println(n.next.value)
}
