package isogram

import "strings"

func IsIsogram(input string) bool {
	// probably would be a better idea to use regex, but for this exercise it does the trick
	input = strings.ReplaceAll(input, "-", "")
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ToLower(input)
	splitInput := strings.Split(input, "")
	var letterOccurences = map[string]bool{}

	for _, letter := range splitInput {
		if !letterOccurences[letter] {
			letterOccurences[letter] = true
		} else {
			return false
		}
	}
	return true
}
