package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
	"time"
)

func main() {
	f, e := ioutil.ReadFile("./data")
	if e != nil {
		log.Fatalln(e)
	}

	input := strings.Split(string(f), "\n")

	start := time.Now()
	log.Println(fmt.Sprintf("Part one(%dns): %d", time.Since(start).Nanoseconds(), solvePartOne(input)))

	start = time.Now()
	log.Println(fmt.Sprintf("Part two(%dns): %0.f", time.Since(start).Nanoseconds(), solvePartTwo(input)))

}

type ToBeInserted struct {
	Index int
	Value rune
}

func solvePartOne(input []string) int {
	polymer := input[0]

	for i := 0; i < 10; i++ {
		toBeInserted := []ToBeInserted{}

		for _, v := range input[2:] {
			insertionRules := [2]rune{
				rune(v[0]),
				rune(v[1]),
			}
			r := rune(strings.Split(v, "> ")[1][0])

			for iElement, element := range polymer[:len(polymer)-1] {
				if insertionRules == [2]rune{element, rune(polymer[iElement+1])} {
					toBeInserted = append(toBeInserted, ToBeInserted{Index: iElement + 1, Value: r})
				}
			}
		}

		sort.SliceStable(toBeInserted, func(i, j int) bool {
			return toBeInserted[i].Index < toBeInserted[j].Index
		})

		shit := 0
		for _, v := range toBeInserted {
			polymer = insert(polymer, v.Index+shit, v.Value)
			shit++
		}
	}

	c := map[rune]int{}

	for _, v := range polymer {
		if _, ok := c[v]; ok {
			c[v]++
		} else {
			c[v] = 1
		}
	}

	min := int(^uint(0) >> 1)
	max := 0

	for _, v := range c {
		if v < min {
			min = v
		}

		if v > max {
			max = v
		}
	}

	return max - min
}

func solvePartTwo(input []string) float64 {
	polymer := map[[2]rune]int{}

	for i := 0; i < len(input[0])-1; i++ {
		r := [2]rune{rune(input[0][i]), rune(input[0][i+1])}
		if _, ok := polymer[r]; ok {
			polymer[r] = polymer[r] + 1
		} else {
			polymer[r] = 1
		}
	}

	for i := 0; i < 40; i++ {
		toBeInserted := map[[2]rune]int{}
		toBeDeleted := [][2]rune{}

		for _, instruction := range input[2:] {
			rules := [2]rune{
				rune(instruction[0]),
				rune(instruction[1]),
			}
			r := rune(strings.Split(instruction, "> ")[1][0])

			if val, ok := polymer[rules]; ok {
				if val2, ok2 := toBeInserted[[2]rune{rules[0], r}]; ok2 {
					toBeInserted[[2]rune{rules[0], r}] = val + val2
				} else {
					toBeInserted[[2]rune{rules[0], r}] = val
				}

				if val2, ok2 := toBeInserted[[2]rune{r, rules[1]}]; ok2 {
					toBeInserted[[2]rune{r, rules[1]}] = val + val2
				} else {
					toBeInserted[[2]rune{r, rules[1]}] = val
				}
				toBeDeleted = append(toBeDeleted, rules)
			}
		}

		for i := range toBeDeleted {
			delete(polymer, toBeDeleted[i])
		}

		for k, v := range toBeInserted {
			polymer[k] = v
		}
	}

	c := map[rune]int{}

	for k, v := range polymer {
		if _, ok := c[k[0]]; ok {
			c[k[0]] = c[k[0]] + v
		} else {
			c[k[0]] = v
		}

		if _, ok := c[k[1]]; ok {
			c[k[1]] = c[k[1]] + v
		} else {
			c[k[1]] = v
		}
	}

	min := int(^uint(0) >> 1)
	max := 0

	for _, v := range c {
		if v < min {
			min = v
		}

		if v > max {
			max = v
		}
	}

	return (float64(max) - float64(min)) / 2
}

func printPolymer(p map[[2]rune]int) {
	for k, v := range p {
		log.Println(string(k[0])+string(k[1])+":", v)
	}
}

// https://stackoverflow.com/a/61822301
// 0 <= index <= len(a)
func insert(s string, index int, value rune) string {
	a := make([]rune, len(s))
	for i, v := range s {
		a[i] = v
	}

	if len(a) == index { // nil or empty slice or after last element
		return string(append(a, value))
	}

	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value

	return string(a)
}
