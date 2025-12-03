package day1

import (
	"aoc2025/helpers"
	"fmt"
	"strconv"
)

func Task2() {
	fmt.Println()
	fmt.Println("Running Task 2...")

	input := helpers.ReadFileHelper("/day1/day_1_input.txt")
	arr := helpers.StringToList(input)

	dial := 50
	zeroCount := 0

	for _, v := range arr {
		prefix := v[0]
		value := v[1:]
		number, _ := strconv.Atoi(value)
		zeroTimes := number / 100
		number = number % 100

		if prefix == 'L' {
			if number > dial && dial != 0 {
				zeroCount++
			}

			dial = dial - number

			if dial < 0 {
				dial = 100 + dial
			}
		} else {
			dial = dial + number

			if dial > 99 {
				dial = dial - 100

				if dial != 0 {
					zeroCount++
				}
			}
		}

		if dial == 0 {
			zeroCount++
		}

		zeroCount += zeroTimes
	}

	fmt.Println("Zero count: ", zeroCount)
}
