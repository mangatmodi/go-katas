package warmup

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

//https://www.hackerrank.com/challenges/diagonal-difference/problem
type DiagonalDifference struct{}

func diagonalDifference(arr [][]int32) int32 {
	size := len(arr)
	if size == 0 {
		return 0
	}

	var diff int32 = 0
	for i, j := 0, size-1; i < size; {
		minus := arr[i][i] - arr [i][j]
		diff = diff + minus
		i++
		j--
	}
	if diff < 0 {
		return -1 * diff
	}

	return diff
}

func (DiagonalDifference) Load(stdin io.Reader, stdout io.Writer) {
	reader := bufio.NewReaderSize(stdin, 16*1024*1024)
	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var arr [][]int32
	for i := 0; i < int(n); i++ {
		arrRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(n) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	result := diagonalDifference(arr)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}
