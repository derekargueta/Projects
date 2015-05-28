package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// Did this one the hard way for the LULz
func convBinary2Decimal(num string) string {

	// will catch if the string is invalid
	if !validateBinaryArgument(num) {
		fmt.Println("Invalid binary number")
		os.Exit(1)
	}

	var totalDecNum int

	for i, char := range num {
		// grab 2**i
		multiFactor := math.Exp2(float64(i))
		// convert the char rune into an integer
		charNumVal := int(char - '0')
		// add to the total
		totalDecNum += charNumVal * int(multiFactor)
	}

	return strconv.Itoa(totalDecNum)
}

// Cheated and used built-in functionality
func convDecimal2Binary(num string) string {

	// will catch if the string is invalid
	if !validateDecimalArgument(num) {
		fmt.Println("Invalid decimal number")
		os.Exit(1)
	}

	n, _ := strconv.Atoi(num)

	// var total
	return strconv.FormatInt(int64(n), 2)
}

func checkArgs() {
	if len(os.Args) < 3 {
		fmt.Println("Not enough arguments! Provide a method (bin|dec) and value")
		os.Exit(1)
	} else if len(os.Args) > 3 {
		fmt.Println("Too many arguments!")
		os.Exit(1)
	}
}

func main() {

	checkArgs()

	var result string
	if os.Args[1] == "bin" {
		// convert last argument decimal -> binary
		result = convDecimal2Binary(os.Args[2])
	} else if os.Args[1] == "dec" {
		// convert last argument binary -> decimal
		result = convBinary2Decimal(os.Args[2])
	} else {
		fmt.Println("Bad method, please use `bin` or `dec`")
		os.Exit(1)
	}

	fmt.Println("result: " + result)
}
