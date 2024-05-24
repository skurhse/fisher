# Binary Search on Common Attributes

Also known as half-interval search or logarithmic search, *binary search* is an algorithm that finds the index of a target value within a sorted collection.

Binary search leverages the ordinal nature of candidate data by repeatedly bifurcating the search interval until the target is found or the interval becomes empty.

Although the basic idea is straightforward, the details can be tricky:
- Variables used to represent indices can result in an arithmetic overflow against large arrays.
- If there are repeated elements, two modified binary searches must be completed, one each for the lower and upper subrange bound.

### Problem Statement

Write two Go functions to perform binary searches against the x-coordinate in a sorted collection of ordered pairs:
```go
RecursiveSearch(numbers []int, target int) int

InterativeSearch(numbers []int, target int) int
```

Benchmark them to compare performance.

##### Notes
- Coordinates may be either natural or negative whole numbers.
- Ordered pairs with x-coordinate values.
