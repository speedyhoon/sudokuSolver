package main

import "fmt"

func (g *Group) solve() {
	m := g.emptyCells()
	if len(m) == 1 {
		*g[m[0]] = g.missingNumbers()[0]
	}
}

func (g *Group) emptyCells() (indexes []int) {
	for i, b := range g {
		if *b == 0 {
			indexes = append(indexes, i)
		}
	}

	return
}

func (g *Group) hasNumber(number byte) bool {
	for _, b := range g {
		if *b == number {
			return true
		}
	}
	return false
}

func (g *Group) missingNumbers() (numbers []byte) {
	for i := mn; i <= mx; i++ {
		if !g.hasNumber(i) {
			numbers = append(numbers, i)
		}
	}
	return
}

func (g *Group) Print() {
	fmt.Printf("[%d %d %d %d %d %d %d %d %d]\n", *g[0], *g[1], *g[2], *g[3], *g[4], *g[5], *g[6], *g[7], *g[8])
}
