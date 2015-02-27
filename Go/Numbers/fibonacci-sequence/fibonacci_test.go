package main

import "testing"

func TestFibOne(t *testing.T) {
    result := FibSeq(int64(1))
    expected := []int64{0}
    if len(result) != len(expected) {
        t.Error("Expected [0], got ", result)
    }
    for i, v := range result {
        if v != expected[i] {
            t.Error("Values did not match.")
        }
    }
}

func TestFibTwo(t *testing.T) {
    result := FibSeq(int64(2))
    expected := []int64{0, 1}
    if len(result) != len(expected) {
        t.Error("Expected [0 1], got ", result)
    }
    for i, v := range result {
        if v != expected[i] {
            t.Error("Values did not match.")
        }
    }
}

func TestFibThree(t *testing.T) {
    result := FibSeq(int64(6))
    expected := []int64{0, 1, 1, 2, 3, 5}
    if len(result) != len(expected) {
        t.Error("Expected [0 1 1 2 3 5], got ", result)
    }
    for i, v := range result {
        if v != expected[i] {
            t.Error("Values did not match.")
        }
    }
}