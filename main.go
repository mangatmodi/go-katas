package main

import (
	"github.com/mangatmodi/go-katas/hackerrank"
	"github.com/mangatmodi/go-katas/hackerrank/warmup"
	"os"
)

func runHackerRank(solver hackerrank.Solver) {
	solver.Load(os.Stdin, os.Stdout)
}
func wait() {
	c := make(chan struct{})
	<-c
}

func main() {
	p := warmup.DiagonalDifference{}
	runHackerRank(&p)
}
