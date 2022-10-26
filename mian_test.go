package main

import "testing"

func TestGreet(t *testing.T) {
	//initial := Greet()
	if Greet() != "Welcome to testing and integration" {
		t.Errorf("\nWanted - 'Welcome to testing and integration'\n got -")
	}
}

func TestCalculate(t *testing.T) {

}
