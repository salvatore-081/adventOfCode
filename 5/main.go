package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	oceanFloor := [999][999]int{}

	f, e := ioutil.ReadFile("./data")
	if e != nil {
		log.Fatalln(e)
	}

	rows := strings.Split(string(f), "\n")

	data := make([][2][2]int, len(rows))

	for index, row := range rows {
		s := strings.Split(row, " -> ")

		x1y1 := strings.Split(s[0], ",")
		x2y2 := strings.Split(s[1], ",")

		x1, e := strconv.Atoi(x1y1[0])
		if e != nil {
			log.Fatalln(e)
		}

		y1, e := strconv.Atoi(x1y1[1])
		if e != nil {
			log.Fatalln(e)
		}

		x2, e := strconv.Atoi(x2y2[0])
		if e != nil {
			log.Fatalln(e)
		}

		y2, e := strconv.Atoi(x2y2[1])
		if e != nil {
			log.Fatalln(e)
		}

		coordinates := [2][2]int{{x1, y1}, {x2, y2}}

		data[index] = coordinates
	}

	log.Println(fmt.Sprintf("Part one: %d", solvePartOne(data, oceanFloor)))
	log.Println(fmt.Sprintf("Part two: %d", solvePartTwo(data, oceanFloor)))

}

func printOceanFloor(oceanFloor [999][999]int) {
	log.Println("")
	for _, row := range oceanFloor {
		fmt.Println(row)
	}
}

func solvePartOne(data [][2][2]int, oceanFloor [999][999]int) int {
	count := 0

	for _, entry := range data {
		if entry[0][0] == entry[1][0] {
			for i := entry[0][1]; i <= entry[1][1]; i++ {
				oceanFloor[i][entry[0][0]]++

			}
			for i := entry[1][1]; i <= entry[0][1]; i++ {
				oceanFloor[i][entry[0][0]]++
			}
		} else if entry[0][1] == entry[1][1] {
			for i := entry[0][0]; i <= entry[1][0]; i++ {
				oceanFloor[entry[0][1]][i]++
			}
			for i := entry[1][0]; i <= entry[0][0]; i++ {
				oceanFloor[entry[0][1]][i]++
			}
		}
	}

	for _, row := range oceanFloor {
		for _, v := range row {
			if v >= 2 {
				count++
			}
		}
	}

	return count
}

func solvePartTwo(data [][2][2]int, oceanFloor [999][999]int) int {
	count := 0

	for _, entry := range data {
		switch math.Atan2((float64(entry[1][1])-float64(entry[0][1])), (float64(entry[1][0])-float64(entry[0][0]))) * (180 / math.Pi) {
		case 0: // [[0,9],[5,9]]
			for i := 0; i <= (entry[1][0] - entry[0][0]); i++ {
				oceanFloor[entry[0][1]][entry[0][0]+i]++
			}
			break
		case 45: // [[0,0],[8,8]]
			for i := 0; i <= entry[1][0]-entry[0][0]; i++ {
				oceanFloor[entry[0][1]+i][entry[0][0]+i]++
			}
			break
		case 90: // [[7,0],[7,4]]
			for i := 0; i <= entry[1][1]-entry[0][1]; i++ {
				oceanFloor[entry[0][1]+i][entry[0][0]]++
			}
			break
		case 135: // [[8,0],[0,8]]
			for i := 0; i <= entry[1][1]-entry[0][1]; i++ {
				oceanFloor[entry[1][1]-i][entry[1][0]+i]++
			}
			break
		case 180: // [[9,4],[3,4]]
			for i := 0; i <= entry[0][0]-entry[1][0]; i++ {
				oceanFloor[entry[0][1]][entry[1][0]+i]++
			}
			break
		case -45: // [[5,5],[8,2]]
			for i := 0; i <= entry[1][0]-entry[0][0]; i++ {
				oceanFloor[entry[1][1]+i][entry[1][0]-i]++
			}
			break
		case -90:
			// [[2,2],[2,1]]
			// [[212,680],[212,136]]
			for i := 0; i <= entry[0][1]-entry[1][1]; i++ {
				oceanFloor[entry[0][1]-i][entry[0][0]]++
			}
			break
		case -135: // [[6,4],[2,0]]
			for i := 0; i <= entry[0][1]-entry[1][1]; i++ {
				oceanFloor[entry[1][1]+i][entry[1][0]+i]++
			}
			break
		default:
			log.Println("default", entry, math.Atan2((float64(entry[1][1])-float64(entry[0][1])), (float64(entry[1][0])-float64(entry[0][0])))*(180/math.Pi))
			break
		}
	}

	for _, row := range oceanFloor {
		for _, v := range row {
			if v >= 2 {
				count++
			}
		}
	}

	return count
}
