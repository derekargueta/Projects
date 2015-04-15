package main

import "fmt"

func main() {

	// create list
	list := LinkedList{length: 0, head: nil}

	// list.head = &n
	n := Node{value: "test", next: nil}
	otherNode := Node{value: "other", next: nil}
	// n.next = &otherNode

	// testing incrementor
	list.AddNode(&n)
	list.AddNode(&otherNode)
	fmt.Println("List length: ", list.length)
	list.Print()

	list.RemoveNode("test")
	list.Print()
	// AND IT WORKS

	// for true {
	// 	// Create a way for users to input command to list
	// 	//reader := bufio.NewReader(os.Stdin)
	// 	fmt.Print("Enter a command, or 'help' for a list of commands: ")
	// 	//text, _ := reader.ReadString('\n')
	// }
}
