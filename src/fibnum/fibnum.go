package fibnum

const (
	FibFirst  = 0
	FibSecond = 1
)

const MaxN = 30

func Fib(n int) int {
	switch n {
	case 0:
		return FibFirst
	case 1:
		return FibSecond
	}

	seq := make([]int, 0, MaxN)
	seq = append(seq, FibFirst, FibSecond)
	n -= 2

	var fib func() int
	fib = func() int {
		len := len(seq)
		next := seq[len-1] + seq[len-2]

		if n == 0 {
			return next
		} else {
			seq = append(seq, next)
			n--
			return fib()
		}
	}

	return fib()
}
