package day

func findLucky(arr []int) int {
	lucky := -1
	counts := make(map[int]int)

	for _, num := range arr {
		counts[num]++
	}

	for num, count := range counts {
		if num == count {
			lucky = max(lucky, num)
		}
	}

	return lucky
}
