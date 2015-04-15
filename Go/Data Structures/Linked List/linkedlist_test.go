package main

import "testing"

func TestAddToList(t *testing.T) {
	ll := LinkedList{length: 0, head: nil}
	testNode1 := Node{value: "test", next: nil}
	ll.AddNode(&testNode1)
	if ll.length != 1 {
		t.Error("Expected length of one")
	}
}
