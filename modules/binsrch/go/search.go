package binsrch

import "errors"

type Point struct {
	X int
	Y int
}

func Search(points []Point, abscissa int) (int, int, error) {
	if len(points) == 0 {
		return 0, 0, errors.New("points empty")
	}

	lower, err := findLower(0, len(points), points, abscissa)
	if err != nil {
		return 0, 0, err
	}
	upper := findUpper(0, len(points), points, abscissa)

	return lower, upper, nil
}

func findLower(lower int, upper int, points []Point, abscissa int) (int, error) {
	size := upper - lower

	if size == 1 {
		if points[lower].X == abscissa {
			return lower, nil
		}
		return 0, errors.New("not found")
	}

	middle := size / 2

	rightLower := lower + middle
	leftUpper := upper - middle

	point := points[rightLower]
	x := point.X

	if x == abscissa {
		preceding := points[rightLower-1]
		if preceding.X != abscissa {
			return rightLower, nil
		}
	}

	if x < abscissa {
		return findLower(rightLower, upper, points, abscissa)
	}
	return findLower(lower, leftUpper, points, abscissa)
}

func findUpper(lower int, upper int, points []Point, abscissa int) int {
	size := upper - lower

	if size == 1 {
		return upper
	}

	middle := size / 2

	rightLower := lower + middle
	leftUpper := upper - middle

	point := points[rightLower]
	x := point.X

	if x == abscissa {
		if rightLower+1 < len(points) {
			succeeding := points[rightLower+1]
			if succeeding.X == abscissa {
				return findUpper(rightLower, upper, points, abscissa)
			}
			return rightLower + 1
		}
		return rightLower + 1
	}

	if x < abscissa {
		return findUpper(rightLower, upper, points, abscissa)
	}
	return findUpper(lower, leftUpper, points, abscissa)
}
