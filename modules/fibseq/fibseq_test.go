package fibseq

import (
	"fmt"
)

func ExampleNumber_one() {
	num := Number(2)
	fmt.Println(num)
	// Output:
	// 1
}

func ExampleNumber_two() {
	num := Number(3)
	fmt.Println(num)
	// Output:
	// 2
}

func ExampleNumber_three() {
	num := Number(4)
	fmt.Println(num)
	// Output:
	// 3
}
