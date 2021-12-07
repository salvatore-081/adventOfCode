package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	f, e := ioutil.ReadFile("./data")
	if e != nil {
		log.Fatalln(e)
	}

	input := strings.Split(string(f), "\n")

	report := make([]int, len(input))

	for index, line := range input {
		entry, e := strconv.Atoi(line)
		if e != nil {
			log.Fatalln(e)
		}
		report[index] = entry
	}

	log.Println(fmt.Sprintf("Part one: %d", solvePartOne(report)))
	log.Println(fmt.Sprintf("Part two: %d", solvePartTwo(report)))

}

func solvePartOne(report []int) int {
	result := make(chan int, 1)

	g := make(chan struct{}, runtime.NumCPU())
	ks := make(chan struct{}, 1)

	for _, entry := range report {
		g <- struct{}{}
		go func(entry int) {
			for _, v := range report {
				select {
				case <-ks:
					ks <- struct{}{}
					<-g
					return
				default:
					if v+entry == 2020 && v != entry {
						result <- v * entry
						ks <- struct{}{}
						<-g
						return
					}
				}
			}
			<-g
		}(entry)
	}

	return <-result
}

func solvePartTwo(report []int) int {
	result := make(chan int, 1)

	g := make(chan struct{}, runtime.NumCPU())
	ks := make(chan struct{}, 1)

	for _, first := range report {
		for _, second := range report {
			g <- struct{}{}
			go func(first int, second int) {
				for _, third := range report {
					select {
					case <-ks:
						ks <- struct{}{}
						<-g
						return
					default:
						if third+second+first == 2020 && third != second && second != first {
							result <- third * second * first
							ks <- struct{}{}
							<-g
							return
						}
					}
				}
				<-g
			}(first, second)
		}
	}

	return <-result
}
