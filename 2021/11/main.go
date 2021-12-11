package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	X int
	Y int
}

func main() {
	f, e := ioutil.ReadFile("./data")
	if e != nil {
		log.Fatalln(e)
	}

	input := strings.Split(string(f), "\n")

	cavern := make([][]int, len(input))

	for ix, x := range input {
		cavern[ix] = make([]int, len(x))
		for iy, y := range x {
			v, e := strconv.Atoi(string(y))
			if e != nil {
				log.Fatalln(e)
			}
			cavern[ix][iy] = v
		}
	}

	for _, x := range cavern {
		log.Println(x)
	}

	start := time.Now()
	log.Println(fmt.Sprintf("Part one(%dns): %d", time.Since(start).Nanoseconds(), solvePartOne(cavern)))

	// re-init cavern
	for ix, x := range input {
		for iy, y := range x {
			v, _ := strconv.Atoi(string(y))
			cavern[ix][iy] = v
		}
	}

	start = time.Now()
	log.Println(fmt.Sprintf("Part two(%dns): %d", time.Since(start).Nanoseconds(), solvePartTwo(cavern)))

}

func flash(cavern [][]int, p Point, flashed map[Point]bool, flashCount *int) {
	if cavern[p.X][p.Y] <= 9 {
		cavern[p.X][p.Y]++
		if cavern[p.X][p.Y] > 9 {
			*flashCount++
			flashed[Point{p.X, p.Y}] = true
			// top
			if p.Y > 0 {
				flash(cavern, Point{X: p.X, Y: p.Y - 1}, flashed, flashCount)

				// top-left
				if p.X > 0 {
					flash(cavern, Point{X: p.X - 1, Y: p.Y - 1}, flashed, flashCount)
				}

				// top-right
				if p.X < len(cavern)-1 {
					flash(cavern, Point{X: p.X + 1, Y: p.Y - 1}, flashed, flashCount)
				}
			}

			// middle-left
			if p.X > 0 {
				flash(cavern, Point{X: p.X - 1, Y: p.Y}, flashed, flashCount)
			}

			// middle-right
			if p.X < len(cavern)-1 {
				flash(cavern, Point{X: p.X + 1, Y: p.Y}, flashed, flashCount)
			}

			// bottom
			if p.Y < len(cavern[p.X])-1 {
				flash(cavern, Point{X: p.X, Y: p.Y + 1}, flashed, flashCount)

				// bottom-left
				if p.X > 0 {
					flash(cavern, Point{X: p.X - 1, Y: p.Y + 1}, flashed, flashCount)
				}

				if p.X < len(cavern)-1 {
					flash(cavern, Point{X: p.X + 1, Y: p.Y + 1}, flashed, flashCount)
				}
			}
		}
	}
}

func solvePartOne(cavern [][]int) int {
	steps := 100
	flashed := map[Point]bool{}
	flashCount := 0

	for i := 0; i < steps; i++ {
		for p := range flashed {
			cavern[p.X][p.Y] = 0
			delete(flashed, p)
		}

		for ix, x := range cavern {
			for iy := range x {
				flash(cavern, Point{X: ix, Y: iy}, flashed, &flashCount)
			}
		}

	}
	return flashCount
}

func solvePartTwo(cavern [][]int) int {
	step := 0
	flashed := map[Point]bool{}
	flashCount := 0

	for true {
		if len(flashed) == len(cavern)*len(cavern[0]) {
			return step
		}
		for p := range flashed {
			cavern[p.X][p.Y] = 0
			delete(flashed, p)
		}

		for ix, x := range cavern {
			for iy := range x {
				flash(cavern, Point{X: ix, Y: iy}, flashed, &flashCount)
			}
		}
		step++

	}
	return flashCount
}
