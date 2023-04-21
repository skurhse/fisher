package tictactoe

import "fmt"

const (
	A       string = "A"
	B       string = "B"
	Draw    string = "Draw"
	Pending string = "Pending"
)

var Players = [2]string{A, B}

func tictactoe(moves [][]int) string {
	var b [3][3]string

	for i, move := range moves {
		m, n := move[0], move[1]
		b[m][n] = Players[i%2]
	}

	pending := false
	ldiag, rdiag := b[0][0], b[0][2]

	if ldiag == "" || rdiag == "" {
		pending = true
	}

	for m := 0; m <= 2; m++ {

		switch b[m][m] {
		case A:
			if ldiag != A {
				ldiag = ""
			}
		case B:
			if ldiag != B {
				ldiag = ""
			}
		default:
			ldiag = ""
			pending = true
		}

		switch b[m][2-m] {
		case A:
			if rdiag != A {
				rdiag = ""
			}
		case B:
			if rdiag != B {
				rdiag = ""
			}
		default:
			rdiag = ""
			pending = true
		}

		row := b[m][0]
		col := b[0][m]

		if row == "" || col == "" {
			pending = true
		}

		for n := 1; n <= 2; n++ {
			switch b[m][n] {
			case A:
				if row != A {
					row = ""
				}
			case B:
				if row != B {
					row = ""
				}
			default:
				row = ""
				pending = true
			}
			switch b[n][m] {
			case A:
				if col != A {
					col = ""
				}
			case B:
				if col != B {
					col = ""
				}
			default:
				col = ""
				pending = true
			}
			if m == 1 {
				fmt.Println(row, n)
			}
		}
		if row != "" {
			return row
		}
		if col != "" {
			return col
		}
	}
	if ldiag != "" {
		return ldiag
	}
	if rdiag != "" {
		return rdiag
	}
	if pending {
		return Pending
	} else {
		return Draw
	}
}
