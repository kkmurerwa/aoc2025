package day2

import (
	"aoc2025/helpers"
	"fmt"
	"strconv"
)

func Task1() {
	fmt.Println()
	fmt.Println("Running Task 1...")

	input := helpers.ReadFileHelper("/day2/day_2_input.txt")
	arr := helpers.CharacterSeparatedStringToList(input, ",")
	result := 0

	for _, v := range arr {
		rangeArr := helpers.CharacterSeparatedStringToList(v, "-")

		start, _ := strconv.Atoi(rangeArr[0])
		end, _ := strconv.Atoi(rangeArr[1])

		for i := start; i <= end; i++ {
			if hasRepeatedDigitsPart1(strconv.Itoa(i)) {
				result += i
			}

		}
	}

	fmt.Println("Result:", result)
}

func hasRepeatedDigitsPart1(digit string) bool {
	digitLen := len(digit)

	if digitLen%2 != 0 {
		return false
	}

	midpoint := digitLen / 2

	firstPart := digit[:midpoint]
	lastPart := digit[midpoint:]

	return firstPart == lastPart
}
