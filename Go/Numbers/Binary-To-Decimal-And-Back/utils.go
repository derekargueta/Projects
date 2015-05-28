package main

import "strconv"

// iterate through all the runes in the string and ensure they are either 1 or 0
func validateBinaryArgument(binaryArg string) bool {
	for _, char := range binaryArg {
		if char != '0' && char != '1' {
			return false
		}
	}
	return true
}

func validateDecimalArgument(decimalArg string) bool {
	// TODO - validate all characters are 1-9
	if _, err := strconv.Atoi(decimalArg); err == nil {
		return true
	}
	return false
}
