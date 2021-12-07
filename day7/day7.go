package main

import (
	"fmt"
	"github.com/Marijus/advent-of-code/common"
	"strconv"
	"strings"
	"time"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getMaxVal(items []int) (maxVal int) {
	for _, item := range items {
		if item > maxVal {
			maxVal = item
		}
	}

	return maxVal
}

func firstPart() {
	crabs := getCrabs("day7_input_test.txt")

	var minIterationFuel int

	for i := 1; i < getMaxVal(crabs); i++ {
		iterationFuel := 0

		for _, item := range crabs {
			iterationFuel += abs(item - i)
		}

		if minIterationFuel == 0 || iterationFuel < minIterationFuel {
			minIterationFuel = iterationFuel
		}
	}

	fmt.Println("result: ", minIterationFuel)

}

func secondPart() {
	crabs := getCrabs("day7_input_test.txt")
	maxVal := getMaxVal(crabs)
	steps := make(map[int]int)

	stepVal := 0
	for i := 1; i < maxVal; i++ {
		stepVal = stepVal + i
		steps[i] = stepVal
	}

	var minIterationFuel int

	for i := 1; i < maxVal; i++ {
		iterationFuel := 0

		for _, item := range crabs {
			iterationFuel += steps[abs(item-i)]
		}

		if minIterationFuel == 0 || iterationFuel < minIterationFuel {
			minIterationFuel = iterationFuel
		}
	}

	fmt.Println("result: ", minIterationFuel)

}

func getCrabs(fileName string) (crabs []int) {
	inputString := common.GetInput(fileName)

	for _, item := range strings.Split(inputString, ",") {
		crab, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}

		crabs = append(crabs, crab)
	}

	return crabs
}

func main() {
	t1 := time.Now()
	secondPart()
	fmt.Println("time elapsed: ", time.Now().Sub(t1))
}
