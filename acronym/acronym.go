// Acronym is used to generated abbreviations of strings.
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate outputs abbreviation of a given string, taking capital letter of each word in the string.
func Abbreviate(s string) string {
	finalAcronym := make([]string, 4)
	isAlpha := regexp.MustCompile(`^[A-Za-z]+`).MatchString

	for _, word := range strings.FieldsFunc(s, SplitHelper) {
		firstLetter := string(word[0])

		for i := 1; len(word)-1 > i && !isAlpha(firstLetter); i++ {
			firstLetter = string(word[i])
		}

		finalAcronym = append(finalAcronym, strings.ToUpper(firstLetter))
	}
	return strings.Join(finalAcronym, "")
}

func SplitHelper(r rune) bool {
	return r == '-' || r == ' '
}
