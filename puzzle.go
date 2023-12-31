package main

import (
	"bytes"
	"fmt"
	"log"
)

const mx byte = 9
const mn byte = 1

// Group is a group of cells either as a row, column or square.
type Group [mx]*Cell

type Puzzle struct {
	Cells    [mx][mx]Cell
	ToSolve  []byte // ToSolve stores which numbers need to be solved.
	ValueQty [mx + 1]byte
}

type Cell struct {
	Value byte
	Pos   []byte // Possibilities.
}

func (p *Puzzle) UnsolvedCells() (u byte) {
	return p.ValueQty[0]
}

func Load(row1, row2, row3, row4, row5, row6, row7, row8, row9 string) (p Puzzle) {
	p.ToSolve = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i, row := range []string{row1, row2, row3, row4, row5, row6, row7, row8, row9} {
		if len(row) != int(mx) {
			log.Panicf("row %d length should be %d, not %d.", i, mx, len(row))
		}

		for j, cell := range row {
			// Ignore non numeric runes.
			if cell <= '0' || cell > '9' {
				p.ValueQty[0]++
				p.Cells[i][j].Pos = p.ToSolve
				continue
			}

			value := byte(cell - '0')
			p.ValueQty[value]++
			p.Cells[i][j].Value = value
		}
	}

	p.updateToSolve()
	p.check()

	return p
}

func (p *Puzzle) updateToSolve() {
	for i := 0; i < len(p.ToSolve); i++ {
		if p.ValueQty[p.ToSolve[i]] == 0 {
			Remove(p.ToSolve, i)
			i--
		}
	}
}

func (p *Puzzle) Solve() {
	var prev byte
	for x := 0; p.UnsolvedCells() > 0 && prev != p.UnsolvedCells() && x <= 99; x++ {
		prev = p.UnsolvedCells()

		for i := range p.Cells {
			for j := range p.Cells[i] {
				if p.Cells[i][j].Value != 0 {
					continue
				}

				p.Possibles(byte(i), byte(j))

				if x == 6 && i == 3 && j == 4 {
					for _, pos := range p.Cells[byte(i)][byte(j)].Pos {
						p.RowCanHaveNumber(pos, byte(i))
					}
				}
			}
		}

		for _, number := range p.ToSolve {
			p.solveOnly1s(number)
		}

		p.updateToSolve()
		log.Println("unsolved cell quantity:", p.UnsolvedCells(), "to complete:", p.ToSolve)
	}
}

func (p *Puzzle) Other(r, c byte) {
	for _, pos := range p.Cells[r][c].Pos {
		p.RowCanHaveNumber(pos, r)
	}
}

func (p *Puzzle) RowCanHaveNumber(val byte, r byte) {
	var canHave [mx]bool
	qty := 0
	var lastIndex byte
	for i, b := range p.Row(r) {
		c := byte(i)
		if b.Value != 0 {
			canHave[i] = false
			continue
		}
		if p.Column(c).hasNumber(val) {
			canHave[i] = false
			continue
		}
		if p.Square(whichSquare(r, c)).hasNumber(val) {
			canHave[i] = false
			continue
		}
		canHave[i] = true
		qty++
		lastIndex = c
	}
	if qty == 1 {
		p.solvedCell(&p.Cells[r][lastIndex], val)
	}
}

func (p *Puzzle) solveOnly1s(number byte) {
	var cantBe [mx][mx]bool
	var hasNumRow, hasNumSquare, hasNumCol [mx]bool
	for i := byte(0); i < mx; i++ {
		hasNumRow[i] = p.Row(i).hasNumber(number)
		hasNumSquare[i] = p.Square(i).hasNumber(number)
		hasNumCol[i] = p.Column(i).hasNumber(number)
	}

	for i := byte(0); i < mx; i++ {
		for j := byte(0); j < mx; j++ {
			cantBe[i][j] = hasNumRow[i] || hasNumSquare[whichSquare(i, j)] || hasNumCol[j] || p.Cells[i][j].Value != 0
		}
	}

	for i := byte(0); i < mx; i++ {
		ok, index := canFillIn(squareBool(cantBe, i))
		if ok {
			p.solvedCell(p.Square(i)[index], number)
		}
	}
}

func (p *Puzzle) solvedCell(cell *Cell, value byte) {
	cell.Value = value
	cell.Pos = nil
	p.ValueQty[0]--
	p.ValueQty[value]++
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

func (p *Puzzle) Print() {
	fmt.Printf("___________________\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n___________________\n", printRow(p.Cells[0]), printRow(p.Cells[1]), printRow(p.Cells[2]), printRow(p.Cells[3]), printRow(p.Cells[4]), printRow(p.Cells[5]), printRow(p.Cells[6]), printRow(p.Cells[7]), printRow(p.Cells[8]))
	fmt.Println("unsolved cell quantity:", p.UnsolvedCells())
}

func printRow(r [mx]Cell) (s string) {
	buf := bytes.NewBufferString("[")
	for i, cell := range r {
		if i != 0 {
			buf.WriteString(" ")
		}

		if cell.Value == 0 {
			buf.WriteString(".")
		} else {
			// Write cell's value as a number.
			buf.WriteByte(cell.Value + '0')
		}
	}

	buf.WriteString("]")
	return buf.String()
}

func (p *Puzzle) check() {
	for i := byte(0); i < mx; i++ {
		if !p.Row(i).check("row", i) || !p.Square(i).check("square", i) || !p.Column(i).check("column", i) {
			log.Panicln("puzzle contains errors")
		}
	}
}

func (g *Group) check(typ string, X byte) (ok bool /*, number byte, dup1, dup2 int*/) {
	x := make(map[byte]int)
	var index int
	for i, b := range g {
		if b.Value == 0 {
			continue
		}

		index, ok = x[b.Value]
		if ok {
			log.Printf("%s %d contains duplicate number %d in cells %d and %d", typ, X+1, *b, index+1, i+1)
			return false //, *b, i, dup1
		}
		x[b.Value] = i
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
