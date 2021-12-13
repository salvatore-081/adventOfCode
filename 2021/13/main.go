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

	thermalImage := [][]string{}
	xMax := 0
	yMax := 0

	dots := []Point{}
	dotsEnd := false

	foldInstructions := []string{}

	for _, row := range input {
		if !dotsEnd {
			if row == "" {
				dotsEnd = true
				continue
			}

			xy := strings.Split(row, ",")
			x, e := strconv.Atoi(xy[0])
			if e != nil {
				log.Fatalln(e)
			}
			y, e := strconv.Atoi(xy[1])
			if e != nil {
				log.Fatalln(e)
			}

			if x > xMax {
				xMax = x
			}

			if y > yMax {
				yMax = y
			}

			dots = append(dots, Point{X: x, Y: y})
		} else {
			foldInstructions = append(foldInstructions, row)
		}
	}

	for i := 0; i <= xMax; i++ {
		y := []string{}
		for j := 0; j <= yMax; j++ {
			y = append(y, ".")
		}
		thermalImage = append(thermalImage, y)
	}

	for _, p := range dots {
		thermalImage[p.X][p.Y] = "#"
	}

	start := time.Now()
	log.Println(fmt.Sprintf("Part one(%dns): %d", time.Since(start).Nanoseconds(), solvePartOne(thermalImage, foldInstructions)))

	start = time.Now()
	solvePartTwo(thermalImage, foldInstructions)
	log.Println(fmt.Sprintf("Part two(%dns)", time.Since(start).Nanoseconds()))

}

func solvePartOne(thermalImage [][]string, foldInstructions []string) int {
	fi := strings.Split(strings.Split(foldInstructions[0], "fold along ")[1], "=")
	dir := fi[0]
	value, e := strconv.Atoi(fi[1])
	if e != nil {
		log.Fatalln(e)
	}

	if dir == "y" {
		c, _ := foldY(value, thermalImage)
		return c
	}

	c, _ := foldX(value, thermalImage)
	return c
}

func solvePartTwo(thermalImage [][]string, foldInstructions []string) {
	tm := thermalImage
	for _, fi := range foldInstructions {
		split := strings.Split(strings.Split(fi, "fold along ")[1], "=")
		dir := split[0]
		value, e := strconv.Atoi(split[1])
		if e != nil {
			log.Fatalln(e)
		}

		if dir == "y" {
			_, tm = foldY(value, tm)
		} else {
			_, tm = foldX(value, tm)
		}

	}

	printThermalImage(tm)
}

func printThermalImage(tm [][]string) {
	for j := range tm[0] {
		for i := range tm {
			fmt.Printf(tm[i][j] + " ")
		}
		fmt.Println()
	}
}

func foldY(value int, tm [][]string) (int, [][]string) {

	dots := 0
	folded := [][]string{}

	for _, x := range tm {
		folded = append(folded, x[:value])
	}
	for ix := range folded {
		for i := value + 1; i < len(tm[ix]); i++ {
			if folded[ix][(i-(len(folded[ix]))*2)*-1] != "#" && tm[ix][i] == "#" {
				folded[ix][(i-(len(folded[ix]))*2)*-1] = "#"
			}
		}
	}

	return dots, folded
}

func foldX(value int, tm [][]string) (int, [][]string) {
	dots := 0
	first := [][]string{}
	second := [][]string{}

	for _, v := range tm[value] {
		if v == "#" {
			return dots, tm
		}
	}

	for i, x := range tm {
		if i < value {
			first = append(first, x)
		}
		if i > value {
			second = append(second, x)
		}
	}

	for i, j := 0, len(second)-1; i < j; i, j = i+1, j-1 {
		second[i], second[j] = second[j], second[i]
	}

	for i, x := range first {
		for j := range x {
			if first[i][j] == "#" {
				dots++
				continue
			} else if second[i][j] == "#" {
				first[i][j] = "#"
				dots++
			}
		}
	}

	return dots, first
}
