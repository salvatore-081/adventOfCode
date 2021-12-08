package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

type Mapping struct {
	Top         rune
	TopLeft     rune
	TopRight    rune
	Middle      rune
	BottomLeft  rune
	BottomRight rune
	Bottom      rune
}

func (m *Mapping) init(patterns []string) {
	sort.Slice(patterns, func(i, j int) bool {
		return len(patterns[i]) < len(patterns[j])
	})

	var one, three, four, five, seven, eight, nine string

	one = patterns[0]
	seven = patterns[1]
	four = patterns[2]
	eight = patterns[9]

	m.Top = findDifferences(seven, one)[0]

	for i := 0; i < 3; i++ {
		d := findDifferences(patterns[6+i], four)
		if len(d) == 2 {
			nine = patterns[6+i]
			if d[0] == m.Top {
				m.Bottom = d[1]
			} else {
				m.Bottom = d[0]
			}

		}
	}

	threeAndFive := []string{}

	for i := 0; i < 3; i++ {
		d := findDifferences(nine, patterns[3+i])
		if len(d) == 1 {
			threeAndFive = append(threeAndFive, patterns[3+i])
		}
	}

	for _, v := range threeAndFive {
		d := findDifferences(one, v)
		if len(d) == 0 {
			three = v
		} else {
			five = v
		}
	}

	for _, v := range findDifferences(three, one) {
		if v != m.Top && v != m.Bottom {
			m.Middle = v
			break
		}
	}

	for _, v := range findDifferences(four, one) {
		if v != m.Middle {
			m.TopLeft = v
			break
		}
	}

	for _, v := range five {
		if v != m.Top && v != m.TopLeft && v != m.Middle && v != m.Bottom {
			m.BottomRight = v
			break
		}
	}

	for _, v := range one {
		if v != m.BottomRight {
			m.TopRight = v
			break
		}
	}

	for _, v := range eight {
		if v != m.Top && v != m.TopLeft && v != m.TopRight && v != m.Middle && v != m.BottomRight && v != m.Bottom {
			m.BottomLeft = v
			break
		}
	}

}

func (m *Mapping) print() {
	log.Println("", string(m.Top)+string(m.Top)+string(m.Top)+string(m.Top), "")
	log.Println(string(m.TopLeft), "  ", string(m.TopRight))
	log.Println(string(m.TopLeft), "  ", string(m.TopRight))
	log.Println("", string(m.Middle)+string(m.Middle)+string(m.Middle)+string(m.Middle), "")
	log.Println(string(m.BottomLeft), "  ", string(m.BottomRight))
	log.Println(string(m.BottomLeft), "  ", string(m.BottomRight))
	log.Println("", string(m.Bottom)+string(m.Bottom)+string(m.Bottom)+string(m.Bottom), "")

}

func (m *Mapping) decode(pattern []string) string {
	output := ""

	for _, digit := range pattern {
		if len(digit) == 2 {
			output += "1"
			continue
		}

		if len(digit) == 3 {
			output += "7"
			continue
		}

		if len(digit) == 4 {
			output += "4"
			continue
		}

		if len(digit) == 7 {
			output += "8"
			continue
		}

		if len(digit) == 5 {
			if len(findDifferences(string([]rune{m.Top, m.TopRight, m.Middle, m.BottomLeft, m.Bottom}), digit)) == 0 {
				output += "2"
				continue
			}

			if len(findDifferences(string([]rune{m.Top, m.TopRight, m.Middle, m.BottomRight, m.Bottom}), digit)) == 0 {
				output += "3"
				continue
			}

			output += "5"
			continue
		}

		if len(digit) == 6 {
			if len(findDifferences(string([]rune{m.Top, m.TopLeft, m.TopRight, m.BottomLeft, m.BottomRight, m.Bottom}), digit)) == 0 {
				output += "0"
				continue
			}

			if len(findDifferences(string([]rune{m.Top, m.TopLeft, m.Middle, m.BottomLeft, m.BottomRight, m.Bottom}), digit)) == 0 {
				output += "6"
				continue
			}

			output += "9"
		}
	}

	return output
}

func findDifferences(s1 string, s2 string) []rune {
	differences := []rune{}

	for _, v1 := range s1 {
		found := false
		for _, v2 := range s2 {
			if v1 == v2 {
				found = true
			}
		}
		if !found {
			differences = append(differences, v1)
		}
	}

	return differences
}

func main() {
	f, e := ioutil.ReadFile("./data")
	if e != nil {
		log.Fatalln(e)
	}

	input := strings.Split(string(f), "\n")

	log.Println(fmt.Sprintf("Part one: %d", solvePartOne(input)))
	log.Println(fmt.Sprintf("Part two: %d", solvePartTwo(input)))

}

func solvePartOne(input []string) int {
	count := 0
	for _, row := range input {
		fourDigitOutputValue := strings.Split(row, "| ")[1]
		for _, digit := range strings.Split(fourDigitOutputValue, " ") {
			segments := len(digit)
			if segments == 2 || segments == 4 || segments == 3 || segments == 7 {
				count++
			}
		}
	}
	return count
}

func solvePartTwo(input []string) int {
	total := 0

	for _, row := range input {
		s := strings.Split(row, " | ")
		patterns := s[0]
		outputs := s[1]

		mapping := Mapping{}
		mapping.init(strings.Split(patterns, " "))
		c, e := strconv.Atoi(mapping.decode(strings.Split(outputs, " ")))
		if e != nil {
			log.Fatalln(e)
		}
		total += c
	}

	return total
}
