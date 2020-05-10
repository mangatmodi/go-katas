package sorting

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

//https://www.hackerrank.com/challenges/lilys-homework/problem
//Sort in both ways and check for which way the number of swaps is smaller. The number of swaps = N - cycles.

// Complete the lilysHomework function below.
func lilysHomework(arr []int32) int32 {
	asc := make([]int32, len(arr))
	desc := make([]int32, len(arr))
	copy(asc, arr)
	copy(desc, arr)

	sort.Slice(asc, func(i, j int) bool {
		return asc[i] < asc[j]
	})

	sort.Slice(desc, func(i, j int) bool {
		return desc[i] > desc[j]
	})

	//Now do the graph magic here. Find how many have to be changed.
	m_asc := make(map[int32]int, len(arr))
	for i := range asc {
		m_asc[asc[i]] = i
	}

	m_desc := make(map[int32]int, len(arr))
	for i := range desc {
		m_desc[desc[i]] = i
	}

	ascCount := 0
	descCount := 0
	for i := range arr {
		if m_asc[arr[i]] != i {
			ascCount++
		}
		if m_desc[arr[i]] != i {
			descCount++
		}
	}
	count := int32(math.Min(float64(ascCount), float64(descCount)))
	if count > 0 {
		return count - 1
	}
	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := lilysHomework(arr)

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
