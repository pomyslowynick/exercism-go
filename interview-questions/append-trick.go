// https://blog.udemy.com/golang-interview-questions/
// Question 12
package main

import "fmt"

func main() {
	// n1 is a slice of integers
	n1 := []int{10, 20, 30, 40}
	// append should return a slice {10, 20, 30, 40}
	n1 = append(n1, 100)
	// I see where the question is going, if I remember correctly from the gopl capacity is gonna be set to n*2
	// when n1 was initiated, it's not gonna increase with a single append.
	fmt.Println(len(n1), cap(n1)) // 5 8
}
