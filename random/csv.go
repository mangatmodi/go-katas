package main

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
)

const (
	letters = "-%@#!-_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	max     = len(letters)
)

var rows = flag.Int("rows", 10, "number of rows")
var cols = flag.Int("cols", 20, "number of cols")

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU() - 2)
	flag.Parse()
	generateBatch()
}

//fast serial
func generate() {
	for r := 0; r < *rows; r++ {
		row()
	}
}

//Slow
func generateBatch() {
	batch := (*rows/runtime.NumCPU() - 2)
	if *rows < batch {
		batch = *rows
	}

	wg := sync.WaitGroup{}
	wg.Add(*rows / batch)
	for r := 0; r < (runtime.NumCPU() - 2); r++ {
		//we create rows/batch goroutines
		go func() {
			runtime.LockOSThread()
			defer runtime.UnlockOSThread()
			for i := 0; i < batch; i++ {
				row()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

//Still slow
func generateAtomic() {
	var count = int64(*rows)
	for r := 0; r < *rows; r++ {
		//we create rows/batch goroutines
		go func() {
			row()
			atomic.AddInt64(&count, -1)
		}()
	}
	for count > 0 {
	} //keep waiting
}

//Slow
func generateWG() {
	wg := sync.WaitGroup{}
	wg.Add(*rows)
	for r := 0; r < *rows; r++ {
		//we create rows/batch goroutines
		go func() {
			row()
			wg.Done()
		}()
	}
	wg.Wait()
}

func row() {
	sb := strings.Builder{}
	for c := 0; c < *cols; c++ {
		sb.Write(randString(20))
		if c < *cols-1 {
			sb.WriteByte(byte(','))
		}
	}
	fmt.Println(sb.String())
}

func randString(n int) []byte {
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		ret[i] = letters[rand.Intn(max)]
	}
	return ret
}
