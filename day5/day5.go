package main

import (
	"fmt"
	"github.com/Marijus/advent-of-code/common"
	"strconv"
	"strings"
)

type CoordinatesStorage struct {
	items [][]int
}

func NewCoordinateStorage() *CoordinatesStorage {
	items := make([][]int, 1024)

	for i := 0; i < 1024; i++ {
		items[i] = make([]int, 1024)
	}

	return &CoordinatesStorage{items: items}
}

func (storage *CoordinatesStorage) addCoordinates(x1, y1, x2, y2 int, countDiagonally bool) {
	if x1 == x2 {
		miny, maxy := getMinMax(y1, y2)
		for j := miny; j < maxy+1; j++ {
			storage.addSingleCoordinate(x1, j)
		}
	} else if y1 == y2 {
		minx, maxx := getMinMax(x1, x2)
		for i := minx; i < maxx+1; i++ {
			storage.addSingleCoordinate(i, y1)
		}
	} else if countDiagonally {
		storage.addSingleCoordinate(x1, y1)
		storage.addSingleCoordinate(x2, y2)

		if x1 < x2 {
			add := false
			if y2 > y1 {
				add = true
			}

			j := y1
			for i := x1 + 1; i < x2; i++ {
				if add {
					j += 1
				} else {
					j -= 1
				}

				storage.addSingleCoordinate(i, j)
			}
		} else {
			add := false
			if y2 < y1 {
				add = true
			}

			j := y2

			for i := x2 + 1; i < x1; i++ {
				if add {
					j += 1
				} else {
					j -= 1
				}

				storage.addSingleCoordinate(i, j)
			}
		}
	}
}

func (storage *CoordinatesStorage) addSingleCoordinate(x, y int) {
	storage.items[x][y] += 1
}

func (storage *CoordinatesStorage) getOverlappingPoints() (count int) {
	count = 0

	for i := 0; i < len(storage.items); i++ {
		for j := 0; j < len(storage.items[i]); j++ {
			if storage.items[i][j] > 1 {
				count++
			}
		}
	}

	return
}

func dayFive(countDiagonally bool) {
	inputString := common.GetInput("day5_input.txt")

	storage := NewCoordinateStorage()

	for _, line := range strings.Split(inputString, "\n") {
		items := strings.Split(line, " -> ")

		first := strings.Split(items[0], ",")
		y1, x1 := first[0], first[1]

		second := strings.Split(items[1], ",")
		y2, x2 := second[0], second[1]

		storage.addCoordinates(convertToInt(x1), convertToInt(y1), convertToInt(x2), convertToInt(y2), countDiagonally)
	}

	result := storage.getOverlappingPoints()
	fmt.Println("result: ", result)
}

func getMinMax(x1, x2 int) (min, max int) {
	if x1 < x2 {
		return x1, x2
	}

	return x2, x1
}

func convertToInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return result
}
