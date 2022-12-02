package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func star1(reader *bufio.Reader) int {

	df := dataframe.ReadCSV(reader, dataframe.WithDelimiter(' '), dataframe.HasHeader(false))

	fmt.Print(df)

	choicesToPointConversion := func(s series.Series) series.Series {
		player1Choice := s.Elem(0).String()
		player2Choice := s.Elem(1).String()
		var outcome int
		switch player2Choice {
		case "X":
			points := 1
			switch player1Choice {
			case "A":
				points = 3
			case "B":
				points = 0
			case "C":
				points = 6
			}
			return series.Ints(points + outcome)
		case "Y":
			points := 2
			switch player1Choice {
			case "A":
				points = 6
			case "B":
				points = 3
			case "C":
				points = 0
			}
			return series.Ints(points + outcome)
		case "Z":
			points := 3
			switch player1Choice {
			case "A":
				points = 0
			case "B":
				points = 6
			case "C":
				points = 3
			}
			return series.Ints(points + outcome)
		}
		return series.Ints(-99999999)
	}

	pointsFromChoiceSeries := df.Select([]string{"X0", "X1"}).Rapply(choicesToPointConversion)
	pointsFromChoiceSeries.SetNames("points")
	df = df.CBind(pointsFromChoiceSeries)

	fmt.Printf("%v\n", df)

	sum := func(s series.Series) series.Series {
		ints, err := s.Int()
		check(err)
		sum := 0
		for _, n := range ints {
			sum += n
		}
		return series.Ints(sum)
	}

	total, err := df.Select([]string{"points"}).Capply(sum).Elem(0, 0).Int()
	check(err)

	fmt.Printf("%v\n", total)

	return total
}

func star2(reader *bufio.Reader) int {

	dfb := dataframe.ReadCSV(reader, dataframe.WithDelimiter(' '), dataframe.HasHeader(false))

	fmt.Print(dfb)

	choicesToPointConversion := func(s series.Series) series.Series {
		player1Choice := s.Elem(0).String()
		player2Choice := s.Elem(1).String()
		var points int
		switch player2Choice {
		case "X":
			outcome := 0
			switch player1Choice {
			case "A":
				points = 3
			case "B":
				points = 1
			case "C":
				points = 2
			}
			return series.Ints(points + outcome)
		case "Y":
			outcome := 3
			switch player1Choice {
			case "A":
				points = 1
			case "B":
				points = 2
			case "C":
				points = 3
			}
			return series.Ints(points + outcome)
		case "Z":
			outcome := 6
			switch player1Choice {
			case "A":
				points = 2
			case "B":
				points = 3
			case "C":
				points = 1
			}
			return series.Ints(points + outcome)
		}
		return series.Ints(-99999999)
	}
	fmt.Printf("%v\n", dfb)

	pointsFromChoiceSeries := dfb.Select([]string{"X0", "X1"}).Rapply(choicesToPointConversion)
	pointsFromChoiceSeries.SetNames("points")
	dfb = dfb.CBind(pointsFromChoiceSeries)

	fmt.Printf("%v\n", dfb)

	sum := func(s series.Series) series.Series {
		ints, err := s.Int()
		check(err)
		sum := 0
		for _, n := range ints {
			sum += n
		}
		return series.Ints(sum)
	}

	total, err := dfb.Select([]string{"points"}).Capply(sum).Elem(0, 0).Int()
	check(err)

	fmt.Printf("%v\n", total)

	return total
}

func main() {

	f1, err := os.Open("./input.txt")
	check(err)

	r1 := bufio.NewReader(f1)

	firstAnswer := star1(r1)
	fmt.Printf("Result 1: %v\n", firstAnswer)

	f2, err := os.Open("./input.txt")
	check(err)

	r2 := bufio.NewReader(f2)
	secondAnswer := star2(r2)
	fmt.Printf("Result 2: %v\n", secondAnswer)

}
