package graph

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//https://www.hackerrank.com/challenges/torque-and-development/problem

type RoadsAndLibraries struct{}
func (r *RoadsAndLibraries) Load() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	stdout  := os.Stdout
	//checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		nmC_libC_road := strings.Split(readLine(reader), " ")

		nTemp, err := strconv.ParseInt(nmC_libC_road[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		mTemp, err := strconv.ParseInt(nmC_libC_road[1], 10, 64)
		checkError(err)
		m := int32(mTemp)

		c_libTemp, err := strconv.ParseInt(nmC_libC_road[2], 10, 64)
		checkError(err)
		c_lib := int32(c_libTemp)

		c_roadTemp, err := strconv.ParseInt(nmC_libC_road[3], 10, 64)
		checkError(err)
		c_road := int32(c_roadTemp)

		var cities [][]int32
		for i := 0; i < int(m); i++ {
			citiesRowTemp := strings.Split(readLine(reader), " ")

			var citiesRow []int32
			for _, citiesRowItem := range citiesRowTemp {
				citiesItemTemp, err := strconv.ParseInt(citiesRowItem, 10, 64)
				checkError(err)
				citiesItem := int32(citiesItemTemp)
				citiesRow = append(citiesRow, citiesItem)
			}

			if len(citiesRow) != 2 {
				panic("Bad input")
			}

			cities = append(cities, citiesRow)
		}

		result := roadsAndLibraries(n, c_lib, c_road, cities)

		fmt.Fprintf(writer, "%d\n", result)
	}

	writer.Flush()
}

func roadsAndLibraries(n int32, c_lib int32, c_road int32, cities [][]int32) int64 {

	return 0
}
