package main

import "testing"

func TestOne(t *testing.T) {
	if !PalindromeCheck("racecar") {
		t.Error("Was expecting true, got false")
	}
}

func TestTwo(t *testing.T) {
	if !PalindromeCheck("tacocat") {
		t.Error("Was expecting true, got false")
	}
}

func TestThree(t *testing.T) {
	if !PalindromeCheck("~ - | * | - ~") {
		t.Error("Was expecting true, got false")
	}
}

func TestFour(t *testing.T) {
	if PalindromeCheck("not a palindrome") {
		t.Error("Was expecting false, got true")
	}
}

func TestFive(t *testing.T) {
	if PalindromeCheck("what what") {
		t.Error("Was expecting false, got true")
	}
}

func TestEight(t *testing.T) {
	if !PalindromeCheck("evenneve") {
		t.Error("Was expecting true, got false")
	}
}

func TestNine(t *testing.T) {
	if !PalindromeCheck("") {
		t.Error("Was expecting true, got false")
	}
}
