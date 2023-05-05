package search

func search(nums []int, target int) int {
	var find func(int, int) int
	find = func(l int, u int) int {

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
			return find(l+j, u)
		default:
			return find(l, u-j)
		}
	}

	return find(0, len(nums))
}
