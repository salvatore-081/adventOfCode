package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"sync"
	"time"
	"unicode"
)

type Node struct {
	value string
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.value)
}

type Graph struct {
	nodes []*Node
	edges map[Node][]*Node
	mutex sync.RWMutex
}

func (g *Graph) init() {
	g.edges = make(map[Node][]*Node)
}

func (g *Graph) addNode(n *Node) {
	g.mutex.Lock()
	g.nodes = append(g.nodes, n)
	g.mutex.Unlock()
}

func (g *Graph) addEdge(n1, n2 *Node) {
	g.mutex.Lock()
	g.edges[*n1] = append(g.edges[*n1], n2)
	g.edges[*n2] = append(g.edges[*n2], n1)
	g.mutex.Unlock()
}

func (g *Graph) String() {
	g.mutex.RLock()
	s := ""
	for i := 0; i < len(g.nodes); i++ {
		s += g.nodes[i].String() + " -> "
		near := g.edges[*g.nodes[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}
	fmt.Println(s)
	g.mutex.RUnlock()
}

func main() {
	f, e := ioutil.ReadFile("./data")
	if e != nil {
		log.Fatalln(e)
	}

	input := strings.Split(string(f), "\n")

	caves := Graph{}
	caves.init()
	caves.String()

	for _, line := range input {
		s := strings.Split(line, "-")

		var c1, c2 bool
		for _, n := range caves.nodes {
			if n.value == s[0] {
				c1 = true
			}

			if n.value == s[1] {
				c2 = true
			}

			if c1 && c2 {
				break
			}
		}

		if !c1 {
			caves.addNode(&Node{value: s[0]})
		}

		if !c2 {
			caves.addNode(&Node{value: s[1]})
		}

		caves.addEdge(&Node{value: s[0]}, &Node{value: s[1]})
	}

	caves.String()

	start := time.Now()
	log.Println(fmt.Sprintf("Part one(%dns): %d", time.Since(start).Nanoseconds(), solvePartOne(&caves)))

	start = time.Now()
	log.Println(fmt.Sprintf("Part two(%dns): %d", time.Since(start).Nanoseconds(), solvePartTwo(&caves)))

}

func solvePartOne(caves *Graph) int {
	count := 0
	countMutex := sync.RWMutex{}

	state := map[Node]bool{}

	for _, n := range caves.edges[Node{value: "start"}] {
		state[Node{value: "start"}] = true
		traversePartOne(n, &count, &countMutex, caves, state)
	}
	return count
}

func traversePartOne(node *Node, count *int, countMutex *sync.RWMutex, caves *Graph, s map[Node]bool) {

	if node.value == "end" {
		countMutex.Lock()
		*count++
		countMutex.Unlock()
		return
	}

	state := map[Node]bool{}
	for k, v := range s {
		state[k] = v
	}

	if unicode.IsLower(rune(node.value[0])) {
		state[*node] = true
	}

	for _, n := range caves.edges[*node] {
		if _, ok := state[*n]; !ok {
			traversePartOne(n, count, countMutex, caves, state)
		}
	}

	return
}

func solvePartTwo(caves *Graph) int {
	count := 0
	countMutex := sync.RWMutex{}

	state := map[Node]bool{}

	lock := false

	for _, n := range caves.edges[Node{value: "start"}] {
		traversePartTwo(n, &count, &countMutex, caves, state, lock)
	}
	return count
}

func traversePartTwo(node *Node, count *int, countMutex *sync.RWMutex, caves *Graph, s map[Node]bool, lock bool) {

	if node.value == "end" {
		countMutex.Lock()
		*count++
		countMutex.Unlock()
		return
	}

	state := map[Node]bool{}
	for k, v := range s {
		state[k] = v
	}

	if node.value != "start" && unicode.IsLower(rune(node.value[0])) {
		if _, ok := state[*node]; ok {
			lock = true
		} else {
			state[*node] = true
		}
	}

	for _, n := range caves.edges[*node] {
		if n.value != "start" {
			if _, ok := state[*n]; !ok || !lock {
				traversePartTwo(n, count, countMutex, caves, state, lock)
			}
		}
	}

	return
}
