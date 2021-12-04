package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type BoardRow struct {
	numbers map[int]bool
}

type BoardColumn struct {
	BoardRow
}

type Board struct {
	values [][]int
}

func main() {

	f, e := ioutil.ReadFile("./data")
	if e != nil {
		panic(e)
	}

	data := strings.SplitN(string(f), "\n", 2)

	numbers := []int{}

	for _, n := range strings.Split(data[0], ",") {
		atoi, e := strconv.Atoi(n)
		if e != nil {
			panic(e)
		}
		numbers = append(numbers, atoi)
	}

	boards := [][][]int{}

	for _, b := range strings.Split(data[1][1:], "\n\n") {
		board := make([][]int, 5)
		for i, r := range strings.Split(b, "\n") {
			row := make([]int, 5)

			for j, n := range strings.Split(strings.Replace(strings.TrimSpace(r), "  ", " ", -1), " ") {
				atoi, e := strconv.Atoi(n)
				if e != nil {
					panic(e)
				}
				row[j] = atoi
			}
			board[i] = row
		}
		boards = append(boards, board)
	}

	log.Println(fmt.Sprintf("Part one: %d", solvePartOne(numbers, boards)))
	log.Println(fmt.Sprintf("Part two: %d", solvePartTwo(numbers, boards)))

}

func solvePartOne(numbers []int, boards [][][]int) int {

	winningBoardIndex := -1
	lastCalled := 0

MainLoop:
	for _, n := range numbers {
		for boardIndex, board := range boards {
			for i, row := range board {
				for j, v := range row {
					if v == n {
						boards[boardIndex][i][j] = -1
						if getCount(board) == 5 {
							winningBoardIndex = boardIndex
							lastCalled = n
							break MainLoop
						}
					}
				}
			}
		}
	}

	sum := 0

	for _, row := range boards[winningBoardIndex] {
		for _, n := range row {
			if n != -1 {
				sum += n
			}
		}
	}

	return sum * lastCalled
}

func solvePartTwo(numbers []int, boards [][][]int) int {

	lastWinningBoardIndex := -1
	lastCalled := 0
	winners := map[int]bool{}

MainLoop:
	for _, n := range numbers {
		for boardIndex, board := range boards {
			if _, ok := winners[boardIndex]; !ok {
				for i, row := range board {
					for j, v := range row {
						if v == n {
							boards[boardIndex][i][j] = -1
							if getCount(board) == 5 {
								winners[boardIndex] = true
								if len(winners) == len(boards) {
									lastCalled = n
									lastWinningBoardIndex = boardIndex
									break MainLoop
								}
							}
						}
					}
				}
			}
		}
	}

	sum := 0

	for _, row := range boards[lastWinningBoardIndex] {
		for _, n := range row {
			if n != -1 {
				sum += n
			}
		}
	}

	return sum * lastCalled
}

func getCount(board [][]int) int {
	count := 0

	for i := 0; i < len(board); i++ {
		xc := 0
		yc := 0
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == -1 {
				yc++
			}
			if board[j][i] == -1 {
				xc++
			}
		}
		if yc > count {
			count = yc
		}
		if xc > count {
			count = xc
		}
	}

	return count
}
