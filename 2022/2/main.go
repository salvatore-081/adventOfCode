package main

import (
	"io/ioutil"
	"log"
	"strings"
)

type StrategyGuide struct {
	ShapeScore map[string]int
	Win        map[string]string
	Equal      map[string]string
}

type SecretStrategyGuide struct {
	Win        map[string]string
	Lose       map[string]string
	ShapeScore map[string]int
}

func main() {
	f, e := ioutil.ReadFile("./data")
	if e != nil {
		log.Fatalln(e)
	}

	input := strings.Split(string(f), "\n")

	// Part 1
	var sg StrategyGuide
	sg.ShapeScore = map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	sg.Win = map[string]string{
		"X": "C",
		"Y": "A",
		"Z": "B",
	}

	sg.Equal = map[string]string{
		"X": "A",
		"Y": "B",
		"Z": "C",
	}

	total := 0

	for _, round := range input {
		total += int(sg.ShapeScore[string(round[2])])
		if sg.Equal[string(round[2])] == string(round[0]) {
			total += 3
		} else if sg.Win[string(round[2])] == string(round[0]) {
			total += 6
		}
	}

	log.Printf("Part 1: %d", total)

	// Part 2
	total = 0
	var ssg SecretStrategyGuide

	ssg.ShapeScore = map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	ssg.Lose = map[string]string{
		"A": "C",
		"B": "A",
		"C": "B",
	}

	ssg.Win = map[string]string{
		"A": "B",
		"B": "C",
		"C": "A",
	}

	for _, round := range input {
		switch string(round[2]) {
		case "X":
			total += ssg.ShapeScore[ssg.Lose[string(round[0])]]
			continue
		case "Y":
			total += (3 + ssg.ShapeScore[string(round[0])])
			continue
		case "Z":
			total += (6 + ssg.ShapeScore[ssg.Win[string(round[0])]])
		}
	}

	log.Printf("Part 2: %d", total)

}
