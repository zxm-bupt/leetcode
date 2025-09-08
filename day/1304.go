package day

func sumZero(n int) []int {
	res := []int{}
	if n%2 != 0 {
		res = append(res, 0)
	}

	for i := 1; len(res) < n; i++ {
		res = append(res, i)
		res = append(res, -1*i)
	}
	return res
}
