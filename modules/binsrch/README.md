# Binary Search With Common Attributes

Also known as half-interval search or logarithmic search, binary search is an algorithm that finds the position of a target value within a sorted array.

Binary search leverages the sorted nature of candidate data by repeatedly bifurcating the search interval until the target is found or the interval is empty.

Although the basic idea of binary search is straightforward, the details can be tricky. Variables used to represent the indices can result in an arithmetic overflow for very large arrays. Additionally, if there are repeat values, two binary searches must be completed, one each for the lower and upper bound.

### Problem Statement

Write a function in Go to perform a binary search against the x-coordinate in an ordered pair

##### Input
- `Coordinates`: A slice of integers (`[]int`), sorted in ascending order and may contain negative, zero, or positive integers.
- `target`: An integer (`int`) representing the value to search for in the array.

##### Output
- Return a pair of indices (`int`) of `target` in the array `numbers` if present.
- Return an error if the target is not
