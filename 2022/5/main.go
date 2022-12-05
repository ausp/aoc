package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func atoiWrap(s string) int {
	i, err := strconv.Atoi(s)
	check(err)
	return i
}

func readMoves(s string) [][]int {
	var moveArray [][]int
	for _, move := range strings.Split(s, "\n") {
		if move == "" {
			break
		}
		ma := strings.Split(move, " ")
		ta := []int{atoiWrap(ma[1]), atoiWrap(ma[3]), atoiWrap(ma[5])}
		moveArray = append(moveArray, ta)
	}
	return moveArray
}

func readCrates(s string) [][]string {
	var crateSlice [][]string
	for _, crateInput := range strings.Split(s, "\n") {
		if crateInput[0:1] == " " {
			break
		}
		var crateRow []string
		for k := 1; k < 10; k++ {
			if len(crateInput) < k*4-3 {
				crateRow = append(crateRow, "")
			} else if strings.Split(crateInput, "")[(k*4)-3] == " " {
				crateRow = append(crateRow, "")
			} else {
				crateRow = append(crateRow, strings.Split(crateInput, "")[(k*4)-3])
			}
		}
		crateSlice = append(crateSlice, crateRow)
		for i := range crateSlice {
			fmt.Printf("%v\n", crateSlice[i])
		}
	}
	return crateSlice
}

func newMatrix(d2, d1 int) [][]string {
	a := make([]string, d2*d1)
	m := make([][]string, d2)
	lo, hi := 0, d1
	for i := range m {
		m[i] = a[lo:hi:hi]
		lo, hi = hi, hi+d1
	}
	return m
}

// Stolen from Stack Overflow
func transpose(a [][]string) [][]string {
	b := newMatrix(len(a[0]), len(a))
	for i := 0; i < len(b); i++ {
		c := b[i]
		for j := 0; j < len(c); j++ {
			c[j] = a[j][i]
		}
	}
	return b
}

func orient(crateBoard [][]string) (newBoard [][]string) {
	for n, a := range crateBoard {
		var b []string
		for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}
		crateBoard[n] = a
		for _, c := range a {
			if c != "" {
				b = append(b, c)
			}
		}
		newBoard = append(newBoard, b)
	}
	return
}

func moveCrate(move []int, crates [][]string) [][]string {
	for i := 0; i < move[0]; i++ {
		crates[move[2]-1] = append(crates[move[2]-1], crates[move[1]-1][len(crates[move[1]-1])-1])
		crates[move[1]-1] = crates[move[1]-1][:len(crates[move[1]-1])-1]
	}
	return crates
}

func moveCrateStack(move []int, crates [][]string) [][]string {
	crates[move[2]-1] = append(crates[move[2]-1], crates[move[1]-1][len(crates[move[1]-1])-move[0]:]...)
	crates[move[1]-1] = crates[move[1]-1][:len(crates[move[1]-1])-move[0]]
	return crates
}

func readResult(crates [][]string) string {
	var sa []string
	for _, cr := range crates {
		sa = append(sa, cr[len(cr)-1])
	}
	return strings.Join(sa, "")
}

func moveCrates(moves [][]int, crates [][]string, stacking bool) string {
	if !stacking {
		for _, m := range moves {
			moveCrate(m, crates)
		}
	} else {
		for _, m := range moves {
			moveCrateStack(m, crates)
		}
	}

	return readResult(crates)
}

func star1(s string) string {

	d := strings.Split(s, "\n\n")

	crates := readCrates(d[0])
	crates = transpose(crates)
	crates = orient(crates)
	moves := readMoves(d[1])

	result := moveCrates(moves, crates, false)

	// fmt.Printf("%v\n", moves)
	// fmt.Printf("%v\n", crates)

	return result
}

func star2(s string) string {

	d := strings.Split(s, "\n\n")

	crates := readCrates(d[0])
	crates = transpose(crates)
	crates = orient(crates)
	moves := readMoves(d[1])

	result := moveCrates(moves, crates, true)

	// fmt.Printf("%v\n", moves)
	// fmt.Printf("%v\n", crates)

	return result
}

func main() {

	f, err := os.Open("./input.txt")
	check(err)

	sb := new(strings.Builder)

	_, err = io.Copy(sb, io.ReadCloser(f))
	check(err)

	firstAnswer := star1(sb.String())
	fmt.Printf("Result 1: %v\n", firstAnswer)

	secondAnswer := star2(sb.String())
	fmt.Printf("Result 2: %v\n", secondAnswer)
}
