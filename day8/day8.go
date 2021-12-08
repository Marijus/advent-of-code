package main

import (
	"fmt"
	"github.com/Marijus/advent-of-code/common"
	"strings"
	"time"
)


func firstPart() {
	//lengthMap := map[int]int{
	//	1: 2,
	//
	//}

	count := 0

	inputString := common.GetInput("day8_input.txt")
	for _, line := range strings.Split(inputString, "\n") {
		lineSplitted := strings.Split(line, " | ")

		outputVals := lineSplitted[1]
		for _, outputVal := range strings.Split(outputVals, " ") {
			if len(outputVal) == 2 || len(outputVal) == 4 || len(outputVal) == 3 || len(outputVal) == 7{
				//fmt.Println(outputVal)
				count += 1
			}
		}
	}

	fmt.Println("result: ", count)
}

func main() {
	t1 := time.Now()
	firstPart()
	fmt.Println("time elapsed: ", time.Now().Sub(t1))
}
