package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

func main() {
	f, e := ioutil.ReadFile("./data")
	if e != nil {
		log.Fatalln(e)
	}

	input := strings.Split(string(f), ",")

	positions := make([]int, len(input))

	for i, v := range input {
		position, e := strconv.Atoi(v)
		if e != nil {
			log.Fatalln(e)
		}
		positions[i] = position
	}

	log.Println(fmt.Sprintf("Part one: %d", solvePartOne(positions)))
	log.Println(fmt.Sprintf("Part two: %d", solvePartTwo(positions)))

}

func solvePartOne(positions []int) int {
	fuel := make(chan int, 1)
	fuel <- int(^uint(0) >> 1)

	var wg sync.WaitGroup

	maxGoroutines := make(chan struct{}, runtime.NumCPU())

	for i := range positions {
		wg.Add(1)
		maxGoroutines <- struct{}{}
		go func(i int) {
			defer wg.Done()
			total := 0

			for _, position := range positions {
				total += int(math.Abs(float64(position - i)))
			}

			f := <-fuel

			if f > total {
				fuel <- total
			} else {
				fuel <- f
			}

			<-maxGoroutines
		}(i)
	}

	wg.Wait()

	return <-fuel
}

func solvePartTwo(positions []int) int {

	fuel := make(chan int, 1)
	fuel <- int(^uint(0) >> 1) // max representable int

	var wg sync.WaitGroup

	maxGoroutines := make(chan struct{}, runtime.NumCPU()) // max number of running goroutines

	for i := range positions {
		wg.Add(1)
		maxGoroutines <- struct{}{} // blocks if maxGoroutines is full
		go func(i int) {
			defer wg.Done()
			total := 0

			for _, position := range positions {
				d := int(math.Abs(float64(position - i)))
				total += (d) * (d + 1) / 2
			}

			f := <-fuel

			if f > total {
				fuel <- total
			} else {
				fuel <- f
			}

			<-maxGoroutines
		}(i)
	}

	wg.Wait()

	return <-fuel
}
