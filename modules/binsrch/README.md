# Binary Search on Common Attributes

Also known as half-interval search or logarithmic search, *binary search* is an algorithm that finds the position of a target value within a sorted array.

Binary search leverages the sorted nature of candidate data by repeatedly bifurcating the search interval until the target is found or the interval becomes empty.

Although the basic idea is straightforward, the details can be tricky:
- Variables used to represent indices can result in an arithmetic overflow against large arrays.
- If there are repeated elements, two modified binary searches must be completed, one each for the lower and upper subrange bound.

### Problem Statement

Write a function `Search` in Go to perform a binary search against the x-coordinate in an collection of ordered pairs.

##### Input
- `coords`: A slice of coordinate pairs, sorted in ascending x order, which may contain natural or negative integers.
- `target`: A x-coordinate representing the value to search for.

##### Output
- Return a slice containing all coordinate pair with an x-coordinate equal to `target`.

##### Notes
- `coords` may have non-unique point values.
