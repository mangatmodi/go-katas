package warmup

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//https://www.hackerrank.com/challenges/plus-minus/problem
type PlusMinus struct{}

func (p *PlusMinus) Load() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

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

	plusMinus(arr)
}

func plusMinus(arr []int32) {
	
}
