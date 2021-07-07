package main

import "fmt"

func func1(n int, ch chan int) {
	//sending into the channel
	ch <- n
}

func main() {
	// declaring a bidirectional channel that transports data of type int
	c := make(chan int)
	fmt.Printf("%T\n", c)

	// starting the goroutine
	go func1(10, c)

	//receiving data from the channel
	n := <-c
	fmt.Println("Value received:", n)

	fmt.Println("Exiting main ...")
}
