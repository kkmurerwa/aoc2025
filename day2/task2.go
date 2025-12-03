package day2

import (
	"aoc2025/helpers"
	"fmt"
	"strconv"
)

func Task2() {
	fmt.Println()
	fmt.Println("Running Task 2...")

	input := helpers.ReadFileHelper("/day2/day_2_input.txt")
	arr := helpers.CharacterSeparatedStringToList(input, ",")
	result := 0

	for _, v := range arr {
		rangeArr := helpers.CharacterSeparatedStringToList(v, "-")

		start, _ := strconv.Atoi(rangeArr[0])
		end, _ := strconv.Atoi(rangeArr[1])

		for i := start; i <= end; i++ {
			if hasRepeatedDigitsPart2(strconv.Itoa(i)) {
				result += i
			}
		}
	}

	fmt.Println("Result:", result)
}

func hasRepeatedDigitsPart2(digit string) bool {
	digitLen := len(digit)

	if digitLen == 1 {
		return false
	}

	division := 1

	for division <= digitLen {
		lastGroup := digit[:division]

		stillEqual := false

		for i := division; i < digitLen; i += division {
			if (i + division) > digitLen {
				stillEqual = false
				break
			}

			currentGroup := digit[i : i+division]

			if currentGroup != lastGroup {
				stillEqual = false
				break
			} else {
				stillEqual = true
			}

			lastGroup = currentGroup
		}

		if stillEqual {
			return true
		}

		division++
	}

	return false
}
