package main

import "testing"

/**
 * ---------- Test Round function
 */
func TestRoundOne(t *testing.T) {
    untruncated := 3.445
    truncated := Round(untruncated, .5, 2)
    if truncated != 3.45 {
        t.Error("Expected 3.45, got ", truncated)
    }
}

func TestRoundTwo(t *testing.T) {
    untruncated := 2.00019
    truncated := Round(untruncated, .5, 4)
    if truncated != 2.0002 {
        t.Error("Expected 2.0002, got ", truncated)
    }
}

/**
 * ---------- Test Reducer function
 */
 func TestReducerOne(t *testing.T) {
    result := Reducer("", 20.00, 20.00)
    if result != 0 {
        t.Error("Expected 0, got ", result)
    }
 }

 func TestReducerTwo(t *testing.T) {
    result := Reducer("", 35, 15)
    if result != 5 {
        t.Error("Expected 5, got ", result)
    }
 }

 func TestReducerThree(t *testing.T) {
    result := Reducer("", 10, 20)
    if result != 10 {
        t.Error("Expected 10, got ", result)
    }
 }

 func TestReducerFour(t *testing.T) {
    result := Reducer("", 20, -5)
    if result != 20 {
        t.Error("Expected 20, got ", result)
    }
 }
