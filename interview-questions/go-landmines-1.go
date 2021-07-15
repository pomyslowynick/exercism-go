package main

import (
	"bytes"
	"fmt"
)

func print(pi *int) { fmt.Println(*pi) } // prints whateber pi points to

func main() {
	for i := 0; i < 10; i++ { // iteration from 0..9
		// defer fmt.Println(i) // prints 0..9? === works correctly, prints from 9..0 because of the reverse order
		defer func() { fmt.Println(i) }() // prints 0? !== prints 10 ten times
		// defer func(i int) { fmt.Println(i) }(i) // prints 0?
		// defer print(&i)                         // prints &i..&i, the address shouldn't change?
		// go fmt.Println(i)                       // prints 0..9
		// go func() { fmt.Println(i) }()          // prints 0..9
	}

	// Empty interface can be of any type
	var any interface{}
	any = true
	any = 12.34
	any = map[string]int{"one": 1}
	any = new(bytes.Buffer)
	any = "hello"

	fmt.Println(any)
}
