package main

import (
	"fmt"
	"github.com/Marijus/advent-of-code/common"
	"strconv"
	"strings"
)

const (
	Red   = "\033[31m"
	Reset = "\033[0m"
)

type Board struct {
	grid [][]*Cell
	won  bool
}

type Cell struct {
	value  int
	filled bool
}

func newBoard(boardLines []string) (board Board) {
	for _, line := range boardLines {
		var gridLine []*Cell

		for _, item := range strings.Fields(line) {
			number, err := strconv.Atoi(item)
			if err != nil {
				panic(err)
			}

			gridLine = append(gridLine, &Cell{
				value:  number,
				filled: false,
			})
		}

		board.grid = append(board.grid, gridLine)
	}

	return board
}

func (b *Board) fillNumber(number int) {
	for _, row := range b.grid {
		for _, col := range row {
			if col.value == number {
				col.filled = true
			}
		}
	}
}

func (b *Board) isWinner() (result bool) {
	// check for row winners
	for _, row := range b.grid {
		rowWinner := true

		for _, col := range row {
			if !col.filled {
				rowWinner = false
				break
			}
		}

		if rowWinner {
			return true
		}
	}

	// cehck for col winners
	matrixLength := len(b.grid[0])

	for j := 0; j < matrixLength; j++ {
		colWinner := true

		for i := 0; i < matrixLength; i++ {
			if !b.grid[i][j].filled {
				colWinner = false
				break
			}
		}

		if colWinner {
			return true
		}
	}

	return false
}

func (b *Board) getUnfilledSum() (sum int) {
	sum = 0
	for _, row := range b.grid {
		for _, col := range row {
			if !col.filled {
				sum += col.value
			}
		}
	}

	return sum
}

func (b *Board) print() {
	for _, row := range b.grid {
		lineString := ""
		for _, col := range row {
			if col.filled {
				lineString += fmt.Sprintf("%s%d%s ", Red, col.value, Reset)
			} else {
				lineString += fmt.Sprintf("%d ", col.value)
			}
		}

		fmt.Println(lineString)
	}

	fmt.Println()
}

func firstPart() {
	luckyNumbers, boards := readBoards("day4_input_test.txt")

	for i := 0; i < len(luckyNumbers); i++ {
		luckyNumber := luckyNumbers[i]
		fmt.Println("lucky number: ", luckyNumber)

		for _, board := range boards {
			board.fillNumber(luckyNumber)
			board.print()
			if board.isWinner() {
				result := board.getUnfilledSum() * luckyNumber
				fmt.Println("WINNER")
				fmt.Println(result)
				return
			}
		}
	}
}

func secondPart() {
	luckyNumbers, boards := readBoards("day4_input.txt")

	for i := 0; i < len(luckyNumbers); i++ {
		luckyNumber := luckyNumbers[i]
		fmt.Println("lucky number: ", luckyNumber)

		for _, board := range boards {
			board.fillNumber(luckyNumber)
			board.print()
			if board.isWinner() {
				result := board.getUnfilledSum() * luckyNumber
				fmt.Println("WINNER")
				fmt.Println(result)

				board.won = true

				boardsInPlay := 0
				for _, b := range boards {
					if !b.won {
						boardsInPlay++
					}
				}

				if boardsInPlay == 0 {
					fmt.Println("Final result")
					return
				}
			}
		}
	}
}

func readBoards(fileName string) (luckyNumbers []int, boards []*Board) {
	inputString := common.GetInput(fileName)
	lines := strings.Split(inputString, "\n")

	//var luckyNumbers []int
	for _, numberString := range strings.Split(lines[0], ",") {
		number, err := strconv.Atoi(numberString)
		if err != nil {
			panic(err)
		}

		luckyNumbers = append(luckyNumbers, number)
	}

	lines = lines[2:]
	numberOfBoards := len(lines) / 5

	//var boards []Board

	for i := 0; i < numberOfBoards; i++ {
		boardLines := lines[0:5]
		board := newBoard(boardLines)
		boards = append(boards, &board)
		board.print()

		if len(lines) > 6 {
			lines = lines[6:]
		}
	}

	return
}
