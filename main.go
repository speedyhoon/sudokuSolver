package main

import "log"

func main() {
	log.SetFlags(log.Lshortfile)

	puz := Load(
		" 654     ",
		" 3   5 76",
		"       2 ",
		"7  8   61",
		"     62  ",
		" 1    4  ",
		" 7   4 53",
		"      1  ",
		"8   9    ",
	)
	puz.Print()
	puz.Solve()
	puz.Print()
}
