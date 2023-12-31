package main

import "log"

func (p *Puzzle) Row(index byte) *Group {
	if index >= mx {
		log.Panicf("incorrect square index %d, expected 0-%d.", index, mx-1)
		return nil
	}

	return &Group{&p.Cells[index][0], &p.Cells[index][1], &p.Cells[index][2], &p.Cells[index][3], &p.Cells[index][4], &p.Cells[index][5], &p.Cells[index][6], &p.Cells[index][7], &p.Cells[index][8]}
}
