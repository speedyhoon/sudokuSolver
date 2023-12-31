package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
)

const mx byte = 9
const mn byte = 1

type group [mx]*byte
type puzzle [mx][mx]byte

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

func loadPuzzle(row1, row2, row3, row4, row5, row6, row7, row8, row9 string) (p puzzle, err error) {
	for i, row := range []string{row1, row2, row3, row4, row5, row6, row7, row8, row9} {
		if len(row) != int(mx) {
			log.Panicf("row %d length should be %d, not %d.", i, mx, len(row))
		}

		for j, cell := range row {
			// Ignore non numeric runes.
			if cell < '0' || cell > '9' {
				continue
			}

			p[i][j] = byte(cell - '0')
		}
	}

	return p, p.check()
}

func (p *puzzle) Solve() {
	pos := [mx][mx][]byte{}

	unsolved := mx * mx
	var prev byte
	for x := 0; prev != unsolved && x <= 99; x++ {
		prev = unsolved

		for i := range p {
			for j := range p[i] {
				if i == 1 && j == 1 {
					fmt.Println()
				}
				if p[i][j] != 0 {
					continue
				}
				pos[i][j] = Possibles(&p[i][j], p.Row(byte(i)), p.Column(byte(j)), p.Square(whichSquare(byte(i), byte(j))))
			}
		}

		for number := byte(1); number <= mx; number++ {
			p.solveOny1s(number)
		}
		unsolved = p.UnsolvedCells()
		log.Println("unsolved cell quantity:", unsolved)
	}
}

func (p *puzzle) solveOny1s(number byte) {
	var cantBe1 [mx][mx]bool
	var hasNumRow, hasNumSquare, hasNumCol [mx]bool
	for i := byte(0); i < mx; i++ {
		hasNumRow[i] = p.Row(i).hasNumber(number)
		hasNumSquare[i] = p.Square(i).hasNumber(number)
		hasNumCol[i] = p.Column(i).hasNumber(number)
	}

	for i := byte(0); i < mx; i++ {
		for j := byte(0); j < mx; j++ {
			cantBe1[i][j] = hasNumRow[i] || hasNumSquare[whichSquare(i, j)] || hasNumCol[j] || p[i][j] != 0
		}
	}

	for i := byte(0); i < mx; i++ {
		ok, index := canFillIn(squareBool(cantBe1, i))
		if ok {
			p.fillInSquare(number, i, index)
		}
	}
}

func (p *puzzle) fillInSquare(number byte, squareIndex byte, cellIndex byte) {
	*p.Square(squareIndex)[cellIndex] = number
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

func (p *puzzle) check() (err error) {
	for i := byte(0); i < mx; i++ {
		if !p.Row(i).check("row", i) || !p.Square(i).check("square", i) || !p.Column(i).check("column", i) {
			return errors.New("puzzle contains errors")
		}
	}

	return nil
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
