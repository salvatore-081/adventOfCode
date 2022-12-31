package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	f, e := ioutil.ReadFile("./data")
	if e != nil {
		log.Fatalln(e)
	}

	input := strings.Split(string(f), "\n")

	// part 1
	fullyContainedPairs := 0

	for _, pairs := range input {
		p := strings.Split(pairs, ",")
		pairsA := strings.Split(p[0], "-")
		aX, _ := strconv.Atoi(pairsA[0])
		aY, _ := strconv.Atoi(pairsA[1])
		pairsB := strings.Split(p[1], "-")
		bX, _ := strconv.Atoi(pairsB[0])
		bY, _ := strconv.Atoi(pairsB[1])
		if (aX <= bX && aY >= bY) || (bX <= aX && bY >= aY) {
			fullyContainedPairs++
		}
	}

	log.Printf("Part 1: %d", fullyContainedPairs)

	// part 2
	overlapPairs := 0

	for _, pairs := range input {
		p := strings.Split(pairs, ",")
		pairsA := strings.Split(p[0], "-")
		aX, _ := strconv.Atoi(pairsA[0])
		aY, _ := strconv.Atoi(pairsA[1])
		pairsB := strings.Split(p[1], "-")
		bX, _ := strconv.Atoi(pairsB[0])
		bY, _ := strconv.Atoi(pairsB[1])
		if aX <= bY && bX <= aY {
			overlapPairs++
		}
	}

	log.Printf("Part 2: %d", overlapPairs)

}
