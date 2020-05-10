package warmup

import (
	"strconv"
	"strings"
)

//https://www.hackerrank.com/challenges/time-conversion/problem

/*
 * Complete the timeConversion function below.
 */
func timeConversion(s string) string {
	arr := strings.Split(s, ":")
	temp := arr[2]
	am := temp[len(temp)-2:]
	sec := temp[0 : len(temp)-2]
	hr := arr[0]
	if am == "AM" {
		if hr == "12" {
			hr = "00"
		}
	} else {
		if hr != "12" {
			t, _ := strconv.Atoi(hr)
			hr = strconv.Itoa(t + 12)
		}
	}
	arr[0] = hr
	arr[2] = sec
	return strings.Join(arr, ":")
}

//func main() {
//	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
//
//	outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
//	checkError(err)
//
//	defer outputFile.Close()
//
//	writer := bufio.NewWriterSize(outputFile, 1024*1024)
//
//	s := readLine(reader)
//
//	result := timeConversion(s)
//
//	fmt.Fprintf(writer, "%s\n", result)
//
//	writer.Flush()
//}
