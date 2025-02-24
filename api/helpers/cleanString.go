package helpers

import (
	"regexp"
	"strings"
)

var DefaultForbiddenCharacters = `[<>'";&|$\[\]{}#%=]`

func CleanSpace(input string) string {
	return strings.TrimSpace(input)
}

func SanitizeInput(input string) string {
	re := regexp.MustCompile(DefaultForbiddenCharacters)
	sanitized := re.ReplaceAllString(input, "")

	return sanitized
}
