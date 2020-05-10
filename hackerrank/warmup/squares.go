package warmup

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Complete the squares function below.
func squares(a int32, b int32) int32 {
	l := int32(math.Floor(math.Sqrt(float64(a))))
	u := int32(math.Ceil(math.Sqrt(float64(b))))

	c := 0
	for i := l; i <= u; i++ {
		s := i * i
		if s >= a && s <= b {
			c++
		}
	}
	return int32(c)
}

func squares1(a int32, b int32) int32 {
	c := 0
	for i := a; i <= b; i++ {
		sqr := math.Sqrt(float64(i))
		if math.Floor(sqr) == math.Ceil(sqr) {
			c++
		}
	}
	return int32(c)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		ab := strings.Split(readLine(reader), " ")

		aTemp, err := strconv.ParseInt(ab[0], 10, 64)
		checkError(err)
		a := int32(aTemp)

		bTemp, err := strconv.ParseInt(ab[1], 10, 64)
		checkError(err)
		b := int32(bTemp)

		result := squares(a, b)

		fmt.Fprintf(writer, "%d\n", result)
	}

	writer.Flush()
}
