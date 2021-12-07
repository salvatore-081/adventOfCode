package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	f, e := ioutil.ReadFile("./example")
	if e != nil {
		log.Fatalln(e)
	}

	input := strings.Split(string(f), ",")

	log.Println(input)

	log.Println(fmt.Sprintf("Part one: %d", solvePartOne()))
	log.Println(fmt.Sprintf("Part two: %d", solvePartTwo()))

}

func solvePartOne() int {
	return 0
}

func solvePartTwo() int {
	return 0
}
