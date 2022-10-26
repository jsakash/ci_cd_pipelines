package main

import "testing"

func TestGreet(t *testing.T) {
	if Greet() != "Welcome to testing and integration" {
		t.Errorf("\nWanted - 'Welcome to testing and integration'\n got -")
	}
}

func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("/n/n Expected 4 ")
	}
}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
	}
	for _, tests := range tests {
		if output := Calculate(tests.input); output != tests.expected {
			t.Error("Test Failed: {} inputed, {} expected, {} received: {}", tests.input, tests.expected, output)
		}
	}
}
