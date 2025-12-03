package main

import (
	"aoc2025/day1"
	"aoc2025/day2"
	"fmt"
	"strings"
)

func main() {
	printDivider()
	fmt.Println("Running Day 1 tasks.")
	day1.Task1()
	day1.Task2()
	printDivider()

	printDivider()
	fmt.Println("Running Day 2 tasks.")
	day2.Task1()
	day2.Task2()
	printDivider()
}

func printDivider() {
	fmt.Println(strings.Repeat("=", 50))
}
