package helper

import (
	"strings"
)

// EqualsIgnoreSpaces compares two strings ignoring all leading and trailing white space
func EqualsIgnoreSpaces(s1, s2 string) bool {
	if s1 == s2 {
		return true
	}

	s1 = strings.ReplaceAll(s1, " ", " ")
	s2 = strings.ReplaceAll(s2, " ", " ")

	return strings.TrimSpace(s1) == strings.TrimSpace(s2)
}
