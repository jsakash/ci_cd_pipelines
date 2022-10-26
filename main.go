package main

import "fmt"

func main() {
	initial := Greet()
	fmt.Println(initial)
}

func Greet() string {
	return "Welcome to testing and integration"
}

func Calculate(x int) (result int) {

	result = x + 2
	return result
}
