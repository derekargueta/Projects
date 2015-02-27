package main

import "testing"

func TestOne(t *testing.T) {
    result := CollatzInit(2)
    if result != 1 {
        t.Error("Expected 1, got ", result)
    }
}

func TestTwo(t *testing.T) {
    result := CollatzInit(27)
    if result != 111 {
        t.Error("Expected 111, got ", result)
    }
}

func TestThree(t *testing.T) {
    result := CollatzInit(5)
    if result != 5 {
        t.Error("Expected 5, got ", result)
    }
}