package main

import (
	"fmt"
	"github.com/Marijus/advent-of-code/common"
	"strconv"
	"strings"
	"time"
)

func printMap(heightMap [][]int) {
	for i := 0; i < len(heightMap); i++ {
		line := ""
		for j := 0; j < len(heightMap[i]); j++ {
			line += fmt.Sprintf("%d ", heightMap[i][j])
		}
		fmt.Println(line)
	}
}

func lowestValue(heightMap [][]int, i, j int) (lowestValue bool) {
	lowestValue = true
	valueToCheck := heightMap[i][j]

	//fmt.Println("val to check: ", valueToCheck)

	// right direction
	if j + 1 <= len(heightMap[i])-1 {
		if heightMap[i][j+1] < valueToCheck {
			//fmt.Println("+")
			return false
		}
	}

	// left direction
	if j - 1 >= 0 {
		if heightMap[i][j-1] < valueToCheck {
			//fmt.Println("++")
			return false
		}
	}

	// up direction
	if i - 1 >= 0 {
		if heightMap[i-1][j] < valueToCheck {
			//fmt.Println("+++")
			return false
		}
	}

	// down direction
	if i + 1 <= len(heightMap) - 1 {
		if heightMap[i+1][j] < valueToCheck {
			//fmt.Println("+++")
			return false
		}
	}

	return lowestValue
}

func firstPart() {
	var heightMap [][]int

	inputString := common.GetInput("day9_input.txt")
	for _, line := range strings.Split(inputString, "\n") {
		var ints []int

		for i := 0; i < len(line); i++ {
			newInt, err := strconv.Atoi(string(line[i]))
			if err != nil {
				panic(err)
			}

			ints = append(ints, newInt)
		}

		heightMap = append(heightMap, ints)
	}

	printMap(heightMap)

	result := 0
	for i := 0; i < len(heightMap); i++ {
		for j := 0; j < len(heightMap[i]); j++ {
			val := heightMap[i][j]

			if lowestValue(heightMap, i, j) {
				//fmt.Println(val)
				result += 1 + val
			}
		}

		//fmt.Println("---")
	}

	fmt.Println("result: ", result)
}

func main() {
	t1 := time.Now()
	firstPart()
	fmt.Println("time elapsed: ", time.Now().Sub(t1))
}
