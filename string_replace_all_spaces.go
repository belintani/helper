package helper

import (
	"strings"
)

// ReplaceAllSpaces replaces all kind of spaces of a string to an alternate text
func ReplaceAllSpaces(text, char string) string {

	text = strings.ReplaceAll(text, " ", char) //yes, it's a strange space
	text = strings.ReplaceAll(text, " ", char) //normal space

	return text
}
