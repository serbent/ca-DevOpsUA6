package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	result := 2 + 3
	expected := 5

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestMultiply(t *testing.T) {
	result := 4 * 5
	expected := 20

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestSquare(t *testing.T) {
	result := 6 * 6
	expected := 36

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}
