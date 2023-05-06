package binsearch

import (
	"fmt"
)

func ExampleOne() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
	// Output:
	// 4
}

func ExampleTwo() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2))
	// Output:
	// -1
}
