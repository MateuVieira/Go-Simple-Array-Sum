package main

import (
	"testing"
)

func testSimpleArraySum(t *testing.T) {
	arTestOne := []int{1, 2, 3, 4, 10, 11}
	arTestTwo := []int{-1, 2, -3, 4, 10, -11}
	arTestThree := []int{-1, -2, -3, -4, 10, -1}

	if simpleArraySum(arTestOne) != 31 {
		t.Error("Expected 1 + 2 + 3 + 4 + 10 + 11 to equal 31 ")
	}

	if simpleArraySum(arTestTwo) != 1 {
		t.Error("Expected -1 + 2 - 3 + 4 + 10 - 11 to equal 1 ")
	}

	if simpleArraySum(arTestThree) != -1 {
		t.Error("Expected 10 - 1 - 2 - 3 - 4 - 1 to equal -1 ")
	}
}
