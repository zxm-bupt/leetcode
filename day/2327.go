package day

func peopleAwareOfSecret(n int, delay int, forget int) int {
	owner := make([]int, forget)
	owner[0] = 1
	for day := 2; day <= n; day++ {
		copy(owner[1:], owner[:forget-1])
		new := 0
		for i := delay; i < forget; i++ {
			new += owner[i]
			new = new % (1e9 + 7)
		}
		owner[0] = new
	}
	sum := 0
	for _, people := range owner {
		sum += people
		sum = sum % (1e9 + 7)
	}
	return sum
}

func PeopleAwareOfSecret(n int, delay int, forget int) int {
	return peopleAwareOfSecret(n, delay, forget)
}
