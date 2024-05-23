# Binary Search

Also known as half-interval search or logarithmic search, binary search is an algorithm that finds the position of a target value within a sorted array.

Binary search leverages the sorted nature of candidate data by repeatedly bifurcating the search interval until the target is found or the search range is empty.

Although the basic idea of binary search is comparatively straightforward, the details can be surprisingly tricky. The variables used to represent the indices can result in an arithmetic overflow for very large arrays. If there are duplicates, two binary searches must be completed, one each for the lower and upper bound.

### Problem Statement

Write a function in Go to perform a binary search.

##### Input
- `numbers`: A slice of integers (`[]int`). The array is sorted in ascending order and may contain negative, zero, or positive integers.
- `target`: An integer (`int`) representing the value to search for in the array.

##### Output
- Return the index (`int`) of `target` in the array `numbers` if present.
- Returns -1 if `target` is not present in the array.
