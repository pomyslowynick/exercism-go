package reverse

import "strings"

// func Reverse(input string) string {
// 	splitString := strings.Split(input, "")
// 	var outString []string

// 	for index := range splitString {
// 		letterIndex := len(splitString) - 1 - index
// 		outString = append(outString, splitString[letterIndex])
// 	}

// 	return strings.Join(outString, "")
// }

func Reverse(s string) string {
	if s != "" {
		splitString := strings.Split(s, "")

		for i, j := len(s)-1, 0; j < i; i, j = i-1, j+1 {
			splitString[i], splitString[j] = splitString[j], splitString[i]
		}
		return strings.Join(splitString, "")
	}

	return ""
}

// func Reverse(s string) string {
// 	r := []rune(s)
// 	left := 0
// 	right := len(r) - 1
// 	for left < right {
// 		r[left], r[right] = r[right], r[left]
// 		left++
// 		right--
// 	}
// 	return string(r)
// }
