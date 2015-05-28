package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter a sentence to count: ")
	reader := bufio.NewReader(os.Stdin)
	sentence, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading the sentence: ", err)
		os.Exit(-1)
	}

	sentence = strings.Replace(sentence, "\n", "", 0)

	wordList := strings.Split(sentence, " ")
	fmt.Println(len(wordList))
}
