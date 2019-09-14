package main

func judgeCircle(moves string) bool {
	x, y := 0, 0
	for _, r := range moves {
		switch r {
		case 'U':
			y++
		case 'D':
			y--
		case 'L':
			x--
		case 'R':
			x++
		}
	}
	return x == 0 && y == 0
}
