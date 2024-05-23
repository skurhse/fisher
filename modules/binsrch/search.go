package binsrch

func Search(numbers []int, target int) int {
	return searchWindow(0, len(numbers), numbers, target)
}

func searchWindow(lower int, upper int, numbers []int, target int) int {
	window := numbers[lower:upper]
	size := len(window)

	if size == 1 {
		if window[0] == target {
			return lower
		} else {
			return -1
		}
	}

	middle := size / 2
	element := window[middle]

	switch {
	case element == target:
		return lower + middle
	case element < target:
		return searchWindow(lower+middle, upper, numbers, target)
	default:
		return searchWindow(lower, upper-middle, numbers, target)
	}
}
