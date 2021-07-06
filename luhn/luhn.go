package luhn

import (
	"strconv"
	"strings"
)

func Valid(input string) bool {
	var sum int
	input = strings.ReplaceAll(input, " ", "")
	splitString := strings.Split(input, "")
	digitsSlice := make([]int, len(splitString))

	if len(digitsSlice) == 0 || len(digitsSlice) == 1 {
		return false
	}

	for index, letter := range splitString {
		val, err := strconv.Atoi(letter)
		if err != nil {
			return false
		}
		digitsSlice[index] = val
	}

	for i := len(splitString) - 2; i >= 0; i -= 2 {
		newVal := digitsSlice[i] * 2
		if newVal > 9 {
			newVal -= 9
		}
		digitsSlice[i] = newVal
	}

	for _, value := range digitsSlice {
		sum += value
	}

	return sum%10 == 0
}
