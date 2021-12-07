package main

import (
	"fmt"
	"github.com/Marijus/advent-of-code/common"
	"strconv"
	"strings"
	"time"
)


func getSum(fishMap map[int]int) (sum int) {
	for _, val := range fishMap {
		sum += val
	}

	return sum
}

func calculateFishes(fileName string, numberOfDays int) {
	inputString := common.GetInput(fileName)

	fishMap := make(map[int]int)
	for i := 0; i < 9; i++ {
		fishMap[i] = 0
	}

	for _, item := range strings.Split(inputString, ",") {
		fish, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}

		fishMap[fish] += 1
	}

	for day := 0; day < numberOfDays; day++ {
		zeroFishes := fishMap[0]

		for i := 1; i <= 8; i++ {
			fishMap[i-1] = fishMap[i]
		}

		fishMap[8] = zeroFishes
		fishMap[6] = fishMap[6] + zeroFishes
	}

	result := getSum(fishMap)

	fmt.Println("result: ", result)
}

func main() {
	t1 := time.Now()
	calculateFishes("day6_input_test.txt", 256)
	fmt.Println("time elapsed: ", time.Now().Sub(t1))
}
