package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	f, e := ioutil.ReadFile("./data")
	if e != nil {
		log.Fatalln(e)
	}

	input := strings.Split(string(f), "\n")
	area := make([][]int, len(input))

	for i := range area {
		area[i] = make([]int, len(input[i]))
	}

	for i, row := range input {
		for j, v := range row {
			n, e := strconv.Atoi(string(v))
			if e != nil {
				log.Fatalln(e)
			}
			area[i][j] = n
		}
	}

	start := time.Now()
	log.Println(fmt.Sprintf("Part one(%dns): %d", time.Since(start).Nanoseconds(), solvePartOne(area)))

	start = time.Now()
	log.Println(fmt.Sprintf("Part two(%dns): %d", time.Since(start).Nanoseconds(), solvePartTwo(area)))

}

func solvePartOne(area [][]int) int {
	risk := make(chan int, 1)
	risk <- 0

	wg := sync.WaitGroup{}

	for ir, row := range area {
		wg.Add(1)
		go func(ir int, row []int) {
			defer wg.Done()
			for iv, v := range row {
				check := 0
				if ir > 0 && ir < len(area)-1 {
					// center
					if iv > 0 && iv < len(row)-1 {
						// top
						if v >= area[ir-1][iv] {
							check = 1
						}
						// left
						if v >= area[ir][iv-1] {
							check = 2
						}
						// right
						if v >= area[ir][iv+1] {
							check = 3
						}
						// bottom
						if v >= area[ir+1][iv] {
							check = 4
						}
					} else if iv == 0 { // center left edge
						// top
						if v >= area[ir-1][iv] {
							check = 5
						}
						// right
						if v >= area[ir][iv+1] {
							check = 6
						}
						// bottom
						if v >= area[ir+1][iv] {
							check = 7
						}
					} else if iv == len(row)-1 { // center right edge
						// top
						if v >= area[ir-1][iv] {
							check = 8
						}
						// left
						if v >= area[ir][iv-1] {
							check = 9
						}
						// bottom
						if v >= area[ir+1][iv] {
							check = 10
						}
					}
				} else if ir == 0 { // top edge
					if iv > 0 && iv < len(row)-1 { // center top edge
						// left
						if v >= area[ir][iv-1] {
							check = 11
						}
						// right
						if v >= area[ir][iv+1] {
							check = 12
						}
						// bottom
						if v >= area[ir+1][iv] {
							check = 13
						}
					} else if iv == 0 { // top left
						// right
						if v >= area[ir][iv+1] {
							check = 14
						}
						// bottom
						if v >= area[ir+1][iv] {
							check = 15
						}
					} else if iv == len(row)-1 { // top right
						// left
						if v >= area[ir][iv-1] {
							check = 16
						}
						// bottom
						if v >= area[ir+1][iv] {
							check = 17
						}
					}
				} else if ir == len(area)-1 { // bottom edge
					if iv > 0 && iv < len(row)-1 { // bottom center edge
						// top
						if v >= area[ir-1][iv] {
							check = 18
						}
						// left
						if v >= area[ir][iv-1] {
							check = 19
						}
						// right
						if v >= area[ir][iv+1] {
							check = 20
						}
					} else if iv == 0 { // bottom left
						// top
						if v >= area[ir-1][iv] {
							check = 21
						}
						// right
						if v >= area[ir][iv+1] {
							check = 22
						}
					} else if iv == len(row)-1 { // bottom right
						// top
						if v >= area[ir-1][iv] {
							check = 23
						}
						// left
						if v >= area[ir][iv-1] {
							check = 24
						}
					}
				}
				if check == 0 {
					currentRisk := <-risk
					currentRisk += v + 1
					risk <- currentRisk
				}
			}
		}(ir, row)
	}

	wg.Wait()

	return <-risk
}

func solvePartTwo(area [][]int) int {
	basins := make(chan [3]int, 1)
	basins <- [3]int{0, 0, 0}

	wg := sync.WaitGroup{}

	for ir, row := range area {
		wg.Add(1)
		go func(ir int, row []int) {
			defer wg.Done()
			for iv, v := range row {
				check := 0
				if ir > 0 && ir < len(area)-1 {
					// center
					if iv > 0 && iv < len(row)-1 {
						// top
						if v >= area[ir-1][iv] {
							check = 1
						}
						// left
						if v >= area[ir][iv-1] {
							check = 2
						}
						// right
						if v >= area[ir][iv+1] {
							check = 3
						}
						// bottom
						if v >= area[ir+1][iv] {
							check = 4
						}
					} else if iv == 0 { // center left edge
						// top
						if v >= area[ir-1][iv] {
							check = 5
						}
						// right
						if v >= area[ir][iv+1] {
							check = 6
						}
						// bottom
						if v >= area[ir+1][iv] {
							check = 7
						}
					} else if iv == len(row)-1 { // center right edge
						// top
						if v >= area[ir-1][iv] {
							check = 8
						}
						// left
						if v >= area[ir][iv-1] {
							check = 9
						}
						// bottom
						if v >= area[ir+1][iv] {
							check = 10
						}
					}
				} else if ir == 0 { // top edge
					if iv > 0 && iv < len(row)-1 { // center top edge
						// left
						if v >= area[ir][iv-1] {
							check = 11
						}
						// right
						if v >= area[ir][iv+1] {
							check = 12
						}
						// bottom
						if v >= area[ir+1][iv] {
							check = 13
						}
					} else if iv == 0 { // top left
						// right
						if v >= area[ir][iv+1] {
							check = 14
						}
						// bottom
						if v >= area[ir+1][iv] {
							check = 15
						}
					} else if iv == len(row)-1 { // top right
						// left
						if v >= area[ir][iv-1] {
							check = 16
						}
						// bottom
						if v >= area[ir+1][iv] {
							check = 17
						}
					}
				} else if ir == len(area)-1 { // bottom edge
					if iv > 0 && iv < len(row)-1 { // bottom center edge
						// top
						if v >= area[ir-1][iv] {
							check = 18
						}
						// left
						if v >= area[ir][iv-1] {
							check = 19
						}
						// right
						if v >= area[ir][iv+1] {
							check = 20
						}
					} else if iv == 0 { // bottom left
						// top
						if v >= area[ir-1][iv] {
							check = 21
						}
						// right
						if v >= area[ir][iv+1] {
							check = 22
						}
					} else if iv == len(row)-1 { // bottom right
						// top
						if v >= area[ir-1][iv] {
							check = 23
						}
						// left
						if v >= area[ir][iv-1] {
							check = 24
						}
					}
				}
				if check == 0 {
					state := map[[2]int]bool{}
					basin := caulcuateBasin(1, area, v, ir, iv, state)

					b := <-basins

					for i, v := range b {
						if v < basin {
							b[i] = basin
							sort.Ints(b[:])
							break
						}
					}

					basins <- b
				}
			}
		}(ir, row)
	}

	wg.Wait()

	b := <-basins

	return b[0] * b[1] * b[2]
}

func caulcuateBasin(count int, area [][]int, v int, x int, y int, state map[[2]int]bool) int {
	state[[2]int{x, y}] = true

	// top
	if _, ok := state[[2]int{x - 1, y}]; !ok && x-1 >= 0 && area[x-1][y] > v && area[x-1][y] != 9 {
		count++
		count = caulcuateBasin(count, area, area[x-1][y], x-1, y, state)
	}
	// right
	if _, ok := state[[2]int{x, y + 1}]; !ok && y+1 < len(area[0]) && area[x][y+1] > v && area[x][y+1] != 9 {
		count++
		count = caulcuateBasin(count, area, area[x][y+1], x, y+1, state)
	}

	// bottom
	if _, ok := state[[2]int{x + 1, y}]; !ok && x+1 < len(area) && area[x+1][y] > v && area[x+1][y] != 9 {
		count++
		count = caulcuateBasin(count, area, area[x+1][y], x+1, y, state)
	}

	// left
	if _, ok := state[[2]int{x, y - 1}]; !ok && y-1 >= 0 && area[x][y-1] > v && area[x][y-1] != 9 {
		count++
		count = caulcuateBasin(count, area, area[x][y-1], x, y-1, state)
	}

	return count
}
