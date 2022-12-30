package main

import (
	"io/ioutil"
	"log"
	"strings"
	"time"
	"unicode"
)

func main() {
	f, e := ioutil.ReadFile("./data")
	if e != nil {
		log.Fatalln(e)
	}

	input := strings.Split(string(f), "\n")

	// part 1
	items := []rune{}

	for _, rucksack := range input {
		check := map[rune]bool{}
		for i, item := range rucksack {
			if i < (len(rucksack) / 2) {
				check[item] = true
			} else {
				if k := check[item]; k {
					delete(check, item)
					items = append(items, item)
				}
			}
		}
	}

	priority := 0

	for _, item := range items {
		if unicode.IsLower(item) {
			priority += int(item) - 96
		} else {
			priority += int(item) - 38
		}
	}

	log.Printf("Part 1: %d", priority)

	// part 2
	priority = 0
	start := time.Now()
OUTER:
	for i := 0; i < len(input); i += 3 {
		for _, k := range input[i] {
			for _, l := range input[i+1] {
				if k == l {
					for _, o := range input[i+2] {
						if l == o {
							if unicode.IsLower(l) {
								priority += int(l) - 96
							} else {
								priority += int(l) - 38
							}
							continue OUTER
						}
					}
				}
			}
		}
	}

	log.Printf("Part 2 (%dns): %d", time.Since(start).Nanoseconds(), priority)
}
