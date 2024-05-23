package binsrch

func Search(nums []int, target int) int {
	return searchWindow(0, len(nums), nums, target)
}

func searchWindow(l int, u int, nums []int, target int) int {
	window := nums[l:u]

	size := len(window)

	if size == 1 {
		if window[0] == target {
			return l
		} else {
			return -1
		}
	}

	j := size / 2
	e := window[j]

	switch {
	case e == target:
		return l + j
	case e < target:
		return searchWindow(l+j, u, nums, target)
	default:
		return searchWindow(l, u-j, nums, target)
	}
}
