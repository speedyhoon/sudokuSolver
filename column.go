package main

import (
	"log"
)

func (p *puzzle) Column(index byte) *group {
	if index >= mx {
		log.Panicf("incorrect square index %d, expected 0-%d.", index, mx-1)
		return nil
	}

	return &group{&p[0][index], &p[1][index], &p[2][index], &p[3][index], &p[4][index], &p[5][index], &p[6][index], &p[7][index], &p[8][index]}
}
