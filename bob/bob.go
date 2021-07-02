// Bob is a package with a single function giving a response from an angsty teenager.
package bob

import (
	"regexp"
	"strings"
)

// Hey returns a string response from an angsty teenager, given string output.
func Hey(remark string) string {
	isAlpha := regexp.MustCompile(`[A-Za-z]+`).MatchString

	switch {
	case strings.ToUpper(remark) == remark && strings.HasSuffix(remark, "?") && isAlpha(remark):
		return "Calm down, I know what I'm doing!"
	case strings.ToUpper(remark) == remark && isAlpha(remark):
		return "Whoa, chill out!"
	case strings.HasSuffix(strings.TrimSpace(remark), "?"):
		return "Sure."
	case len(remark) == 0 || len(strings.TrimSpace(remark)) == 0:
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
