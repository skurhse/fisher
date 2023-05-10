package eseive

func MakePrimes(lower int, upper int) []int {
	if upper < 2 {
		return []int{}
	}
	if lower < 2 {
		lower = 2
	}

	primes := make([]int, 0)

	composites := make(map[int]bool, lower-upper)

	var i int
	for i = 2; i*i < upper; i++ {
		if composites[i] {
			continue
		}
		for multiple := i + i; multiple < upper; multiple += i {
			if multiple >= lower {
				composites[multiple] = true
			}
		}
	}
	for i := lower; i < upper; i++ {
		if !composites[i] {
			primes = append(primes, i)
		}
	}

	return primes
}
