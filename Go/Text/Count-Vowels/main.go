package main

import "fmt"
import "os"
import "bufio"

func CountVowels(s string) map[rune]int {

	letterCount := make(map[rune]int)
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}

	// iterate through input string
	for _, r := range s {

		// iterate through vowels
		for _, vowel := range vowels {

			// If current letter is a vowel...
			if r == vowel {
				if _, ok := letterCount[r]; ok {
					// increment it in the map
					letterCount[vowel]++
				} else {
					// or instantiate it with '1'
					letterCount[vowel] = 1
				}

				// no need to keep checking vowels
				break
			}
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
