//https://www.hackerrank.com/challenges/magic-square-forming/problem
package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// Complete the formingMagicSquare function below.
//Calculate all diff sums, for each sum get the cost of changing
func formingMagicSquare(s [][]int32) int32 {
	c, _ := cost(0, 45, s, 0, 0)
	return c
}

//There can be only 1 such change for each overlapping row and column.NO!! WE HAVE TO CHOOSE MIN
func cost(sum float64, total int32, s [][]int32, i, j int) (int32, bool) {
	// every point we have 3 answers. Either we dont change it,
	//or we change it for vertical and try horizontal, or we try
	//or we try horizontal and try for next vertical

	var verSum int32 = 0
	var horSum int32 = 0
	for k := 0; k <= 2; k++ {
		verSum += s[k][j]
		horSum += s[i][k]
	}
	if i == 1 && j == 1 {
		//We are good
		if total == 15 && verSum == 15 && horSum == 15 {
			return int32(sum), true
		}
		if (45-total) == (15-verSum) && (45-total) == (15-horSum) {
			return int32(sum + math.Abs(float64(45-total))), true
		}
		//its a bad state
		return int32(sum), false
	}

	i, j = goNext(i, j)
	var min int32 = 1<<31 - 1
	//No change
	change, ok := cost(sum, total, s, i, j)
	if ok && change < min {
		min = change
	}
	//same as Vertical
	diff := 15 - verSum
	if diff != 0 {
		s[i][j] += diff
		change, ok = cost(sum+math.Abs(float64(diff)), total+diff, s, i, j)
		s[i][j] -= diff //make same
		if ok && change < min {
			min = change
		}
	}

	//same as horizontal
	diff = 15 - horSum
	if diff != 0 {
		s[i][j] += diff
		change, ok = cost(sum+math.Abs(float64(diff)), total+diff, s, i, j)
		s[i][j] -= diff //make same
		if ok && change < min {
			min = change
		}
	}
	return cost(sum+float64(min), total, s, i, j)
}

func goNext(i, j int) (int, int) {
	switch {
	case i == 2 && j == 0:
		return i - 1, j
	case i != 2 && j < 2:
		return i, j + 1
	case j != 2 && i < 2:
		return i + 1, j
	case i == 1 && j == 1:
		return -1, -1
	case i == 2:
		return i, j - 1
	case j == 2:
		return i + 1, j
	}
	return -1, -1
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	var s [][]int32
	for i := 0; i < 3; i++ {
		sRowTemp := strings.Split(readLine(reader), " ")

		var sRow []int32
		for _, sRowItem := range sRowTemp {
			sItemTemp, err := strconv.ParseInt(sRowItem, 10, 64)
			checkError(err)
			sItem := int32(sItemTemp)
			sRow = append(sRow, sItem)
		}

		if len(sRow) != 3 {
			panic("Bad input")
		}

		s = append(s, sRow)
	}

	result := formingMagicSquare(s)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
