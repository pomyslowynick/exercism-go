package main

import "fmt"

// That is some seriously twisted code.
// My understanding of it before I'll read what it does:
// adder() is a regular function, it's body is an anonymous function which takes an integer argument
// and returns an integer. It also has a sum variable within it's scope and returns another anonymous function
// which takes a named argument x, type int and returns an int. It also accesses an outer scope variable scope and adds
// the x argument to it and then returns the value of the outer variable.

// UPDATE: okay, well it's clear right now. it isn't a nested anonymous function, adders return type is a func(int) int.
// That makes a lot more sense :D
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// pos and neg are pointing to anonymous function within adder
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		// at each iteration of the loop both pos and neg are called with increasing i values
		// my understanding is that both pos and neg keep the values of variables within their respective scopes
		// in a way they "save the functio environment". Hence each time sum value is incremented it will keep it's value
		// within the scope of the function.

		// adder returns an anonymous function which is then called (is it anonymous if it has a variable pointing to it? :P)
		// I wonder if we need the outer anonymous function.
		fmt.Println(
			pos(i),    // I predict: 0, 1, 3, 6, 10, 15, 21, 28, 36, 45
			neg(-2*i), // prediction: 0, -2, -6, -12, -20, -30, -42, -56, -72, -90
		)
	}
}
