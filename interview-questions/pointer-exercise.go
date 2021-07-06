// https://blog.udemy.com/golang-interview-questions/
// Question 20
package main

import "fmt"

func main() {
	var someString string
	fmt.Println(&someString)

	var someInt int
	var intPointer *int = &someInt
	fmt.Println(intPointer)
}
