package slices

func Sum(n []int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}
	return sum
}

func SumAll(n ...[]int) []int {
	r := make([]int, len(n))

	for i, v := range n {
		r[i] = Sum(v)
	}

	return r
}

func SumAllTails(n ...[]int) []int {
	r := make([]int, len(n))

	for i, v := range n {
		if len(v) == 0 {
			r[i] = 0
		} else {
			t := v[1:]
			r[i] = Sum(t)
		}
	}

	return r
}
