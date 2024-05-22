package fibnum

import (
	"fmt"
)

func ExampleFib_one() {
	fibNum := Fib(2)
	fmt.Println(fibNum)
	// Output:
	// 1
}

func ExampleFib_two() {
	fibNum := Fib(3)
	fmt.Println(fibNum)
	// Output:
	// 2
}

func ExampleFib_three() {
	fibNum := Fib(4)
	fmt.Println(fibNum)
	// Output:
	// 3
}
