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

	input := strings.Split(string(f), "\n\n")

	// Part 1
	max := 0

	for _, elf := range input {
		total := 0
		elfC := strings.Split(string(elf), "\n")
		for _, v := range elfC {
			n, _ := strconv.Atoi(v)
			total += n
		}
		if max < total {
			max = total
		}
	}

	log.Printf("Part 1: %d", max)

	// Part 2
	topThree := []int{0, 0, 0}

	for _, elf := range input {
		total := 0
		elfC := strings.Split(string(elf), "\n")
		for _, v := range elfC {
			n, _ := strconv.Atoi(v)
			total += n
		}

		for i, top := range topThree {
			if top < total {
				topThree = insert(topThree, i, total)
				topThree = topThree[0:3]
				break
			}
		}
	}

	log.Printf("Part 2: %d", topThree[0]+topThree[1]+topThree[2])

}

func insert(a []int, index int, value int) []int {
	a = append(a[:index+1], a[index:]...)
	a[index] = value
	return a
}
