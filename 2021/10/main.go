package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
	"sync"
	"time"
)

type Stack []rune

func (s *Stack) push(r rune) {
	*s = append(*s, r)
}

func (s *Stack) pop(r rune) (e error) {
	firstOut := (*s)[len(*s)-1]

	if firstOut != r {
		return errors.New("illegal character")
	}
	*s = (*s)[:len(*s)-1]
	return e
}

func (s *Stack) score() int {
	score := 0
	for i := len(*s) - 1; i > -1; i-- {
		switch (*s)[i] {
		case '(':
			score = (score * 5) + 1
			break
		case '[':
			score = (score * 5) + 2
			break
		case '{':
			score = (score * 5) + 3
		case '<':
			score = (score * 5) + 4
			break
		}
	}

	return score
}

func main() {
	f, e := ioutil.ReadFile("./data")
	if e != nil {
		log.Fatalln(e)
	}

	input := strings.Split(string(f), "\n")

	start := time.Now()
	log.Println(fmt.Sprintf("Part one(%dns): %d", time.Since(start).Nanoseconds(), solvePartOne(input)))

	start = time.Now()
	log.Println(fmt.Sprintf("Part two(%dns): %d", time.Since(start).Nanoseconds(), solvePartTwo(input)))

}

func solvePartOne(input []string) int {
	illegals := make(chan int, len(input))

	wg := sync.WaitGroup{}

	for _, row := range input {
		wg.Add(1)
		go func(row string) {
			defer wg.Done()

			var s Stack
			for _, r := range row {
				switch r {
				case '(':
					s.push(r)
					break
				case '[':
					s.push(r)
					break
				case '{':
					s.push(r)
					break
				case '<':
					s.push(r)
					break
				case ')':
					e := s.pop('(')
					if e != nil {
						illegals <- 3
						return
					}
					break
				case ']':
					e := s.pop('[')
					if e != nil {
						illegals <- 57
						return
					}
					break
				case '}':
					e := s.pop('{')
					if e != nil {
						illegals <- 1197
						return
					}
					break
				case '>':
					e := s.pop('<')
					if e != nil {
						illegals <- 25137
						return
					}
					break
				}
			}
			illegals <- 0
		}(row)
	}

	wg.Wait()

	close(illegals)

	total := 0

	for illegal := range illegals {
		total += illegal
	}

	return total
}

func solvePartTwo(input []string) int {
	scores := make(chan int, len(input))

	wg := sync.WaitGroup{}

	for _, row := range input {
		wg.Add(1)
		go func(row string) {
			defer wg.Done()

			var s Stack
			for _, r := range row {
				switch r {
				case '(':
					s.push(r)
					break
				case '[':
					s.push(r)
					break
				case '{':
					s.push(r)
					break
				case '<':
					s.push(r)
					break
				case ')':
					e := s.pop('(')
					if e != nil {
						return
					}
					break
				case ']':
					e := s.pop('[')
					if e != nil {
						return
					}
					break
				case '}':
					e := s.pop('{')
					if e != nil {
						return
					}
					break
				case '>':
					e := s.pop('<')
					if e != nil {
						return
					}
					break
				}
			}
			if len(s) > 0 {
				scores <- s.score()
			}
		}(row)
	}

	wg.Wait()

	close(scores)

	s := []int{}

	for score := range scores {
		s = append(s, score)
	}

	sort.Ints(s)

	return s[len(s)/2]
}
