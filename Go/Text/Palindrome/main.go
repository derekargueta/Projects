package main

import "fmt"
import "os"

func PalindromeCheck(s string) bool {

	// convert string to array of runes
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}

	return true
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please enter a string to check for being a palindrome")
	}

	fmt.Println(PalindromeCheck(os.Args[1]))
}
