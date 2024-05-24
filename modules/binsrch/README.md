# Binary Search on Common Attributes

Also known as logarithmic search, *binary search* is an algorithm that finds the index of a target value within a sorted collection.

Binary search leverages the ordinal nature of the sorted candidate data by repeatedly bifurcating the search interval until the target is found or the interval becomes empty (no result).

Although the basic idea behind the algorithm is straightforward, the details can be tricky:
- Variables used to represent indices can result in an arithmetic overflow against large collections.
- If elements are repeated, two modified binary searches must be completed in order to capture the lower and upper bounds.

### Problem Statement

Write two functions to perform binary searches against the abscissa (x-coordinate) in a sorted collection of cartesian points:
```go
type Point struct {
    X int
    Y int
}

func Search(points []Point, abscissa int) (int, int, error) {}
```

Return the half-open interval of the found subrange.

##### Notes
- Coordinates may be either natural or negative whole numbers.
- Ordered pairs with x-coordinate values.
