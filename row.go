package main

import "log"

func (p *puzzle) Row(index byte) *group {
	if index >= mx {
		log.Panicf("incorrect square index %d, expected 0-%d.", index, mx-1)
		return nil
	}

	return &group{&p[index][0], &p[index][1], &p[index][2], &p[index][3], &p[index][4], &p[index][5], &p[index][6], &p[index][7], &p[index][8]}
}
