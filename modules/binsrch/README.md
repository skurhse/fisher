# Binary Search

Also known as half-interval search or logarithmic search, *binary search* is a search algorithm that finds the position of a target value within a sorted array.

Binary search leverages the sorted nature of the candidate data by repeatedly dividing the search interval in half until the target is found or the search range is empty.

Although the basic idea of binary search is comparatively straightforward, the details can be surprisingly tricky. In a practical implementation, the variables used to represent the indices will often be of fixed size, and this can result in an arithmetic overflow for very large arrays

### Problem Statement

Write a function in Go to perform a binary search.

##### Input
- `numbers`: A slice of integers (`[]int`). The array is sorted in ascending order and may contain negative, zero, or positive integers.
- `target`: An integer (`int`) representing the value to search for in the array.

##### Output
- Return the index (`int`) of `target` in the array `numbers` if present.
- Returns -1 if `target` is not present in the array.
