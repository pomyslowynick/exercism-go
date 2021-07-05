package hamming

import (
	"errors"
	"strings"
)

// Distance calcualtes the Hamming distance between the two strings
func Distance(a, b string) (differences int, err error) {
	if len(a) != len(b) {
		return 0, errors.New("the strings are not the same length")
	}

	splitA := strings.Split(a, "")
	splitB := strings.Split(b, "")
	for index, letter := range splitA {
		if letter != splitB[index] {
			differences += 1
		}
	}
	// Go magic which might be less readable: it just returns the declared return variables
	// as they are after the program execution. err is initialize to nil and differences was
	// incremented to the correct value.
	return

	// Here is a more readable version I would actually prefer in production code.
	return differences, nil

}
