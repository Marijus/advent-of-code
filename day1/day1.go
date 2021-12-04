package main

import (
	"fmt"
	"github.com/Marijus/advent-of-code/common"
	"strconv"
	"strings"
)

func dayOnePartTwo() {
	data := getInputData()

	var sumsList []int
	for i := 0; i < len(data) - 2; i++ {
		sum := 0
		for _, item := range data[i:i+3] {
			sum += item
		}
		sumsList = append(sumsList, sum)
	}

	count := getCount(sumsList)
	fmt.Println(count)
}

func dayOne() {
	data := getInputData()
	count := getCount(data)
	fmt.Println(count)
}

func getCount(data []int) (count int) {
	count = 0
	for i := 1; i < len(data); i++ {
		if data[i] > data[i-1] {
			count++
		}
	}

	return count
}

func getInputData() []int {
	inputString := common.GetInput("day1_input.txt")
	var data []int

	for _, item := range strings.Split(inputString, "\n") {
		val, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}

		data = append(data, val)
	}

	return data
}