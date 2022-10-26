package main

import "testing"

func TestMain(t *testing.T) {
	initial := Greet()
	if initial != "Welcome to testing and integration" {
		t.Errorf("\nWanted - 'Welcome to testing and integration'\n got - %s\n\n", initial)
	}

}
