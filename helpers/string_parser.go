package helpers

import "strings"

func StringToList(input string) []string {
	words := strings.Fields(input)

	return words
}
