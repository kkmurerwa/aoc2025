package helpers

import "strings"

func StringToList(input string) []string {
	words := strings.Fields(input)

	return words
}

func CharacterSeparatedStringToList(input string, character string) []string {
	arr := strings.Split(input, character)

	return arr
}
