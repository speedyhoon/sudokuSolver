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
	Value    byte
	Pos      []byte // Possibilities.
	row, col byte
}

func (p *Puzzle) UnsolvedCells() (u int) {
	return int(p.ValueQty[0])
}

func Load(row1, row2, row3, row4, row5, row6, row7, row8, row9 string) (p Puzzle) {
	p.ToSolve = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i, row := range []string{row1, row2, row3, row4, row5, row6, row7, row8, row9} {
		if len(row) != int(mx) {
			log.Panicf("row %d length should be %d, not %d.", i, mx, len(row))
		}

		for j, cell := range row {
			p.Cells[i][j].row = byte(i)
			p.Cells[i][j].col = byte(j)

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

	p.check()
	p.updateToSolve()

	return p
}

func (p *Puzzle) updatePossibilities() {
	for r := byte(0); r < mx; r++ {
		for c := byte(0); c < mx; c++ {
			if p.Cells[r][c].Value != 0 {
				continue
			}

			p.Possibles(r, c)
		}
	}
	p.updateToSolve()
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
	p.updatePossibilities()

	var prev int
	for p.UnsolvedCells() > 0 && prev != p.UnsolvedCells() {
		prev = p.UnsolvedCells()
		for r := byte(0); r < mx; r++ {
			for c := byte(0); c < mx; c++ {
				if p.Cells[r][c].Value != 0 {
					continue
				}

				for i := 0; i < len(p.Cells[r][c].Pos); i++ {
					pos := p.Cells[r][c].Pos[i]
					if p.Cells[r][c].ColHasNum(p, pos) {
						p.Cells[r][c].Pos = Remove(p.Cells[r][c].Pos, i)
						i--
					}
					if len(p.Cells[r][c].Pos) == 1 {
						p.solvedCell(&p.Cells[r][c], p.Cells[r][c].Pos[0])
					}
				}
				for _, pos := range p.Cells[r][c].Pos {
					p.RowCanHaveNumber(pos, r)
					p.ColCanHaveNumber(pos, c)
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

func (p *Puzzle) RowCanHaveNumber(val byte, r byte) {
	var qty, lastIndex byte
	rowHasNum := p.Row(r).hasNumber(val)
	for i, b := range p.Row(r) {
		c := byte(i)
		if b.Value != 0 || rowHasNum || p.Column(c).hasNumber(val) || p.Square(whichSquare(r, c)).hasNumber(val) {
			continue
		}

		qty++
		lastIndex = c
	}
	if qty == 1 {
		p.solvedCell(&p.Cells[r][lastIndex], val)
	}
}

func (p *Puzzle) ColCanHaveNumber(val byte, c byte) {
	var qty, lastIndex byte
	colHasNum := p.Column(c).hasNumber(val)
	for i, b := range p.Column(c) {
		r := byte(i)
		if b.Value != 0 || p.Row(r).hasNumber(val) || colHasNum || p.Square(whichSquare(r, c)).hasNumber(val) {
			continue
		}

		qty++
		lastIndex = r
	}
	if qty == 1 {
		p.solvedCell(&p.Cells[lastIndex][c], val)
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

	for r := byte(0); r < mx; r++ {
		for c := byte(0); c < mx; c++ {
			cantBe[r][c] = hasNumRow[r] || hasNumSquare[whichSquare(r, c)] || hasNumCol[c] || p.Cells[r][c].Value != 0
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
	p.Print()

	// Update adjacent cell possibilities.
	p.Row(cell.row).RemovePos(value, p)
	p.Column(cell.col).RemovePos(value, p)
	p.Square(whichSquare(cell.row, cell.col)).RemovePos(value, p)
}

func (g *Group) RemovePos(value byte, p *Puzzle) {
	for _, cell := range g {
		if cell.Value != 0 {
			continue
		}

		cell.Pos = removeValue(cell.Pos, value)
		if len(cell.Pos) == 1 {
			p.solvedCell(cell, cell.Pos[0])
		}
	}
}

func removeValue(pos []byte, value byte) []byte {
	for i, p := range pos {
		if p == value {
			return Remove(pos, i)
		}
	}
	return pos
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
	fmt.Printf("____1___3___5___7___\n0%v\n1%v\n2%v\n3%v\n4%v\n5%v\n6%v\n7%v\n8%v\n____1___3___5___7___\n", printRow(p.Cells[0]), printRow(p.Cells[1]), printRow(p.Cells[2]), printRow(p.Cells[3]), printRow(p.Cells[4]), printRow(p.Cells[5]), printRow(p.Cells[6]), printRow(p.Cells[7]), printRow(p.Cells[8]))
	fmt.Println("unsolved cell quantity:", p.UnsolvedCells())
}

func (p *Puzzle) Printer() string {
	return fmt.Sprintf("___________________\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n___________________\nunsolved cell quantity: %d\n", printRow(p.Cells[0]), printRow(p.Cells[1]), printRow(p.Cells[2]), printRow(p.Cells[3]), printRow(p.Cells[4]), printRow(p.Cells[5]), printRow(p.Cells[6]), printRow(p.Cells[7]), printRow(p.Cells[8]), p.UnsolvedCells())
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
