package eseive

func MakePrimes(lower int, upper int) []int {
	if upper < 2 {
		return []int{}
	}
	if lower < 2 {
		lower = 2
	}

	primes := []int{}
	composites := make(map[int]bool, lower-upper)

	for i := 2; i*i < upper; i++ {
		if composites[i] {
			continue
		}
		for m := i + i; m < upper; m += i {
			if m >= lower {
				composites[m] = true
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
