package main

import "testing"

func TestBinaryNumberValidation(t *testing.T) {
	if validateBinaryArgument("asdf") {
		t.Error("Expected asdf to be invalid, but passed as valid")
	}

	if validateBinaryArgument("12345") {
		t.Error("Expected 12345 to be invalid, but passed as valid")
	}

	if validateBinaryArgument(",.'[;'") {
		t.Error("Expected ,.'[;' to be invalid, but passed as valid")
	}

	if !validateBinaryArgument("101") {
		t.Error("Expected 101 to be valid, but failed as invalid")
	}

	if !validateBinaryArgument("0") {
		t.Error("Expected 0 to be valid, but failed as invalid")
	}

	if !validateBinaryArgument("1") {
		t.Error("Expected 1 to be valid, but failed as invalid")
	}

	if validateBinaryArgument("-1") {
		t.Error("Expected -1 to be invalid, but passed as valid")
	}
}

func TestDecimalNumberValidation(t *testing.T) {
	if validateDecimalArgument("asdf") {
		t.Error("Expected asdf to be invalid, but passed as valid")
	}

	if !validateDecimalArgument("12345") {
		t.Error("Expected 12345 to be valid, but failed as invalid")
	}

	if validateDecimalArgument(",.'[;'") {
		t.Error("Expected ,.'[;' to be invalid, but passed as valid")
	}

	if !validateDecimalArgument("101") {
		t.Error("Expected 101 to be valid, but failed as invalid")
	}

	if !validateDecimalArgument("0") {
		t.Error("Expected 0 to be valid, but failed as invalid")
	}

	if !validateDecimalArgument("1") {
		t.Error("Expected 1 to be valid, but failed as invalid")
	}

	if !validateDecimalArgument("-1") {
		t.Error("Expected -1 to be valid, but failed as invalid")
	}

	if validateDecimalArgument("123aab") {
		t.Error("Expected 123aab to be invalid, but passed as valid")
	}
}
