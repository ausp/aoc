package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func star1(df dataframe.DataFrame) int {

	split := func(s series.Series) series.Series {
		left := s.Elem(0).String()[:len(s.Elem(0).String())/2]
		right := s.Elem(0).String()[len(s.Elem(0).String())/2:]

		s2 := series.Strings([]string{left, right})

		// fmt.Printf("%s,  %s, %s, %s\n", s, left, right, s2)

		return s2
	}

	match := func(s series.Series) series.Series {
		left := s.Elem(0).String()
		right := s.Elem(1).String()

		for _, c := range strings.Split(left, "") {
			if strings.Contains(right, c) {
				return series.Strings([]string{left, right, c})
			}

		}
		return series.Strings([]string{left, right, "sadface"})

	}

	convert := func(s series.Series) series.Series {
		left := s.Elem(0).String()
		right := s.Elem(1).String()
		match := s.Elem(2).String()
		point := int([]rune(s.Elem(2).String())[0]) - 96 // Normalises to a=1

		// fmt.Printf("%s,  %s, %s, %v, %v\n", left, right, match, point, int([]rune(s.Elem(2).String())[0]))

		if point > 0 { // Rune values of uppercase characters are lower than lowercase counterparts; 0 is a determiner after normalising.
			return series.Strings([]string{left, right, match, strconv.Itoa(point)}) // Lowercase
		} else {
			return series.Strings([]string{left, right, match, strconv.Itoa(point + 58)}) // Uppercase
		}

	}

	sum := func(s series.Series) series.Series {
		ints, err := s.Int()
		check(err)
		sum := 0
		for _, n := range ints {
			sum += n
		}
		return series.Ints(sum)
	}

	converted := df.Select([]string{"X0"}).Rapply(split).Rapply(match).Rapply(convert)
	converted.SetNames("left", "right", "match", "point")

	// fmt.Print(converted)

	total, err := converted.Select([]string{"point"}).Capply(sum).Elem(0, 0).Int()
	check(err)

	return total
}

func star2(s string) int {

	// Grab slice
	sa := strings.Split(s, "\n")

	// fmt.Printf("%v\n", sa[0])

	var ga [][]string
	var ia []int
	i := 0

	for i = 0; i <= len(sa); i = i + 3 {
		if sa[i] != "" {
			ga = append(ga, sa[i:i+2])

			for _, c := range strings.Split(sa[i], "") {
				if strings.Contains(sa[i+1], c) && strings.Contains(sa[i+2], c) {
					iv := int([]rune(c)[0]) - 96
					if iv > 0 {
						ia = append(ia, iv)
					} else {
						ia = append(ia, iv+58)
					}
					// fmt.Printf("%v, %v, %v, %v, %v\n", sa[i], sa[i+1], sa[i+2], iv, c)
					break
				}

			}
		}
	}
	// fmt.Printf("%v\n%v\n%v\n%v\n", ga, ia, len(ga), len(ia))

	total := 0
	for _, n := range ia {
		total += n
	}

	return total
}

func main() {

	f1, err := os.Open("./input.txt")
	check(err)

	r := bufio.NewReader(f1)
	df := dataframe.ReadCSV(r, dataframe.WithDelimiter(' '), dataframe.HasHeader(false))

	firstAnswer := star1(df)
	fmt.Printf("Result 1: %v\n", firstAnswer)

	f, err := os.ReadFile("./input.txt")
	check(err)

	secondAnswer := star2(string(f))
	fmt.Printf("Result 2: %v\n", secondAnswer)

}
