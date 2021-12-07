package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	days := 256

	f, e := ioutil.ReadFile("./data")
	if e != nil {
		log.Fatalln(e)
	}

	input := strings.Split(string(f), ",")

	fishes := []int{}

	for _, v := range input {
		lifespan, e := strconv.Atoi(v)
		if e != nil {
			log.Fatalln(e)
		}
		fishes = append(fishes, lifespan)
	}

	log.Println(fmt.Sprintf("Part one: %d", solvePartOne(18, fishes))) // takes all the ram available and crash if days is too big
	log.Println(fmt.Sprintf("Part two: %.0f", solvePartTwo(days, fishes)))

}

func solvePartOne(days int, input []int) int {
	fishes := []int{}

	for _, n := range input {
		fishes = append(fishes, n)
	}

	for i := 0; i < days; i++ {
		newFishes := 0
		for index, fish := range fishes {
			if fish == 0 {
				fishes[index] = 6
				newFishes++
			} else {
				fishes[index]--
			}
		}
		for j := 0; j < newFishes; j++ {
			fishes = append(fishes, 8)
		}

	}
	return len(fishes)
}

func solvePartTwo(days int, fishes []int) float64 {
	state := map[int]float64{}

	for i := 0; i <= 8; i++ {
		state[i] = 0
	}

	for i := 0; i < len(fishes); i++ {
		state[fishes[i]]++
	}

	for i := 0; i < days; i++ {
		newFishes := state[0]
		state[0] = state[1]
		state[1] = state[2]
		state[2] = state[3]
		state[3] = state[4]
		state[4] = state[5]
		state[5] = state[6]
		state[6] = newFishes + state[7]
		state[7] = state[8]
		state[8] = newFishes
	}

	total := float64(0)
	for _, v := range state {
		total += v
	}

	return total
}
