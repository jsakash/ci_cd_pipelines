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
