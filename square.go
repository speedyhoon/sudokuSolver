package main

import (
	"log"
)

const (
	topLeft byte = iota
	topCent
	topRight
	midLeft
	center
	midRight
	bottomLeft
	bottomCent
	bottomRight
)

func whichSquare(row, col byte) byte {
	return [mx][mx]byte{
		{0, 0, 0, 1, 1, 1, 2, 2, 2},
		{0, 0, 0, 1, 1, 1, 2, 2, 2},
		{0, 0, 0, 1, 1, 1, 2, 2, 2},
		{3, 3, 3, 4, 4, 4, 5, 5, 5},
		{3, 3, 3, 4, 4, 4, 5, 5, 5},
		{3, 3, 3, 4, 4, 4, 5, 5, 5},
		{6, 6, 6, 7, 7, 7, 8, 8, 8},
		{6, 6, 6, 7, 7, 7, 8, 8, 8},
		{6, 6, 6, 7, 7, 7, 8, 8, 8},
	}[row][col]
}

func (p *Puzzle) Square(index byte) *Group {
	switch index {
	case topLeft:
		return &Group{
			&p.Cells[0][0], &p.Cells[0][1], &p.Cells[0][2],
			&p.Cells[1][0], &p.Cells[1][1], &p.Cells[1][2],
			&p.Cells[2][0], &p.Cells[2][1], &p.Cells[2][2],
		}
	case topCent:
		return &Group{
			&p.Cells[0][3], &p.Cells[0][4], &p.Cells[0][5],
			&p.Cells[1][3], &p.Cells[1][4], &p.Cells[1][5],
			&p.Cells[2][3], &p.Cells[2][4], &p.Cells[2][5],
		}
	case topRight:
		return &Group{
			&p.Cells[0][6], &p.Cells[0][7], &p.Cells[0][8],
			&p.Cells[1][6], &p.Cells[1][7], &p.Cells[1][8],
			&p.Cells[2][6], &p.Cells[2][7], &p.Cells[2][8],
		}
	case midLeft:
		return &Group{
			&p.Cells[3][0], &p.Cells[3][1], &p.Cells[3][2],
			&p.Cells[4][0], &p.Cells[4][1], &p.Cells[4][2],
			&p.Cells[5][0], &p.Cells[5][1], &p.Cells[5][2],
		}
	case center:
		return &Group{
			&p.Cells[3][3], &p.Cells[3][4], &p.Cells[3][5],
			&p.Cells[4][3], &p.Cells[4][4], &p.Cells[4][5],
			&p.Cells[5][3], &p.Cells[5][4], &p.Cells[5][5],
		}
	case midRight:
		return &Group{
			&p.Cells[3][6], &p.Cells[3][7], &p.Cells[3][8],
			&p.Cells[4][6], &p.Cells[4][7], &p.Cells[4][8],
			&p.Cells[5][6], &p.Cells[5][7], &p.Cells[5][8],
		}
	case bottomLeft:
		return &Group{
			&p.Cells[6][0], &p.Cells[6][1], &p.Cells[6][2],
			&p.Cells[7][0], &p.Cells[7][1], &p.Cells[7][2],
			&p.Cells[8][0], &p.Cells[8][1], &p.Cells[8][2],
		}
	case bottomCent:
		return &Group{
			&p.Cells[6][3], &p.Cells[6][4], &p.Cells[6][5],
			&p.Cells[7][3], &p.Cells[7][4], &p.Cells[7][5],
			&p.Cells[8][3], &p.Cells[8][4], &p.Cells[8][5],
		}
	case bottomRight:
		return &Group{
			&p.Cells[6][6], &p.Cells[6][7], &p.Cells[6][8],
			&p.Cells[7][6], &p.Cells[7][7], &p.Cells[7][8],
			&p.Cells[8][6], &p.Cells[8][7], &p.Cells[8][8],
		}
	}

	log.Panicf("incorrect square index %d, expected 0-%d.", index, mx-1)
	return nil
}

func (g *Group) has(val byte) bool {
	for _, b := range g {
		if b.Value == val {
			return true
		}
	}
	return false
}

func contains(row, column, square *Group, val byte) bool {
	return row.has(val) || column.has(val) || square.has(val)
}

func (p *Puzzle) Possibles(r, c byte) {
	row := p.Row(r)
	column := p.Column(c)
	square := p.Square(whichSquare(r, c))
	var b []byte

	for _, list := range p.Cells[r][c].Pos {
		if !contains(row, column, square, list) {
			b = append(b, list)
		}
	}

	if len(b) == 1 {
		p.solvedCell(&p.Cells[r][c], b[0])
	} else {
		p.Cells[r][c].Pos = b
	}
}

func squareBool(p [mx][mx]bool, index byte) [mx]bool {
	switch index {
	case topLeft:
		return [mx]bool{
			p[0][0], p[0][1], p[0][2],
			p[1][0], p[1][1], p[1][2],
			p[2][0], p[2][1], p[2][2],
		}
	case topCent:
		return [mx]bool{
			p[0][3], p[0][4], p[0][5],
			p[1][3], p[1][4], p[1][5],
			p[2][3], p[2][4], p[2][5],
		}
	case topRight:
		return [mx]bool{
			p[0][6], p[0][7], p[0][8],
			p[1][6], p[1][7], p[1][8],
			p[2][6], p[2][7], p[2][8],
		}
	case midLeft:
		return [mx]bool{
			p[3][0], p[3][1], p[3][2],
			p[4][0], p[4][1], p[4][2],
			p[5][0], p[5][1], p[5][2],
		}
	case center:
		return [mx]bool{
			p[3][3], p[3][4], p[3][5],
			p[4][3], p[4][4], p[4][5],
			p[5][3], p[5][4], p[5][5],
		}
	case midRight:
		return [mx]bool{
			p[3][6], p[3][7], p[3][8],
			p[4][6], p[4][7], p[4][8],
			p[5][6], p[5][7], p[5][8],
		}
	case bottomLeft:
		return [mx]bool{
			p[6][0], p[6][1], p[6][2],
			p[7][0], p[7][1], p[7][2],
			p[8][0], p[8][1], p[8][2],
		}
	case bottomCent:
		return [mx]bool{
			p[6][3], p[6][4], p[6][5],
			p[7][3], p[7][4], p[7][5],
			p[8][3], p[8][4], p[8][5],
		}
	case bottomRight:
		return [mx]bool{
			p[6][6], p[6][7], p[6][8],
			p[7][6], p[7][7], p[7][8],
			p[8][6], p[8][7], p[8][8],
		}
	}

	log.Panicf("incorrect square index %d, expected 0-%d.", index, mx-1)
	return [mx]bool{}
}

func (g *Group) squareColsHasNum(number byte) (column0, column1, column2 bool) {
	column0 = g[0].hasPos(number) || g[3].hasPos(number) || g[6].hasPos(number)
	column1 = g[1].hasPos(number) || g[4].hasPos(number) || g[7].hasPos(number)
	column2 = g[2].hasPos(number) || g[5].hasPos(number) || g[8].hasPos(number)

	return column0 && !column1 && !column2,
		!column0 && column1 && !column2,
		!column0 && !column1 && column2
}

func (g *Group) squareColHasNum(col, number byte) bool {
	column1 := g[0].hasPos(number) || g[3].hasPos(number) || g[6].hasPos(number)
	column2 := g[1].hasPos(number) || g[4].hasPos(number) || g[7].hasPos(number)
	column3 := g[2].hasPos(number) || g[5].hasPos(number) || g[8].hasPos(number)

	// Convert column index into square column index.
	switch col % 3 {
	case 0:
		return column1 && !column2 && !column3
	case 1:
		return !column1 && column2 && !column3
	case 2:
		return !column1 && !column2 && column3
	}

	log.Panicf("not a valid column %d", col)
	return false
}

func (c *Cell) hasPos(number byte) bool {
	for _, pos := range c.Pos {
		if pos == number {
			return true
		}
	}
	return false
}

func adjacentSquaresVertical(squareIndex byte) (top byte, bottom byte) {
	switch squareIndex {
	case 0:
		return 3, 6
	case 1:
		return 4, 7
	case 2:
		return 5, 8
	case 3:
		return 0, 6
	case 4:
		return 1, 7
	case 5:
		return 2, 8
	case 6:
		return 0, 3
	case 7:
		return 1, 7
	case 8:
		return 2, 5
	default:
		log.Panicln("out of bounds")
		return 0, 0
	}
}
