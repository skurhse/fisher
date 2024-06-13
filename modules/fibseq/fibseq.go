package fibseq

const (
	First  = 0
	Second = 1
)

func Number(n int) int {
	switch n {
	case 0:
		return First
	case 1:
		return Second
	}

	seq := make([]int, 0, n)
	seq = append(seq, First, Second)
	n -= 2

	return number(n, seq)
}

func number(n int, seq []int) int {
	l := len(seq)
	next := seq[l-1] + seq[l-2]

	if n == 0 {
		return next
	}

	seq = append(seq, next)
	n--
	return number(n, seq)
}
