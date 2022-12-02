package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func star1(scanner *bufio.Scanner) int {

	scanner.Split(bufio.ScanLines)

	bestElf := 0

	currentElf := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			if currentElf > bestElf {
				bestElf = currentElf
			}
			currentElf = 0
		} else {
			pineCones, err := strconv.Atoi(scanner.Text())
			check(err)
			currentElf += pineCones
		}
	}

	return bestElf
}

func star2(byteData []byte, n int) int {

	var intData = []int{}                     // Initialise integer slice for later use.
	strData := string(byteData)               // Convert input bytes to string.
	arrData := strings.Split(strData, "\n\n") // Convert single string into array of strings, each representing an elf.

	for _, d := range arrData {
		sum := 0
		v := strings.Split(d, "\n")
		for _, y := range v {
			x, err := strconv.Atoi(y) // This assumes no newline at the end of the input.
			check(err)
			sum += x
		}
		intData = append(intData, sum)
	}

	// Sort by value (ascending).
	sort.Ints(intData)

	// Grab the last n values of intData - the n highest.
	intData = intData[len(intData)-n:]

	tpc := 0

	for _, pc := range intData {
		tpc += pc
	}

	return tpc
}

func main() {

	f, err := os.Open("./input.txt")
	check(err)

	r := bufio.NewReader(f)
	s := bufio.NewScanner(r)

	firstAnswer := star1(s)
	fmt.Printf("Result 1: %v\n", firstAnswer)

	rf, err := os.ReadFile("./input.txt")
	check(err)

	secondAnswer := star2(rf, 3)
	fmt.Printf("Result 2: %v\n", secondAnswer)

}
