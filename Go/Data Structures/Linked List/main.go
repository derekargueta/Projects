package main

import "fmt"
import "os"
import "bufio"

type Node struct {
	value string
	next  *Node
}

type LinkedList struct {
	length uint
	head   *Node
}

func AddNodeToList(list *LinkedList, node *Node) {

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

func PrintList(list LinkedList) {
	node := list.head
	for node.next != nil {
		fmt.Println(node.value)
		node = node.next
	}

	// and print the last node
	fmt.Println(node.value)
}

func main() {

	// create list
	list := LinkedList{length: 0, head: nil}

	// list.head = &n
	n := Node{value: "test", next: nil}
	otherNode := Node{value: "other", next: nil}
	// n.next = &otherNode

	// testing incrementor
	AddNodeToList(&list, &n)
	AddNodeToList(&list, &otherNode)
	fmt.Println("List length: ", list.length)
	PrintList(list)
	// AND IT WORKS

	// Create a way for users to input command to list
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a command, or 'help' for a list of commands: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
