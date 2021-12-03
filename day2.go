package main

import (
	"fmt"
	"strconv"
	"strings"
)

func dayTwo() {
	inputString := getInput("day2_input.txt")
	instructions := strings.Split(inputString, "\n")

	horizontal := 0
	vertical := 0

	for _, instruction := range instructions {
		instructionValues := strings.Split(instruction, " ")
		direction := instructionValues[0]
		val, err := strconv.Atoi(instructionValues[1])
		if err != nil {
			panic(err)
		}

		switch direction {
		case "forward":
			horizontal += val
		case "down":
			vertical += val
		case "up":
			vertical -= val
		}
	}

	fmt.Println(horizontal, vertical)
	fmt.Println(horizontal * vertical)
}

func dayTwoPartTwo() {
	inputString := getInput("day2_input.txt")
	instructions := strings.Split(inputString, "\n")

	horizontal := 0
	vertical := 0
	aim := 0

	for _, instruction := range instructions {
		instructionValues := strings.Split(instruction, " ")
		direction := instructionValues[0]
		val, err := strconv.Atoi(instructionValues[1])
		if err != nil {
			panic(err)
		}

		switch direction {
		case "forward":
			horizontal += val
			vertical += aim * val
		case "down":
			aim += val
		case "up":
			aim -= val
		}
	}

	fmt.Println(horizontal, vertical)
	fmt.Println(horizontal * vertical)
}
