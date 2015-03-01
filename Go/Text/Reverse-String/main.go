package main

import "fmt"
import "os"

func Reverse(s string) string {

	// convert string to array of runes
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please enter a string to reverse")
		return
	}

	fmt.Println(Reverse(os.Args[1]))
}
