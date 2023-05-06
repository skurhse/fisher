package binsrch

import (
	"fmt"
)

func ExampleSearch_one() {
	fmt.Println(Search([]int{-1, 0, 3, 5, 9, 12}, 9))
	// Output:
	// 4
}

func ExampleSearch_two() {
	fmt.Println(Search([]int{-1, 0, 3, 5, 9, 12}, 2))
	// Output:
	// -1
}
