package main

import "fmt"
import "os"
import "bufio"

func CountVowels(s string) map[rune]int {

	letterCount := map[rune]int{
		'a': 0,
		'e': 0,
		'i': 0,
		'o': 0,
		'u': 0,
	}

	// iterate through input string
	for _, r := range s {

		if _, ok := letterCount[r]; ok {
			// increment it in the map
			letterCount[r]++
		}
	}

	return letterCount
}

func main() {
	// setup input reader
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a word: ")
	text, _ := reader.ReadString('\n')
	letterCount := CountVowels(text)

	// print the results
	for k, v := range letterCount {
		fmt.Printf("%s -- %d\n", string(k), v)
	}
}
