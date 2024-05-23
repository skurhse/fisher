# Binary Search on Common Attributes

Also known as half-interval search or logarithmic search, *binary search* is an algorithm that finds the position of a target value within a sorted array.

Binary search leverages the sorted nature of candidate data by repeatedly bifurcating the search interval until the target is found or the interval becomes empty.

Although the basic idea of binary search is straightforward, the details can be tricky. Variables used to represent indices can result in an arithmetic overflow against large arrays. Additionally, if there are repeat elements, two binary searches must be completed, one each for the lower and upper bound of the element range.

### Problem Statement

Write a function in Go to perform a binary search against the x-coordinate in an set of ordered pairs.

##### Input
- `coords`: A slice of coordinate pairs, sorted in ascending x order, which may contain natural or negative integers.
- `target`: A x-coordinate representing the value to search for.

##### Output
- Return a slice containing all coordinate pair with an x-coordinate equal to `target`.
