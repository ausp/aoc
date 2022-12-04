package main

import (
	"encoding/csv"
	"fmt"
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

func star1(sa [][]string) int {
	count := 0
	for _, s := range sa {
		l := strings.Split(s[0], "-")
		r := strings.Split(s[1], "-")
		if (atoiWrap(l[0]) <= atoiWrap(r[0]) && atoiWrap(l[1]) >= atoiWrap(r[1])) || (atoiWrap(l[0]) >= atoiWrap(r[0]) && atoiWrap(l[1]) <= atoiWrap(r[1])) {
			// fmt.Printf("Yes: %v\n", s)
			count++
		} else {
			// fmt.Printf("No: %v\n", s)
		}
	}
	return count
}

func star2(sa [][]string) int {
	count := 0
	for _, s := range sa {
		l := strings.Split(s[0], "-")
		r := strings.Split(s[1], "-")
		if !(atoiWrap(l[1]) < atoiWrap(r[0])) && !(atoiWrap(l[0]) > atoiWrap(r[1])) {
			// fmt.Printf("Yes: %v\n", s)
			count++
		} else {
			// fmt.Printf("No: %v\n", s)
		}
	}
	return count
}

func main() {

	f, err := os.Open("./input.txt")
	check(err)

	s, err := csv.NewReader(f).ReadAll()
	check(err)

	firstAnswer := star1(s)
	fmt.Printf("Result 1: %v\n", firstAnswer)

	secondAnswer := star2(s)
	fmt.Printf("Result 2: %v\n", secondAnswer)

}
