package main

import "testing"

func TestOne(t *testing.T) {
	result := Reverse("teststring")
	if result != "gnirtstset" {
		t.Error("Was expecting gnirtstset, got ", result)
	}
}

func TestTwo(t *testing.T) {
	result := Reverse("itworks")
	if result != "skrowti" {
		t.Error("Was expecting skrowti, got ", result)
	}
}
