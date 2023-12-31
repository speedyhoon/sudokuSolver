package main

import "log"

func main() {
	log.SetFlags(log.Lshortfile)

	puz, err := loadPuzzle(` 654     
 3   5 76
       2 
7  8   61
     62  
 1    4  
 7   4 53
      1  
8   9    `)

	if err != nil {
		log.Println(err)
		return
	}

	puz.Print()
	puz.Solve()
	puz.Print()
}

/*puz, err := loadPuzzle(` 5     1
3 2   9 7
 6 9 7 5
  86 17
    3
  58 94
 7 5 3 9
5 4   3 1
 1     2 `)*/

/*puz := puzzle{	// Easy.
	0: row(1, 2, 3, 0, 5, 6, 7, 8, 9),
	1: row(4, 5, 6, 7, 8, 9, 1, 2, 0),
	2: row(7, 8, 0, 1, 2, 3, 0, 5, 6),
	3: row(2, 3, 4, 5, 0, 7, 8, 9, 0),
	4: row(5, 6, 7, 8, 9, 0, 2, 3, 0),
	5: row(0, 9, 0, 2, 3, 4, 5, 6, 7),
	6: row(3, 0, 5, 0, 0, 8, 9, 1, 2),
	7: row(6, 7, 8, 9, 1, 2, 3, 0, 0),
	8: row(9, 0, 0, 0, 0, 0, 0, 0, 0),
}*/
/*	puz := puzzle{ //Medium
	0: row(0, 0, 0, 0, 5, 0, 6, 0, 0),
	1: row(0, 0, 6, 1, 0, 2, 9, 0, 0),
	2: row(1, 8, 0, 0, 0, 0, 0, 2, 0),
	3: row(0, 5, 0, 0, 2, 0, 0, 7, 0),
	4: row(2, 0, 0, 9, 0, 3, 0, 0, 6),
	5: row(0, 7, 0, 0, 8, 0, 0, 4, 0),
	6: row(0, 6, 0, 0, 0, 0, 0, 9, 1),
	7: row(0, 0, 9, 4, 0, 7, 5, 0, 0),
	8: row(0, 0, 1, 0, 6, 0, 0, 0, 0),
}*/

/*	puz := loadPuzzle(` 5     1
	3 2   9 7
	 6 9 7 5
	  86 17
	    3
	  58 94
	 7 5 3 9
	5 4   3 1
	 1     2 `)*/

/*	puz := loadPuzzle(` 4 8    5
	 9     3
	3 8  7  2
	 3    9
	    6   7
	1 59   2
	     95
	  4
	8 25   1 `)*/

/*	puz := loadPuzzle(`  1     8
	 2  9 5 6
	9  4
	  7  1
	       3
	 6  4 2 9
	 3     4
	    5  8
	  67  3 5`)*/

//	puz := loadPuzzle(`43 5 69
// 9      8
//     2
//  7    1
//    3
//6  4 9 8
//    2 5
//  63 5  7
//3   8    `)

/*
		puz := loadPuzzle(` 4 8    5
	 9     3
	3 8  7  2
	 3
	    6   7
	1 59   2
	     95
	  4
	8 25   1 `)*/ /*

		puz := loadPuzzle(`   5 9  6
	 1
	54 8   9
	89 4   6
	    2 3
	  7
	  67
	  1    5
	75  4 9 2`)*/

/*
___________________
[1 2 3 0 5 6 7 8 9]
[4 5 6 7 8 9 1 2 0]
[7 8 0 1 2 3 0 5 6]
[2 3 4 5 0 7 8 9 0]
[5 6 7 8 9 0 2 3 0]
[0 9 0 2 3 4 5 6 7]
[3 0 5 0 0 8 9 1 2]
[6 0 8 9 1 2 3 0 0]
[0 0 0 0 0 0 0 0 0]
___________________

RUN 1
___________________
[1 2 3 4 5 6 7 8 9]
[4 5 6 7 8 9 1 2 3]
[7 8 9 1 2 3 4 5 6]
[2 3 4 5 0 7 8 9 0]
[5 6 7 8 9 0 2 3 0]
[0 9 0 2 3 4 5 6 7]
[3 0 5 0 0 8 9 1 2]
[6 0 8 9 1 2 3 0 0]
[0 0 0 0 0 0 6 0 0]
___________________
RUN 2
___________________
[1 2 3 4 5 6 7 8 9]
[4 5 6 7 8 9 1 2 3]
[7 8 9 1 2 3 4 5 6]
[2 3 4 5 0 7 8 9 0]
[5 6 7 8 9 0 2 3 0]
[0 9 0 2 3 4 5 6 7]
[3 0 5 0 0 8 9 1 2]
[6 0 8 9 1 2 3 0 0]
[0 0 0 0 0 0 6 0 0]
___________________
*/
