package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/RyanCarrier/dijkstra"
)

func main() {
	f, e := ioutil.ReadFile("./data")
	if e != nil {
		log.Fatalln(e)
	}

	input := strings.Split(string(f), "\n")

	start := time.Now()
	log.Println(fmt.Sprintf("Part one: %d (%dns)", solvePartOne(&input), time.Since(start).Nanoseconds()))

	start = time.Now()
	log.Println(fmt.Sprintf("Part two: %d (%dns)", solvePartTwo(&input), time.Since(start).Nanoseconds()))
}

func solvePartOne(input *[]string) int {
	matrix := make([][]int, len((*input)[0]))

	for x := range matrix {
		y := make([]int, len(*input))
		for i := 0; i < len(*input); i++ {
			v, e := strconv.Atoi(string((*input)[i][x]))
			if e != nil {
				log.Fatalln(e)
			}
			y[i] = v
		}
		matrix[x] = y
	}
	ids := map[[2]int]int{}

	g := dijkstra.NewGraph()
	for x := range matrix {
		for y := range matrix[x] {
			k := [2]int{x, y}
			ids[k] = len(ids)
			g.AddVertex(ids[k])
		}
	}

	for x, row := range matrix {
		for y := range row {
			if x+1 < len(matrix) {
				g.AddArc(ids[[2]int{x, y}], ids[[2]int{x + 1, y}], int64(matrix[x+1][y]))
				g.AddArc(ids[[2]int{x + 1, y}], ids[[2]int{x, y}], int64(matrix[x][y]))

			}
			if y+1 < len(row) {
				g.AddArc(ids[[2]int{x, y}], ids[[2]int{x, y + 1}], int64(matrix[x][y+1]))
				g.AddArc(ids[[2]int{x, y + 1}], ids[[2]int{x, y}], int64(matrix[x][y]))

			}
		}
	}

	d, e := g.Shortest(ids[[2]int{0, 0}], ids[[2]int{len(matrix) - 1, len(matrix) - 1}])
	if e != nil {
		log.Fatalln(e)
	}

	return int(d.Distance)
}

func solvePartTwo(input *[]string) int {

	matrix := make([][]int, len((*input))*5)

	for x := range *input {

		for m := 0; m < 5; m++ {
			matrix[x+len(*input)*m] = make([]int, len(matrix))
		}

		for y := range (*input)[x] {
			v, e := strconv.Atoi(string((*input)[x][y]))
			if e != nil {
				log.Fatalln(e)
			}

			for m1 := 0; m1 < 5; m1++ {
				for m2 := 0; m2 < 5; m2++ {
					value := v + m1 + m2
					if value > 9 {
						value -= 9
					}
					matrix[x+len(*input)*m1][y+len(*input)*m2] = value
				}
			}
		}
	}

	ids := map[[2]int]int{}

	g := dijkstra.NewGraph()
	for x := range matrix {
		for y := range matrix[x] {
			k := [2]int{x, y}
			ids[k] = len(ids)
			g.AddVertex(ids[k])
		}
	}

	for x, row := range matrix {
		for y := range row {
			if x+1 < len(matrix) {
				g.AddArc(ids[[2]int{x, y}], ids[[2]int{x + 1, y}], int64(matrix[x+1][y]))
				g.AddArc(ids[[2]int{x + 1, y}], ids[[2]int{x, y}], int64(matrix[x][y]))

			}
			if y+1 < len(row) {
				g.AddArc(ids[[2]int{x, y}], ids[[2]int{x, y + 1}], int64(matrix[x][y+1]))
				g.AddArc(ids[[2]int{x, y + 1}], ids[[2]int{x, y}], int64(matrix[x][y]))

			}
		}
	}

	d, e := g.Shortest(ids[[2]int{0, 0}], ids[[2]int{len(matrix) - 1, len(matrix) - 1}])
	if e != nil {
		log.Fatalln(e)
	}

	return int(d.Distance)
}
