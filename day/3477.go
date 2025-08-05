package day

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	count := 0
	hash := make(map[int]bool, len(baskets))
	for _, fruit := range fruits {
		flag := 0
		for i, basket := range baskets {
			if fruit <= basket && !hash[i] {
				flag = 1
				hash[i] = true
				break
			}
		}
		if flag == 0 {
			count++
		}
	}
	return count
}
