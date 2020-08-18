package helper

import (
	"strings"
)

func ReplaceAllSpaces(text, char string) string {

	text = strings.ReplaceAll(text, " ", char) //yes, it's a strange space
	text = strings.ReplaceAll(text, " ", char) //normal space

	return text
}
