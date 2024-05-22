package eseive

import (
	"fmt"
	"testing"
)

func ExampleMakePrimes_one() {
	primes := MakePrimes(0, 100)
	fmt.Println(primes)
	// Output: [2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97]
}

func ExampleMakePrimes_two() {
	primes := MakePrimes(10000, 10100)
	fmt.Println(primes)
	// Output: [10007 10009 10037 10039 10061 10067 10069 10079 10091 10093 10099]
}

func ExampleMakePrimes_three() {
	primes := MakePrimes(1000000, 1000100)
	fmt.Println(primes)
	// Output: [1000003 1000033 1000037 1000039 1000081 1000099]
}

func BenchmarkMakePrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MakePrimes(2, i)
	}
}
