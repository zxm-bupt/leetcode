package day

func FindSmallestInteger(nums []int, value int) int {
	return findSmallestInteger(nums, value)
}

func findSmallestInteger(nums []int, value int) int {
	hash := make(map[int]int)
	for _, num := range nums {
		hash[modLikePython(num, value)]++
	}
	count := 0
	for {
		if v, ok := hash[count%value]; ok && v > 0 {
			hash[count%value]--
			count++
		} else {
			return count
		}
	}
}

func modLikePython(a, b int) int {
	return ((a % b) + b) % b
}
