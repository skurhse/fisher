package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	FibFirst  = 0
	FibSecond = 1
)

const MaxN = 30

func main() {
	fibs := [MaxN]string{}

	for i := 0; i < MaxN; i++ {
		fibs[i] = strconv.Itoa(fib(i))
	}
	fibStr := strings.Join(fibs[:], ", ")

	fmt.Printf("Fibonacci sequence for %d numbers:\n", MaxN)
	fmt.Println(fibStr)
}

func fib(n int) int {
	switch n {
	case 0:
		return FibFirst
	case 1:
		return FibSecond
	}

	seq := make([]int, 0, MaxN)
	seq = append(seq, FibFirst, FibSecond)
	n -= 2

	var fib func() int
	fib = func() int {
		len := len(seq)
		next := seq[len-1] + seq[len-2]

		if n == 0 {
			return next
		} else {
			seq = append(seq, next)
			n--
			return fib()
		}
	}

	return fib()
}
