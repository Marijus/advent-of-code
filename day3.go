package main

import (
	"fmt"
	"strconv"
	"strings"
)

func dayThree() {
	inputString := getInput("day3_input_test.txt")
	var items []string
	for _, item := range strings.Split(inputString, "\n") {
		items = append(items, item)
	}

	mostCommonResult := ""
	leastCommonResult := ""

	for i := 0; i < len(items[0]); i++ {
		zeros := 0
		ones := 0

		for _, item := range items {
			if item[i:i+1] == "0" {
				zeros += 1
			} else {
				ones += 1
			}
		}

		if zeros > ones {
			mostCommonResult += "0"
			leastCommonResult += "1"
		} else {
			mostCommonResult += "1"
			leastCommonResult += "0"
		}
	}

	fmt.Println(mostCommonResult)
	mostCommonResultDecimal, err := strconv.ParseInt(mostCommonResult, 2, 64)
	if err != nil {
		panic(err)
	}

	fmt.Println(leastCommonResult)
	leastCommonResultDecimal, err := strconv.ParseInt(leastCommonResult, 2, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println("result: ", mostCommonResultDecimal*leastCommonResultDecimal)
}

func dayThreePartTwo() {
	inputString := getInput("day3_input.txt")
	var items []string
	for _, item := range strings.Split(inputString, "\n") {
		items = append(items, item)
	}

	oxygen := getRating(items, false)
	oxygenDecimal, err := strconv.ParseInt(oxygen, 2, 64)
	if err != nil {
		panic(err)
	}

	co2 := getRating(items, true)
	co2Decimal, err := strconv.ParseInt(co2, 2, 64)
	if err != nil {
		panic(err)
	}

	fmt.Println("oxygen: ", oxygenDecimal)
	fmt.Println("co2: ", co2Decimal)
	fmt.Println("result: ", oxygenDecimal*co2Decimal)
}

func getRating(items []string, reverse bool) (result string) {
	firstItem := items[0]

	for i := 0; i < len(firstItem); i++ {
		var zeroItems []string
		var oneItems []string

		for _, item := range items {
			if item[i:i+1] == "0" {
				zeroItems = append(zeroItems, item)
			} else {
				oneItems = append(oneItems, item)
			}
		}

		if !reverse {
			if len(zeroItems) > len(oneItems) {
				items = zeroItems
			} else if len(zeroItems) < len(oneItems) {
				items = oneItems
			} else {
				items = oneItems
			}
		} else {
			if len(zeroItems) < len(oneItems) {
				items = zeroItems
			} else if len(zeroItems) > len(oneItems) {
				items = oneItems
			} else {
				items = zeroItems
			}
		}

		if len(items) == 1 {
			result = items[0]
			return
		}
	}

	panic("could not calculate result")
}
