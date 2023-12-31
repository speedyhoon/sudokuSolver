package main

import (
	"bytes"
	"fmt"
	"log"
)

const mx byte = 9
const mn byte = 1

type group [mx]*byte
type puzzle [mx][mx]byte

// toSolve stores which numbers need to be solved.
var toSolve []byte
var valueQty [mx + 1]byte

func (p *puzzle) UnsolvedCells() (u byte) {
	for i := range p {
		for j := range p[i] {
			if p[i][j] == 0 {
				u++
			}
		}
	}
	return
}

func loadPuzzle(row1, row2, row3, row4, row5, row6, row7, row8, row9 string) (p puzzle) {
	toSolve = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i, row := range []string{row1, row2, row3, row4, row5, row6, row7, row8, row9} {
		if len(row) != int(mx) {
			log.Panicf("row %d length should be %d, not %d.", i, mx, len(row))
		}

		for j, cell := range row {
			// Ignore non numeric runes.
			if cell < '0' || cell > '9' {
				valueQty[0]++
				continue
			}

			value := byte(cell - '0')
			valueQty[value]++
			p[i][j] = value
		}
	}

	updateToSolve()
	p.check()

	return p
}

func updateToSolve() {
	for i := 0; i < len(toSolve); i++ {
		if valueQty[toSolve[i]] == 0 {
			Remove(toSolve, i)
			i--
		}
	}
}

func (p *puzzle) Solve() {
	pos := [mx][mx][]byte{}

	unsolved := mx * mx
	var prev byte
	for x := 0; prev != unsolved && x <= 99; x++ {
		prev = unsolved

		for i := range p {
			for j := range p[i] {
				if p[i][j] != 0 {
					continue
				}
				pos[i][j] = Possibles(&p[i][j], p.Row(byte(i)), p.Column(byte(j)), p.Square(whichSquare(byte(i), byte(j))))
			}
		}

		for _, number := range toSolve {
			p.solveOny1s(number)
		}

		updateToSolve()
		log.Println("unsolved cell quantity:", p.UnsolvedCells())
	}
}

func (p *puzzle) solveOny1s(number byte) {
	var cantBe [mx][mx]bool
	var hasNumRow, hasNumSquare, hasNumCol [mx]bool
	for i := byte(0); i < mx; i++ {
		hasNumRow[i] = p.Row(i).hasNumber(number)
		hasNumSquare[i] = p.Square(i).hasNumber(number)
		hasNumCol[i] = p.Column(i).hasNumber(number)
	}

	for i := byte(0); i < mx; i++ {
		for j := byte(0); j < mx; j++ {
			cantBe[i][j] = hasNumRow[i] || hasNumSquare[whichSquare(i, j)] || hasNumCol[j] || p[i][j] != 0
		}
	}

	for i := byte(0); i < mx; i++ {
		ok, index := canFillIn(squareBool(cantBe, i))
		if ok {
			solvedCell(p.Square(i)[index], number)
		}
	}
}

func solvedCell(cell *byte, value byte) {
	*cell = value
	valueQty[0]--
	valueQty[value]++
}

func tern(cond bool) string {
	if cond {
		return "."
	}
	return "X"
}

func canFillIn(r [mx]bool) (_ bool, index byte) {
	qty := 0
	for i, val := range r {
		if !val {
			qty++
			if qty >= 2 {
				return false, 0
			}
			index = byte(i)
		}
	}
	return qty == 1, index
}

func (p *puzzle) Print() {
	fmt.Printf("___________________\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n___________________\n", printRow(p[0]), printRow(p[1]), printRow(p[2]), printRow(p[3]), printRow(p[4]), printRow(p[5]), printRow(p[6]), printRow(p[7]), printRow(p[8]))
	fmt.Println("unsolved cell quantity:", p.UnsolvedCells())
}

func printRow(r [mx]byte) (s string) {
	buf := bytes.NewBufferString("[")
	for i, cell := range r {
		if i != 0 {
			buf.WriteString(" ")
		}

		if cell == 0 {
			buf.WriteString(".")
		} else {
			// Write cell's value as a number.
			buf.WriteByte(cell + '0')
		}
	}

	buf.WriteString("]")
	return buf.String()
}

func (p *puzzle) check() {
	for i := byte(0); i < mx; i++ {
		if !p.Row(i).check("row", i) || !p.Square(i).check("square", i) || !p.Column(i).check("column", i) {
			log.Panicln("puzzle contains errors")
		}
	}
}

func (g *group) check(typ string, X byte) (ok bool /*, number byte, dup1, dup2 int*/) {
	x := make(map[byte]int)
	var index int
	for i, b := range g {
		if *b == 0 {
			continue
		}

		index, ok = x[*b]
		if ok {
			log.Printf("%s %d contains duplicate number %d in cells %d and %d", typ, X+1, *b, index+1, i+1)
			return false //, *b, i, dup1
		}
		x[*b] = i
	}
	return true //, 0 ,0 ,0
}

func Remove[T any](t []T, index int) []T {
	if index < 0 {
		return t
	}

	l := len(t)
	if l == 0 || index >= l {
		return t
	}

	switch index {
	case 0:
		return t[1:]
	case l - 1:
		return t[:index]
	default:
		return append(t[:index], t[index+1:]...)
	}
}
