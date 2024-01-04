package main

import (
	"log"
)

func (p *Puzzle) Column(index byte) *Group {
	if index >= mx {
		log.Panicf("incorrect square index %d, expected 0-%d.", index, mx-1)
		return nil
	}

	return &Group{&p.Cells[0][index], &p.Cells[1][index], &p.Cells[2][index], &p.Cells[3][index], &p.Cells[4][index], &p.Cells[5][index], &p.Cells[6][index], &p.Cells[7][index], &p.Cells[8][index]}
}

func (c *Cell) ColHasNum(p *Puzzle, number byte) bool {
	top, bottom := adjacentSquaresVertical(whichSquare(c.row, c.col))
	return p.Column(c.col).hasNumber(number) ||
		p.Square(top).squareColHasNum(c.col, number) ||
		p.Square(bottom).squareColHasNum(c.col, number)
}

func (p *Puzzle) ColHasNum(row, col, number byte) bool {
	top, bottom := adjacentSquaresVertical(whichSquare(row, col))
	return p.Column(col).hasNumber(number) ||
		p.Square(top).squareColHasNum(1, number) ||
		p.Square(bottom).squareColHasNum(1, number)
}
