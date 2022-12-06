package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func allAreOne(a []int) bool {
	for _, x := range a {
		if x != 1 {
			fmt.Printf("D: %v\n", a)
			return false
		}
	}
	return true
}

func star1(s string, n int) int {
	for i := n; i < len(s); i++ {
		ss := s[i-n : i]
		x := make([]int, n)
		for j, c := range strings.Split(ss, "") {
			if strings.Count(ss, c) > 1 {
				break
			}
			x[j] = x[j] + 1
			if j == n-1 && allAreOne(x) {
				fmt.Printf(":D %v, %v, %v, %v\n", i-n+1, ss, c, x)
				return i
			}
		}
	}

	return (-99999)
}

// Look, consistency is important even if pointless.
func star2(s string, n int) int {
	return star1(s, n)
}

func main() {

	f, err := os.Open("./input.txt")
	check(err)

	sb := new(strings.Builder)

	_, err = io.Copy(sb, io.ReadCloser(f))
	check(err)

	firstAnswer := star1(sb.String(), 4)
	fmt.Printf("Result 1: %v\n", firstAnswer)

	secondAnswer := star2(sb.String(), 14)
	fmt.Printf("Result 2: %v\n", secondAnswer)
}
