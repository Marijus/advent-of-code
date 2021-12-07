package main

import (
	"fmt"
	"github.com/Marijus/advent-of-code/common"
	"strconv"
	"strings"
	"time"
)

func firstPart() {
	inputString := common.GetInput("day6_test_input.txt")
	fmt.Println(inputString)
	var fishList []int

	for _, item := range strings.Split(inputString, ",") {
		fish, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}

		fishList = append(fishList, fish)
	}

	for day := 0; day < 80; day++ {
		currentLength := len(fishList)
		for i := 0; i < currentLength; i++ {
			if fishList[i] > 0 {
				fishList[i]--
			} else if fishList[i] == 0 {
				fishList[i] = 6
				fishList = append(fishList, 8)
			}
		}
		//fmt.Println(fishList)
	}

	fmt.Println(len(fishList))
}


func printMap(fishMap map[int]int) {
	line := ""
	for key, val := range fishMap {
		line += fmt.Sprintf("%d: %d ", key, val)
	}
	fmt.Println(line)
}

func secondPart() {
	inputString := common.GetInput("day6_test_input.txt")

	fishMap := map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
		8: 0,
		//9: 0,
	}

	for i := 0; i < 9; i++ {
		fishMap[i] = 0
	}

	printMap(fishMap)
	for _, item := range strings.Split(inputString, ",") {
		fish, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}

		//fmt.Println(item)
		fmt.Println(item)
		fmt.Println(fish)
		fishMap[fish] += 1
	}
	fmt.Println("before start")
	printMap(fishMap)

	for day := 0; day < 80; day++ {
		fmt.Println("----")
		fmt.Println(day+1)

		fmt.Println("before")
		printMap(fishMap)
		for i := 1; i <= 8; i++ {
			fishMap[i-1] = fishMap[i]
		}

		//fishMap[8] = fishMap[7] + fishMap[0]
		//fishMap[6] = fishMap[6] + fishMap[0]
		//fishMap[0] = 0

		//fmt.Println(fishMap)
		printMap(fishMap)
		fmt.Println("after")

	}

	result := 0
	for i := 1; i <= 8; i++ {
		result += fishMap[i]
	}

	fmt.Println(fishMap)
	fmt.Println("result: ", result)
}

func main() {
	t1 := time.Now()
	secondPart()
	fmt.Println("time elapsed: ", time.Now().Sub(t1))
}
